package infras

import (
	"context"

	"go-coffeeshop/internal/counter/usecases/orders"
	"go-coffeeshop/pkg/rabbitmq/publisher"

	"github.com/google/wire"
)

var (
	BaristaEventPublisherSet = wire.NewSet(NewBaristaEventPublisher)
	KitchenEventPublisherSet = wire.NewSet(NewKitchenEventPublisher)
)

type (
	baristaEventPublisher struct {
		pub publisher.EventPublisher
	}
	kitchenEventPublisher struct {
		pub publisher.EventPublisher
	}
)

func NewBaristaEventPublisher(pub publisher.EventPublisher) orders.BaristaEventPublisher {
	return &baristaEventPublisher{
		pub: pub,
	}
}

func (p *baristaEventPublisher) Configure(opts ...publisher.Option) {
	p.pub.Configure(opts...)
}

func (p *baristaEventPublisher) Publish(ctx context.Context, body []byte, contentType string) error {
	return p.pub.Publish(ctx, body, contentType)
}

func NewKitchenEventPublisher(pub publisher.EventPublisher) orders.KitchenEventPublisher {
	return &kitchenEventPublisher{
		pub: pub,
	}
}

func (p *kitchenEventPublisher) Configure(opts ...publisher.Option) {
	p.pub.Configure(opts...)
}

func (p *kitchenEventPublisher) Publish(ctx context.Context, body []byte, contentType string) error {
	return p.pub.Publish(ctx, body, contentType)
}
