-- name: CreateOrder :one
INSERT INTO orders (user_pk, status, total_price, shipping_address, shipping_date, delivered_date, is_paid, email, phone)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING
    *;

-- name: GetOrder :one
SELECT
    *
FROM
    orders
WHERE
    pk = $1;

-- name: GetOrderByUser :many
SELECT
    *
FROM
    orders
WHERE
    user_pk = $1;

-- name: ListOrders :many
SELECT
    *
FROM
    orders
LIMIT $1 offset $2;

-- name: ListOrdersByStatus :many
SELECT
    *
FROM
    orders
WHERE
    status = $1
LIMIT $2 offset $3;

-- name: ListOrdersByUser :many
SELECT
    *
FROM
    orders
WHERE
    user_pk = $1
LIMIT $2 offset $3;

-- name: UpdateOrder :one
UPDATE
    orders
SET
    user_pk = $2,
    status = $3,
    total_price = $4,
    shipping_address = $5,
    shipping_date = $6,
    delivered_date = $7,
    is_paid = $8,
    email = $9,
    phone = $10
WHERE
    pk = $1
RETURNING
    *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE pk = $1;
