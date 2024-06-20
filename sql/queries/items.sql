-- name: SelectItem :one
SELECT * FROM items
WHERE id = $1;

-- name: SelectItems :many
SELECT * FROM items
ORDER BY name;

-- name: CreateItem :one
INSERT INTO items (
  name, description, year
) VALUES (
  $1, $2, $3
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

-- name: CreateItemHasUploadRelationship :one
INSERT INTO item_has_upload (
  item_id, upload_id
) VALUES (
  $1, $2
)
RETURNING *;

-- https://github.com/sqlc-dev/sqlc/issues/3238

-- name: SelectItemsWithAuthorsAndUploads :many
SELECT items.id, items.name, items.description, items.year,
CAST(
  CASE
    WHEN COUNT(authors.id) > 0
    THEN jsonb_agg(distinct jsonb_build_object('id', authors.id, 'name', authors.name))::jsonb
  END
AS jsonb) as authors,
CAST(
  CASE
    WHEN COUNT(uploads.id) > 0
    THEN jsonb_agg(distinct jsonb_build_object('id', uploads.id, 'filename', uploads.name, 'sum', uploads.sum))::jsonb
  END
AS jsonb) as uploads
FROM items
left join item_has_author on items.id = item_has_author.item_id
left join authors on item_has_author.author_id = authors.id
left join item_has_upload on items.id = item_has_upload.item_id
left join uploads on item_has_upload.upload_id = uploads.id
GROUP BY items.id
LIMIT $1 OFFSET $2;

-- name: SearchItems :many
SELECT items.id, items.name, items.description, items.year,
CAST(
  CASE
    WHEN COUNT(authors.id) > 0
    THEN jsonb_agg(distinct jsonb_build_object('id', authors.id, 'name', authors.name))::jsonb
  END
AS jsonb) as authors,
CAST(
  CASE
    WHEN COUNT(uploads.id) > 0
    THEN jsonb_agg(distinct jsonb_build_object('id', uploads.id, 'filename', uploads.name, 'sum', uploads.sum))::jsonb
  END
AS jsonb) as uploads
FROM items
left join item_has_author on items.id = item_has_author.item_id
left join authors on item_has_author.author_id = authors.id
left join item_has_upload on items.id = item_has_upload.item_id
left join uploads on item_has_upload.upload_id = uploads.id
where lower(unaccent(items.name)) like $1
group by items.id
LIMIT $2 OFFSET $3;

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
