package app

import (
	"context"
	"encoding/json"

	"go-coffeeshop/cmd/counter/config"
	"go-coffeeshop/internal/counter/domain"
	"go-coffeeshop/internal/counter/events"
	ordersUC "go-coffeeshop/internal/counter/usecases/orders"
	shared "go-coffeeshop/internal/pkg/event"
	"go-coffeeshop/pkg/postgres"
	pkgConsumer "go-coffeeshop/pkg/rabbitmq/consumer"
	pkgPublisher "go-coffeeshop/pkg/rabbitmq/publisher"
	"go-coffeeshop/proto/gen"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type App struct {
	Cfg       *config.Config
	PG        postgres.DBEngine
	AMQPConn  *amqp.Connection
	Publisher pkgPublisher.EventPublisher
	Consumer  pkgConsumer.EventConsumer

	BaristaOrderPub ordersUC.BaristaEventPublisher
	KitchenOrderPub ordersUC.KitchenEventPublisher

	ProductDomainSvc  domain.ProductDomainService
	UC                ordersUC.UseCase
	CounterGRPCServer gen.CounterServiceServer

	baristaHandler events.BaristaOrderUpdatedEventHandler
	kitchenHandler events.KitchenOrderUpdatedEventHandler
}

func New(
	cfg *config.Config,
	pg postgres.DBEngine,
	amqpConn *amqp.Connection,
	publisher pkgPublisher.EventPublisher,
	consumer pkgConsumer.EventConsumer,

	baristaOrderPub ordersUC.BaristaEventPublisher,
	kitchenOrderPub ordersUC.KitchenEventPublisher,
	productDomainSvc domain.ProductDomainService,
	uc ordersUC.UseCase,
	counterGRPCServer gen.CounterServiceServer,

	baristaHandler events.BaristaOrderUpdatedEventHandler,
	kitchenHandler events.KitchenOrderUpdatedEventHandler,
) *App {
	return &App{
		Cfg: cfg,

		PG:        pg,
		AMQPConn:  amqpConn,
		Publisher: publisher,
		Consumer:  consumer,

		BaristaOrderPub: baristaOrderPub,
		KitchenOrderPub: kitchenOrderPub,

		ProductDomainSvc:  productDomainSvc,
		UC:                uc,
		CounterGRPCServer: counterGRPCServer,

		baristaHandler: baristaHandler,
		kitchenHandler: kitchenHandler,
	}
}

func (a *App) Worker(ctx context.Context, messages <-chan amqp.Delivery) {
	for delivery := range messages {
		slog.Info("processDeliveries", "delivery_tag", delivery.DeliveryTag)
		slog.Info("received", "delivery_type", delivery.Type)

		switch delivery.Type {
		case "barista-order-updated":
			var payload shared.BaristaOrderUpdated

			err := json.Unmarshal(delivery.Body, &payload)
			if err != nil {
				slog.Error("failed to Unmarshal message", err)
			}

			err = a.baristaHandler.Handle(ctx, &payload)

			if err != nil {
				if err = delivery.Reject(false); err != nil {
					slog.Error("failed to delivery.Reject", err)
				}

				slog.Error("failed to process delivery", err)
			} else {
				err = delivery.Ack(false)
				if err != nil {
					slog.Error("failed to acknowledge delivery", err)
				}
			}
		case "kitchen-order-updated":
			var payload shared.KitchenOrderUpdated

			err := json.Unmarshal(delivery.Body, &payload)
			if err != nil {
				slog.Error("failed to Unmarshal message", err)
			}

			err = a.kitchenHandler.Handle(ctx, &payload)

			if err != nil {
				if err = delivery.Reject(false); err != nil {
					slog.Error("failed to delivery.Reject", err)
				}

				slog.Error("failed to process delivery", err)
			} else {
				err = delivery.Ack(false)
				if err != nil {
					slog.Error("failed to acknowledge delivery", err)
				}
			}
		default:
			slog.Info("default")
		}
	}

	slog.Info("Deliveries channel closed")
}
