package main

import (
	"context"
	"go-coffeeshop/cmd/product/config"
	"go-coffeeshop/internal/product/app"
	"os"

	mylogger "go-coffeeshop/pkg/logger"

	"github.com/golang/glog"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		glog.Fatal(err)
	}

	mylog := mylogger.New(cfg.Level)

	a := app.New(mylog, cfg)
	if err = a.Run(context.Background()); err != nil {
		glog.Fatal(err)
		os.Exit(1)
	}
}
