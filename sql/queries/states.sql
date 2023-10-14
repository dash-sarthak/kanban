-- name: CreateState :one
INSERT INTO states (id, name, project)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetStatesByProject :many
SELECT * FROM states
WHERE project = $1;
