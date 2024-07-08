-- name: CreateSession :one
INSERT INTO sessions (id, user_id, ip_address, user_agent, last_activity)
VALUES (sqlc.arg(id), sqlc.arg(user_id), sqlc.arg(ip_address), sqlc.arg(user_agent), sqlc.arg(last_activity))
RETURNING *;

-- name: GetSessionById :one
SELECT *
FROM sessions
WHERE id = $1;

-- name: DeleteSessionById :exec
DELETE
FROM sessions
WHERE id = $1;