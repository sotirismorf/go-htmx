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

-- name: SelectItemPopulated :one
SELECT items.id, items.name, items.description,
CAST(
  CASE
    WHEN COUNT(authors.id) > 0
    THEN jsonb_agg(distinct jsonb_build_object('id', authors.id, 'name', authors.name))::jsonb
  END
AS jsonb) as authors,
CAST(
  CASE
    WHEN COUNT(uploads.id) > 0
    THEN jsonb_agg(distinct jsonb_build_object('id', uploads.id, 'filename', uploads.name, 'sum', uploads.sum, 'type', uploads.type))::jsonb
  END
AS jsonb) as uploads
FROM items
left join item_has_author on items.id = item_has_author.item_id
left join authors on item_has_author.author_id = authors.id
left join item_has_upload on items.id = item_has_upload.item_id
left join uploads on item_has_upload.upload_id = uploads.id
where items.id = $1
group by items.id;
