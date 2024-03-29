package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	"go-coffeeshop/cmd/proxy/config"
	mylog "go-coffeeshop/pkg/logger"
	gen "go-coffeeshop/proto/gen"

	"github.com/golang/glog"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func newGateway(ctx context.Context, productConn *grpc.ClientConn, counterConn *grpc.ClientConn, opts []gwruntime.ServeMuxOption) (http.Handler, error) {
	mux := gwruntime.NewServeMux(opts...)

	for _, f := range []func(context.Context, *gwruntime.ServeMux, *grpc.ClientConn) error{
		gen.RegisterProductServiceHandler,
	} {
		if err := f(ctx, mux, productConn); err != nil {
			return nil, err
		}
	}

	for _, f := range []func(context.Context, *gwruntime.ServeMux, *grpc.ClientConn) error{
		gen.RegisterCounterServiceHandler,
	} {
		if err := f(ctx, mux, counterConn); err != nil {
			return nil, err
		}
	}

	return mux, nil
}

func dial(ctx context.Context, network, addr string) (*grpc.ClientConn, error) {
	switch network {
	case "tcp":
		return dialTCP(ctx, addr)
	case "unix":
		return dialUnix(ctx, addr)
	default:
		return nil, fmt.Errorf("unsupported network type %q", network)
	}
}

func dialTCP(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func dialUnix(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	d := func(ctx context.Context, addr string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "unix", addr)
	}

	return grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(d))
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)

				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))

	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
}

func withLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		glog.Infof("Run  %s %s", r.Method, r.URL)

		h.ServeHTTP(w, r)
	})
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		glog.Fatalf("Config error: %s", err)
	}

	logger := mylog.New(cfg.Log.Level)
	logger.Info("Init %s %s\n", cfg.Name, cfg.Version)

	productConn, err := dial(ctx, "tcp", fmt.Sprintf("%s:%d", cfg.ProductHost, cfg.ProductPort))
	if err != nil {
		logger.Fatal("%s", err)
	}

	go func() {
		<-ctx.Done()

		if err = productConn.Close(); err != nil {
			glog.Errorf("Failed to close a product client connection to the gRPC server: %v", err)
		}
	}()

	counterConn, err := dial(ctx, "tcp", fmt.Sprintf("%s:%d", cfg.CounterHost, cfg.CounterPort))
	if err != nil {
		logger.Fatal("%s", err)
	}

	go func() {
		<-ctx.Done()

		if err = counterConn.Close(); err != nil {
			glog.Errorf("Failed to close a counter client connection to the gRPC server: %v", err)
		}
	}()

	mux := http.NewServeMux()

	gw, err := newGateway(ctx, productConn, counterConn, nil)
	if err != nil {
		logger.Fatal("%s", err)
	}

	mux.Handle("/", gw)

	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: allowCORS(withLogger(mux)),
	}

	go func() {
		<-ctx.Done()
		glog.Infof("Shutting down the http server")

		if err := s.Shutdown(context.Background()); err != nil {
			glog.Errorf("Failed to shutdown http server: %v", err)
		}
	}()

	glog.Infof("Starting listening at %s", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))

	if err := s.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		glog.Errorf("Failed to listen and serve: %v", err)
	}
}
