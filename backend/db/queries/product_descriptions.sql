-- name: CreateProductDescription :one
INSERT INTO product_descriptions (product_translation_pk, introduction, prize, item_info, recommendation)
    VALUES ($1, $2, $3, $4, $5)
RETURNING *
    ;

-- name: GetProductDescriptionsByTranslationPk :one 
SELECT
    *
FROM
    product_descriptions
WHERE
    product_translation_pk = $1;

-- name: GetProductDescriptionByPk :one
SELECT
    *
FROM
    product_descriptions
WHERE
    pk = $1
    ;

-- name: UpdateProductDescription :one
UPDATE
    product_descriptions
SET
    product_translation_pk = $2,
    introduction = $3,
    prize = $4,
    item_info = $5,
    recommendation = $6
WHERE
    pk = $1
RETURNING *
    ;

-- name: DeleteProductDescription :exec
DELETE FROM product_descriptions
WHERE pk = $1
    ;
