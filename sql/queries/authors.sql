-- name: CreateAuthor :one
INSERT INTO authors (id, name, username, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAuthors :many
SELECT * FROM authors;

-- name: GetAuthorID :one
SELECT id FROM authors
WHERE username =  $1;

-- name: DeleteAuthor :one
DELETE FROM authors
WHERE username = $1 and password = $2
RETURNING username;
