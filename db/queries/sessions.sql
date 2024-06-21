-- name: CreateSession :one
INSERT INTO sessions (user_id, ip_address, user_agent)
VALUES (sqlc.arg(user_id), sqlc.arg(ip_address), sqlc.arg(user_agent))
RETURNING *;

-- name: GetSessionById :one
SELECT *
FROM sessions
WHERE id = $1;

-- name: DeleteSessionById :exec
DELETE
FROM sessions
WHERE id = $1;