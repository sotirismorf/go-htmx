// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: items.sql

package schema

import (
	"context"
)

const createItem = `-- name: CreateItem :one
INSERT INTO items (
  name, description, year
) VALUES (
  $1, $2, $3
)
RETURNING id, name, description, group_id, year
`

type CreateItemParams struct {
	Name        string
	Description *string
	Year        int16
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRow(ctx, createItem, arg.Name, arg.Description, arg.Year)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.GroupID,
		&i.Year,
	)
	return i, err
}

const createItemHasAuthorRelationship = `-- name: CreateItemHasAuthorRelationship :one
INSERT INTO item_has_author (
  item_id, author_id
) VALUES (
  $1, $2
)
RETURNING item_id, author_id
`

type CreateItemHasAuthorRelationshipParams struct {
	ItemID   int64
	AuthorID int64
}

func (q *Queries) CreateItemHasAuthorRelationship(ctx context.Context, arg CreateItemHasAuthorRelationshipParams) (ItemHasAuthor, error) {
	row := q.db.QueryRow(ctx, createItemHasAuthorRelationship, arg.ItemID, arg.AuthorID)
	var i ItemHasAuthor
	err := row.Scan(&i.ItemID, &i.AuthorID)
	return i, err
}

const createItemHasUploadRelationship = `-- name: CreateItemHasUploadRelationship :one
INSERT INTO item_has_upload (
  item_id, upload_id
) VALUES (
  $1, $2
)
RETURNING item_id, upload_id
`

type CreateItemHasUploadRelationshipParams struct {
	ItemID   int64
	UploadID int64
}

func (q *Queries) CreateItemHasUploadRelationship(ctx context.Context, arg CreateItemHasUploadRelationshipParams) (ItemHasUpload, error) {
	row := q.db.QueryRow(ctx, createItemHasUploadRelationship, arg.ItemID, arg.UploadID)
	var i ItemHasUpload
	err := row.Scan(&i.ItemID, &i.UploadID)
	return i, err
}

const deleteItem = `-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1
`

func (q *Queries) DeleteItem(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteItem, id)
	return err
}

const searchItems = `-- name: SearchItems :many
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
AS jsonb) as uploads,
COUNT(*) OVER()
FROM items
left join item_has_author on items.id = item_has_author.item_id
left join authors on item_has_author.author_id = authors.id
left join item_has_upload on items.id = item_has_upload.item_id
left join uploads on item_has_upload.upload_id = uploads.id
where lower(unaccent(items.name)) like $1
group by items.id
LIMIT $2 OFFSET $3
`

type SearchItemsParams struct {
	Name   string
	Limit  int32
	Offset int32
}

type SearchItemsRow struct {
	ID          int64
	Name        string
	Description *string
	Year        int16
	Authors     []byte
	Uploads     []byte
	Count       int64
}

func (q *Queries) SearchItems(ctx context.Context, arg SearchItemsParams) ([]SearchItemsRow, error) {
	rows, err := q.db.Query(ctx, searchItems, arg.Name, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchItemsRow
	for rows.Next() {
		var i SearchItemsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Year,
			&i.Authors,
			&i.Uploads,
			&i.Count,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectItem = `-- name: SelectItem :one
SELECT id, name, description, group_id, year FROM items
WHERE id = $1
`

func (q *Queries) SelectItem(ctx context.Context, id int64) (Item, error) {
	row := q.db.QueryRow(ctx, selectItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.GroupID,
		&i.Year,
	)
	return i, err
}

const selectItemPopulated = `-- name: SelectItemPopulated :one
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
group by items.id
`

type SelectItemPopulatedRow struct {
	ID          int64
	Name        string
	Description *string
	Authors     []byte
	Uploads     []byte
}

func (q *Queries) SelectItemPopulated(ctx context.Context, id int64) (SelectItemPopulatedRow, error) {
	row := q.db.QueryRow(ctx, selectItemPopulated, id)
	var i SelectItemPopulatedRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Authors,
		&i.Uploads,
	)
	return i, err
}

const selectItems = `-- name: SelectItems :many
SELECT id, name, description, group_id, year FROM items
ORDER BY name
`

func (q *Queries) SelectItems(ctx context.Context) ([]Item, error) {
	rows, err := q.db.Query(ctx, selectItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.GroupID,
			&i.Year,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectItemsWithAuthorsAndUploads = `-- name: SelectItemsWithAuthorsAndUploads :many

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
AS jsonb) as uploads,
COUNT(*) OVER()
FROM items
left join item_has_author on items.id = item_has_author.item_id
left join authors on item_has_author.author_id = authors.id
left join item_has_upload on items.id = item_has_upload.item_id
left join uploads on item_has_upload.upload_id = uploads.id
GROUP BY items.id
LIMIT $1 OFFSET $2
`

type SelectItemsWithAuthorsAndUploadsParams struct {
	Limit  int32
	Offset int32
}

type SelectItemsWithAuthorsAndUploadsRow struct {
	ID          int64
	Name        string
	Description *string
	Year        int16
	Authors     []byte
	Uploads     []byte
	Count       int64
}

// https://github.com/sqlc-dev/sqlc/issues/3238
func (q *Queries) SelectItemsWithAuthorsAndUploads(ctx context.Context, arg SelectItemsWithAuthorsAndUploadsParams) ([]SelectItemsWithAuthorsAndUploadsRow, error) {
	rows, err := q.db.Query(ctx, selectItemsWithAuthorsAndUploads, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectItemsWithAuthorsAndUploadsRow
	for rows.Next() {
		var i SelectItemsWithAuthorsAndUploadsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Year,
			&i.Authors,
			&i.Uploads,
			&i.Count,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectSingleItemWithAuthors = `-- name: SelectSingleItemWithAuthors :one
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
group by items.id
`

type SelectSingleItemWithAuthorsRow struct {
	ID          int64
	Name        string
	Description *string
	Authors     []byte
}

func (q *Queries) SelectSingleItemWithAuthors(ctx context.Context, id int64) (SelectSingleItemWithAuthorsRow, error) {
	row := q.db.QueryRow(ctx, selectSingleItemWithAuthors, id)
	var i SelectSingleItemWithAuthorsRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Authors,
	)
	return i, err
}
