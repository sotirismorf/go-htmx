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
