//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-coffeeshop/cmd/counter/config"
	"go-coffeeshop/internal/counter/app/router"
	"go-coffeeshop/internal/counter/events/handlers"
	"go-coffeeshop/internal/counter/infras"
	infrasGRPC "go-coffeeshop/internal/counter/infras/grpc"
	"go-coffeeshop/internal/counter/infras/repo"
	ordersUC "go-coffeeshop/internal/counter/usecases/orders"
	"go-coffeeshop/pkg/postgres"
	"go-coffeeshop/pkg/rabbitmq"
	pkgConsumer "go-coffeeshop/pkg/rabbitmq/consumer"
	pkgPublisher "go-coffeeshop/pkg/rabbitmq/publisher"
	"google.golang.org/grpc"
)

func InitApp(
	cfg *config.Config,
	dbConnStr postgres.DBConnString,
	rabbitMQConnStr rabbitmq.RabbitMQConnStr,
	grpcServer *grpc.Server,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
		rabbitMQFunc,
		pkgPublisher.EventPublisherSet,
		pkgConsumer.EventConsumerSet,

		infras.BaristaEventPublisherSet,
		infras.KitchenEventPublisherSet,
		infrasGRPC.ProductGRPCClientSet,
		router.CounterGRPCServerSet,
		repo.RepositorySet,
		ordersUC.UseCaseSet,
		handlers.BaristaOrderUpdatedEventHandlerSet,
		handlers.KitchenOrderUpdatedEventHandlerSet,
	))
}

func dbEngineFunc(url postgres.DBConnString) (postgres.DBEngine, func(), error) {
	db, err := postgres.NewPostgresDB(url)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}

func rabbitMQFunc(url rabbitmq.RabbitMQConnStr) (*amqp.Connection, func(), error) {
	conn, err := rabbitmq.NewRabbitMQConn(url)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}
