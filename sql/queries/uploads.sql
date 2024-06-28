-- name: CreateUpload :one
INSERT INTO uploads (
  sum, name, type, size
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: SelectUploads :many
SELECT * FROM uploads
ORDER BY id;

-- name: SearchUploads :many
SELECT * FROM uploads
WHERE name like $1
ORDER BY name;

-- name: SelectSingleUpload :many
SELECT * FROM uploads
WHERE id = $1;

-- name: SelectUploadsOfItemByItemID :many
SELECT * FROM uploads
INNER JOIN item_has_upload on uploads.id = item_has_upload.upload_id
AND item_has_upload.item_id = $1;

-- name: DeleteSingleUpload :exec
DELETE FROM uploads
WHERE id = $1;
