-- name: CreateProductTranslation :one
INSERT INTO product_translations (product_pk, language, name, category)
    VALUES ($1, $2, $3, $4)
RETURNING
    *;

-- name: GetProductTranslation :one
SELECT
    *
FROM
    product_translations
WHERE
    product_pk = $1
    AND
    language = $2
LIMIT 1;

-- name: GetProductTranslationByName :one
SELECT
    *
FROM
    product_translations
WHERE
    name = $1
    AND
    language = $2;

-- name: GetProductTranslations :many
SELECT
    *
FROM
    product_translations
WHERE
    product_pk = $1;

-- name: UpdateProductTranslation :one
UPDATE
    product_translations
SET
    name = $3,
    category = $4
WHERE
    product_pk = $1
    AND
    language = $2
RETURNING
    *;

-- name: DeleteProductTranslation :exec
DELETE FROM product_translations
WHERE product_pk = $1
    AND
    language = $2;
