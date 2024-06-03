-- name: CreateOrderDetail :one
INSERT INTO order_details (order_pk, product_pk, quantity, price)
    VALUES ($1, $2, $3, $4)
RETURNING
    *;

-- name: GetOrderDetail :one
SELECT
    *
FROM
    order_details
WHERE
    pk = $1;

-- name: GetOrderDetailByOrder :many
SELECT
    *
FROM
    order_details
WHERE
    order_pk = $1;

-- name: UpdateOrderDetail :one
UPDATE
    order_details
SET
    order_pk = $2,
    product_pk = $3,
    quantity = $4,
    price = $5,
    discount = $6
WHERE
    pk = $1
RETURNING
    *;

-- name: DeleteOrderDetail :exec
DELETE FROM order_details
WHERE pk = $1;
