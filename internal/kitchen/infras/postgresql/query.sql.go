// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createOrder = `-- name: CreateOrder :one

INSERT INTO
    kitchen.kitchen_orders (
        id,
        order_id,
        item_type,
        item_name,
        time_up,
        created,
        updated
    )
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, order_id, item_type, item_name, time_up, created, updated
`

type CreateOrderParams struct {
	ID       uuid.UUID    `json:"id"`
	OrderID  uuid.UUID    `json:"order_id"`
	ItemType int32        `json:"item_type"`
	ItemName string       `json:"item_name"`
	TimeUp   time.Time    `json:"time_up"`
	Created  time.Time    `json:"created"`
	Updated  sql.NullTime `json:"updated"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (KitchenKitchenOrder, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.ID,
		arg.OrderID,
		arg.ItemType,
		arg.ItemName,
		arg.TimeUp,
		arg.Created,
		arg.Updated,
	)
	var i KitchenKitchenOrder
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.ItemType,
		&i.ItemName,
		&i.TimeUp,
		&i.Created,
		&i.Updated,
	)
	return i, err
}
