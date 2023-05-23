package products

import (
	"context"

	"go-coffeeshop/internal/product/domain"
)

type UseCase interface {
	GetItemTypes(context.Context) ([]*domain.ItemTypeDto, error)
	GetItemsByType(context.Context, string) ([]*domain.ItemDto, error)
}
