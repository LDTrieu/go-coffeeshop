package eventhandlers

import (
	"context"

	"go-coffeeshop/internal/pkg/event"
)

type KitchenOrderedEventHandler interface {
	Handle(context.Context, event.KitchenOrdered) error
}
