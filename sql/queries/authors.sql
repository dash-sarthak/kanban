-- name: CreateAuthor :one
INSERT INTO authors (id, name, username)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAuthors :many
SELECT * FROM authors;

-- name: DeleteAuthor :one
DELETE FROM authors
WHERE username = $1
RETURNING username;
