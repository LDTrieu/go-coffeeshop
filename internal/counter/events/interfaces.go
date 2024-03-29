package events

import (
	"context"

	"go-coffeeshop/internal/pkg/event"
)

type (
	BaristaOrderUpdatedEventHandler interface {
		Handle(context.Context, *event.BaristaOrderUpdated) error
	}

	KitchenOrderUpdatedEventHandler interface {
		Handle(context.Context, *event.KitchenOrderUpdated) error
	}
)
