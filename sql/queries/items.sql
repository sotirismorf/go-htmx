-- name: SelectItem :one
SELECT * FROM items
WHERE id = $1;

-- name: SelectItems :many
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

-- name: SelectItemsWithAuthors :many
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

-- name: SelectSingleItemWithAuthors :one
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
