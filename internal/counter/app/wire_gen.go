package app

import (
	"go-coffeeshop/cmd/counter/config"
	"go-coffeeshop/internal/counter/app/router"
	"go-coffeeshop/internal/counter/events/handlers"
	"go-coffeeshop/internal/counter/infras"
	grpc2 "go-coffeeshop/internal/counter/infras/grpc"
	"go-coffeeshop/internal/counter/infras/repo"
	"go-coffeeshop/internal/counter/usecases/orders"
	"go-coffeeshop/pkg/postgres"
	"go-coffeeshop/pkg/rabbitmq"
	"go-coffeeshop/pkg/rabbitmq/consumer"
	"go-coffeeshop/pkg/rabbitmq/publisher"

	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitApp(cfg *config.Config, dbConnStr postgres.DBConnString, rabbitMQConnStr rabbitmq.RabbitMQConnStr, grpcServer *grpc.Server) (*App, func(), error) {
	dbEngine, cleanup, err := dbEngineFunc(dbConnStr)
	if err != nil {
		return nil, nil, err
	}
	connection, cleanup2, err := rabbitMQFunc(rabbitMQConnStr)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	eventPublisher, err := publisher.NewPublisher(connection)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	eventConsumer, err := consumer.NewConsumer(connection)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	baristaEventPublisher := infras.NewBaristaEventPublisher(eventPublisher)
	kitchenEventPublisher := infras.NewKitchenEventPublisher(eventPublisher)
	productDomainService, err := grpc2.NewGRPCProductClient(cfg)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	orderRepo := repo.NewOrderRepo(dbEngine)
	useCase := orders.NewUseCase(orderRepo, productDomainService, baristaEventPublisher, kitchenEventPublisher)
	counterServiceServer := router.NewGRPCCounterServer(grpcServer, cfg, useCase)
	baristaOrderUpdatedEventHandler := handlers.NewBaristaOrderUpdatedEventHandler(orderRepo)
	kitchenOrderUpdatedEventHandler := handlers.NewKitchenOrderUpdatedEventHandler(orderRepo)
	app := New(cfg, dbEngine, connection, eventPublisher, eventConsumer, baristaEventPublisher, kitchenEventPublisher, productDomainService, useCase, counterServiceServer, baristaOrderUpdatedEventHandler, kitchenOrderUpdatedEventHandler)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

func dbEngineFunc(url postgres.DBConnString) (postgres.DBEngine, func(), error) {
	db, err := postgres.NewPostgresDB(url)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}

func rabbitMQFunc(url rabbitmq.RabbitMQConnStr) (*amqp091.Connection, func(), error) {
	conn, err := rabbitmq.NewRabbitMQConn(url)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}
