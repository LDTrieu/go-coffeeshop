package eventhandlers

import (
	"context"

	"go-coffeeshop/internal/pkg/event"
)

type BaristaOrderedEventHandler interface {
	Handle(context.Context, event.BaristaOrdered) error
}
