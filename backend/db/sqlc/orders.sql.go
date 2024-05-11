// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: orders.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (user_pk, status, total_price, shipping_address, shipping_date, delivered_date, is_paid)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    pk, id, user_pk, status, is_paid, total_price, shipping_address, shipping_date, delivered_date, created_at, updated_at
`

type CreateOrderParams struct {
	UserPk          int64              `json:"user_pk"`
	Status          OrderStatus        `json:"status"`
	TotalPrice      int32              `json:"total_price"`
	ShippingAddress string             `json:"shipping_address"`
	ShippingDate    pgtype.Timestamptz `json:"shipping_date"`
	DeliveredDate   pgtype.Timestamptz `json:"delivered_date"`
	IsPaid          bool               `json:"is_paid"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, createOrder,
		arg.UserPk,
		arg.Status,
		arg.TotalPrice,
		arg.ShippingAddress,
		arg.ShippingDate,
		arg.DeliveredDate,
		arg.IsPaid,
	)
	var i Order
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.UserPk,
		&i.Status,
		&i.IsPaid,
		&i.TotalPrice,
		&i.ShippingAddress,
		&i.ShippingDate,
		&i.DeliveredDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders
WHERE pk = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, pk int64) error {
	_, err := q.db.Exec(ctx, deleteOrder, pk)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT
    pk, id, user_pk, status, is_paid, total_price, shipping_address, shipping_date, delivered_date, created_at, updated_at
FROM
    orders
WHERE
    pk = $1
`

func (q *Queries) GetOrder(ctx context.Context, pk int64) (Order, error) {
	row := q.db.QueryRow(ctx, getOrder, pk)
	var i Order
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.UserPk,
		&i.Status,
		&i.IsPaid,
		&i.TotalPrice,
		&i.ShippingAddress,
		&i.ShippingDate,
		&i.DeliveredDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrderByUser = `-- name: GetOrderByUser :many
SELECT
    pk, id, user_pk, status, is_paid, total_price, shipping_address, shipping_date, delivered_date, created_at, updated_at
FROM
    orders
WHERE
    user_pk = $1
`

func (q *Queries) GetOrderByUser(ctx context.Context, userPk int64) ([]Order, error) {
	rows, err := q.db.Query(ctx, getOrderByUser, userPk)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.UserPk,
			&i.Status,
			&i.IsPaid,
			&i.TotalPrice,
			&i.ShippingAddress,
			&i.ShippingDate,
			&i.DeliveredDate,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOrders = `-- name: ListOrders :many
SELECT
    pk, id, user_pk, status, is_paid, total_price, shipping_address, shipping_date, delivered_date, created_at, updated_at
FROM
    orders
LIMIT $1 offset $2
`

type ListOrdersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error) {
	rows, err := q.db.Query(ctx, listOrders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.UserPk,
			&i.Status,
			&i.IsPaid,
			&i.TotalPrice,
			&i.ShippingAddress,
			&i.ShippingDate,
			&i.DeliveredDate,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrder = `-- name: UpdateOrder :one
UPDATE
    orders
SET
    user_pk = $2,
    status = $3,
    total_price = $4,
    shipping_address = $5,
    shipping_date = $6,
    delivered_date = $7,
    is_paid = $8
WHERE
    pk = $1
RETURNING
    pk, id, user_pk, status, is_paid, total_price, shipping_address, shipping_date, delivered_date, created_at, updated_at
`

type UpdateOrderParams struct {
	Pk              int64              `json:"pk"`
	UserPk          int64              `json:"user_pk"`
	Status          OrderStatus        `json:"status"`
	TotalPrice      int32              `json:"total_price"`
	ShippingAddress string             `json:"shipping_address"`
	ShippingDate    pgtype.Timestamptz `json:"shipping_date"`
	DeliveredDate   pgtype.Timestamptz `json:"delivered_date"`
	IsPaid          bool               `json:"is_paid"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, updateOrder,
		arg.Pk,
		arg.UserPk,
		arg.Status,
		arg.TotalPrice,
		arg.ShippingAddress,
		arg.ShippingDate,
		arg.DeliveredDate,
		arg.IsPaid,
	)
	var i Order
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.UserPk,
		&i.Status,
		&i.IsPaid,
		&i.TotalPrice,
		&i.ShippingAddress,
		&i.ShippingDate,
		&i.DeliveredDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
