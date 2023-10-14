-- name: CreateProject :one
INSERT INTO projects (id, name, author, description)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetProjects :many
SELECT * FROM projects;

-- name: GetProjectsByAuthor :many
SELECT * FROM projects
WHERE author = $1;