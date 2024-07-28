// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: product_translations.sql

package db

import (
	"context"
)

const createProductTranslation = `-- name: CreateProductTranslation :one
INSERT INTO product_translations (product_pk, language, name, category)
    VALUES ($1, $2, $3, $4)
RETURNING
    pk, product_pk, name, language, category, created_at, updated_at
`

type CreateProductTranslationParams struct {
	ProductPk int64        `json:"product_pk"`
	Language  LanguageCode `json:"language"`
	Name      string       `json:"name"`
	Category  string       `json:"category"`
}

func (q *Queries) CreateProductTranslation(ctx context.Context, arg CreateProductTranslationParams) (ProductTranslation, error) {
	row := q.db.QueryRow(ctx, createProductTranslation,
		arg.ProductPk,
		arg.Language,
		arg.Name,
		arg.Category,
	)
	var i ProductTranslation
	err := row.Scan(
		&i.Pk,
		&i.ProductPk,
		&i.Name,
		&i.Language,
		&i.Category,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProductTranslation = `-- name: DeleteProductTranslation :exec
DELETE FROM product_translations
WHERE product_pk = $1
    AND
    language = $2
`

type DeleteProductTranslationParams struct {
	ProductPk int64        `json:"product_pk"`
	Language  LanguageCode `json:"language"`
}

func (q *Queries) DeleteProductTranslation(ctx context.Context, arg DeleteProductTranslationParams) error {
	_, err := q.db.Exec(ctx, deleteProductTranslation, arg.ProductPk, arg.Language)
	return err
}

const getProductTranslation = `-- name: GetProductTranslation :one
SELECT
    pk, product_pk, name, language, category, created_at, updated_at
FROM
    product_translations
WHERE
    product_pk = $1
    AND
    language = $2
LIMIT 1
`

type GetProductTranslationParams struct {
	ProductPk int64        `json:"product_pk"`
	Language  LanguageCode `json:"language"`
}

func (q *Queries) GetProductTranslation(ctx context.Context, arg GetProductTranslationParams) (ProductTranslation, error) {
	row := q.db.QueryRow(ctx, getProductTranslation, arg.ProductPk, arg.Language)
	var i ProductTranslation
	err := row.Scan(
		&i.Pk,
		&i.ProductPk,
		&i.Name,
		&i.Language,
		&i.Category,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProductTranslationByName = `-- name: GetProductTranslationByName :one
SELECT
    pk, product_pk, name, language, category, created_at, updated_at
FROM
    product_translations
WHERE
    name = $1
    AND
    language = $2
`

type GetProductTranslationByNameParams struct {
	Name     string       `json:"name"`
	Language LanguageCode `json:"language"`
}

func (q *Queries) GetProductTranslationByName(ctx context.Context, arg GetProductTranslationByNameParams) (ProductTranslation, error) {
	row := q.db.QueryRow(ctx, getProductTranslationByName, arg.Name, arg.Language)
	var i ProductTranslation
	err := row.Scan(
		&i.Pk,
		&i.ProductPk,
		&i.Name,
		&i.Language,
		&i.Category,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProductTranslations = `-- name: GetProductTranslations :many
SELECT
    pk, product_pk, name, language, category, created_at, updated_at
FROM
    product_translations
WHERE
    product_pk = $1
`

func (q *Queries) GetProductTranslations(ctx context.Context, productPk int64) ([]ProductTranslation, error) {
	rows, err := q.db.Query(ctx, getProductTranslations, productPk)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductTranslation
	for rows.Next() {
		var i ProductTranslation
		if err := rows.Scan(
			&i.Pk,
			&i.ProductPk,
			&i.Name,
			&i.Language,
			&i.Category,
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

const updateProductTranslation = `-- name: UpdateProductTranslation :one
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
    pk, product_pk, name, language, category, created_at, updated_at
`

type UpdateProductTranslationParams struct {
	ProductPk int64        `json:"product_pk"`
	Language  LanguageCode `json:"language"`
	Name      string       `json:"name"`
	Category  string       `json:"category"`
}

func (q *Queries) UpdateProductTranslation(ctx context.Context, arg UpdateProductTranslationParams) (ProductTranslation, error) {
	row := q.db.QueryRow(ctx, updateProductTranslation,
		arg.ProductPk,
		arg.Language,
		arg.Name,
		arg.Category,
	)
	var i ProductTranslation
	err := row.Scan(
		&i.Pk,
		&i.ProductPk,
		&i.Name,
		&i.Language,
		&i.Category,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
