-- name: SelectAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: SelectAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: SearchAuthors :many
SELECT * FROM authors
WHERE name like $1
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE authors
set name = $2,
bio = $3
WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;
