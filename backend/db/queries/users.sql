-- name: CreateUser :one
INSERT INTO users (email, phone, line_id, birth_year, gender, PASSWORD, first_name, last_name,
    LANGUAGE, address, longitude, latitude)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING
    *;

-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    pk = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = $1;

-- name: GetUserByPhone :one
SELECT
    *
FROM
    users
WHERE
    phone = $1;

-- name: GetUserById :one
SELECT
    *
FROM
    users
WHERE
    id = $1;

-- name: ListUsers :many
SELECT
    *
FROM
    users
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
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
    address = $8,
    line_id = $9,
    gender = $10,
    birth_year = $11,
    longitude = $12,
    latitude = $13
WHERE
    pk = $1
RETURNING
    *;

-- name: UpdateUserById :one
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
    address = $8,
    line_id = $9,
    gender = $10,
    birth_year = $11,
    longitude = $12,
    latitude = $13
WHERE
    id = $1
RETURNING
    *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE pk = $1;
