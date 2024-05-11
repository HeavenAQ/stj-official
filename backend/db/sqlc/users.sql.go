// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, phone, PASSWORD, first_name, last_name,
    LANGUAGE, address)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    pk, id, email, phone, password, first_name, last_name, language, address, last_login, created_at, updated_at
`

type CreateUserParams struct {
	Email     string       `json:"email"`
	Phone     string       `json:"phone"`
	Password  string       `json:"password"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Language  LanguageCode `json:"language"`
	Address   string       `json:"address"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Email,
		arg.Phone,
		arg.Password,
		arg.FirstName,
		arg.LastName,
		arg.Language,
		arg.Address,
	)
	var i User
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.Language,
		&i.Address,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE pk = $1
`

func (q *Queries) DeleteUser(ctx context.Context, pk int64) error {
	_, err := q.db.Exec(ctx, deleteUser, pk)
	return err
}

const getUser = `-- name: GetUser :one
SELECT
    pk, id, email, phone, password, first_name, last_name, language, address, last_login, created_at, updated_at
FROM
    users
WHERE
    pk = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, pk int64) (User, error) {
	row := q.db.QueryRow(ctx, getUser, pk)
	var i User
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.Language,
		&i.Address,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT
    pk, id, email, phone, password, first_name, last_name, language, address, last_login, created_at, updated_at
FROM
    users
WHERE
    email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.Language,
		&i.Address,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByPhone = `-- name: GetUserByPhone :one
SELECT
    pk, id, email, phone, password, first_name, last_name, language, address, last_login, created_at, updated_at
FROM
    users
WHERE
    phone = $1
`

func (q *Queries) GetUserByPhone(ctx context.Context, phone string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByPhone, phone)
	var i User
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.Language,
		&i.Address,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT
    pk, id, email, phone, password, first_name, last_name, language, address, last_login, created_at, updated_at
FROM
    users
LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.Email,
			&i.Phone,
			&i.Password,
			&i.FirstName,
			&i.LastName,
			&i.Language,
			&i.Address,
			&i.LastLogin,
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

const updateUser = `-- name: UpdateUser :one
UPDATE
    users
SET
    email = $2,
    phone = $3,
    PASSWORD = $4,
    first_name = $5,
    last_name = $6,
    LANGUAGE =
    $7,
    address = $8
WHERE
    pk = $1
RETURNING
    pk, id, email, phone, password, first_name, last_name, language, address, last_login, created_at, updated_at
`

type UpdateUserParams struct {
	Pk        int64        `json:"pk"`
	Email     string       `json:"email"`
	Phone     string       `json:"phone"`
	Password  string       `json:"password"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Language  LanguageCode `json:"language"`
	Address   string       `json:"address"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Pk,
		arg.Email,
		arg.Phone,
		arg.Password,
		arg.FirstName,
		arg.LastName,
		arg.Language,
		arg.Address,
	)
	var i User
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.Language,
		&i.Address,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
