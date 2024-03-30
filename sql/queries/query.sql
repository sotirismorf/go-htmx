-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
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

-- name: GetItem :one
SELECT * FROM items
WHERE id = $1;

-- name: ListItems :many
SELECT * FROM items
ORDER BY name;

-- name: CreateItem :one
INSERT INTO items (
  name, description
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1;

-- name: CreateItemHasAuthorRelationship :one
INSERT INTO item_has_author (
  item_id, author_id
) VALUES (
  $1, $2
)
RETURNING *;

-- https://github.com/sqlc-dev/sqlc/issues/3238

-- name: ListItemsWithAuthors :many
SELECT items.id, items.name, items.description,
CAST(
  CASE
    WHEN (array_length(array_remove(array_agg(authors.id), null), 1) > 0)
    THEN jsonb_agg((authors.id, authors.name))::jsonb
  END
AS jsonb) as authors
FROM items
left join item_has_author on items.id = item_has_author.item_id
left join authors on item_has_author.author_id = authors.id
group by items.id;

-- name: ListSingleItemWithAuthors :one
SELECT items.id, items.name, items.description,
CAST(
  CASE
    WHEN (array_length(array_remove(array_agg(authors.id), null), 1) > 0)
    THEN jsonb_agg((authors.id, authors.name))::jsonb
  END
AS jsonb) as authors
FROM items
left join item_has_author on items.id = item_has_author.item_id
left join authors on item_has_author.author_id = authors.id
where items.id = $1
group by items.id;
