-- name: CreateUser :one
INSERT INTO users (id, name, password_hash, created_at, updated_at)
values (sqlc.arg(id), sqlc.arg(name), sqlc.arg(password_hash), sqlc.arg(created_at), sqlc.arg(updated_at))
RETURNING *;

-- name: ListUsers :many
SELECT *
FROM users;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByName :one
SELECT *
FROM users
WHERE name = $1;

-- name: UpdateUser :one
UPDATE users
SET password_hash = sqlc.arg(password_hash)
WHERE id = $1
RETURNING *;
