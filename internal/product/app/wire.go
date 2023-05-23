//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"go-coffeeshop/cmd/product/config"
	"go-coffeeshop/internal/product/app/router"
	"go-coffeeshop/internal/product/infras/repo"
	productsUC "go-coffeeshop/internal/product/usecases/products"
	"google.golang.org/grpc"
)

func InitApp(
	cfg *config.Config,
	grpcServer *grpc.Server,
) (*App, error) {
	panic(wire.Build(
		New,
		router.ProductGRPCServerSet,
		repo.RepositorySet,
		productsUC.UseCaseSet,
	))
}
