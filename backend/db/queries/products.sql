-- name: CreateProduct :one
INSERT INTO products (price, "imageURLs", status, quantity)
    VALUES ($1, $2, $3, $4)
RETURNING
    *;

-- name: GetProduct :one
SELECT
    *
FROM
    products
WHERE
    pk = $1;

-- name: ListProducts :many
SELECT
    *
FROM
    products
LIMIT $1 offset $2;

-- name: UpdateProduct :one
UPDATE
    products
SET
    price = $2,
    "imageURLs" = $3,
    status = $4,
    quantity = $5
WHERE
    pk = $1
RETURNING
    *;

-- name: UpdateProductIdxImageURL :one
UPDATE
    products
SET
    "imageURLs"[$2] = $3
WHERE
    pk = $1
RETURNING
    *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE pk = $1;

-- name: DeleteProductById :exec
DELETE FROM products
WHERE id = $1;

-- name: GetProdcutByStatus :many
SELECT
    *
FROM
    products
WHERE
    status = $1;

-- name: GetProductByPrice :many
SELECT
    *
FROM
    products
WHERE
    price >= $1
    AND price <= $2;

-- name: GetProductByQuantity :many
SELECT
    *
FROM
    products
WHERE
    quantity >= $1
    AND quantity <= $2;

-- name: GetProductByID :one
SELECT
    *
FROM
    products
WHERE
    id = $1;
