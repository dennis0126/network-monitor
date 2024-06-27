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
