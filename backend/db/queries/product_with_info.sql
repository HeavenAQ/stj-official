-- name: GetProductWithInfo :one
SELECT
    products.price,
    products.id,
    products."imageURLs",
    products.status,
    products.quantity,
    products.created_at,
    products.updated_at,
    product_translations.category,
    product_translations.name,
    product_translations.description
FROM
    products
    INNER JOIN product_translations ON products.pk = product_translations.product_pk
WHERE
    products.id = $1
    AND product_translations.language = $2;

-- name: ListProductWithInfo :many
SELECT
    products.price,
    products.id,
    products."imageURLs",
    products.status,
    products.quantity,
    products.created_at,
    products.updated_at,
    product_translations.category,
    product_translations.name,
    product_translations.description
FROM
    products
    INNER JOIN product_translations ON products.pk = product_translations.product_pk
WHERE
    product_translations.language = $1
LIMIT $2 offset $3;
