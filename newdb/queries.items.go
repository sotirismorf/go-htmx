package newdb

import (
	"encoding/json"
)

const selectItemsPopulated = `
SELECT items.id, items.name, items.group_id, groups.name as group_name,
CASE
	WHEN COUNT(authors.id) > 0
	THEN jsonb_agg(distinct jsonb_build_object('id', authors.id, 'name', authors.name))
END
as json_authors,
COUNT(*) OVER()
FROM items
LEFT JOIN groups on items.group_id = groups.id
LEFT JOIN item_has_author on items.id = item_has_author.item_id
LEFT JOIN authors on item_has_author.author_id = authors.id
GROUP BY items.id, groups.name
ORDER BY items.id
LIMIT $2 OFFSET $1;
`

func SelectItemsPopulated(page int, limit int) (ItemData, error) {
	dbItems := []schemaItem{}
	items := []Item{}

	err := db.Select(&dbItems, selectItemsPopulated, calcOffset(page, limit), limit)
	if err != nil {
		return ItemData{}, err
	}

	for _, v := range dbItems {
		var authors []Author

		json.Unmarshal([]byte(v.JsonAuthors), &authors)

		item := Item{
			ID:      v.ID,
			Name:    v.Name,
			Authors: authors,
		}

		if v.GroupID != nil && v.GroupName != nil {
			item.Group = Group{
				ID:   *v.GroupID,
				Name: *v.GroupName,
			}
		}

		items = append(items, item)
	}

	meta := Metadata{
		PageIndex: page,
		Limit:     limit,
	}

	if len(dbItems) > 0 {
		meta.TotalResults = dbItems[0].Count
		meta.TotalPages = calcPageCount(meta.TotalResults, limit)
	} else {
		meta.TotalResults = 0
		meta.TotalPages = 0
	}

	data := ItemData{
		Results:  items,
		Metadata: meta,
	}

	return data, nil
}

const selectItem = `
SELECT items.id, items.name, items.description, items.group_id, groups.name as group_name,
CASE
	WHEN COUNT(authors.id) > 0
	THEN jsonb_agg(distinct jsonb_build_object('id', authors.id, 'name', authors.name))
END
as json_authors,
CASE
  WHEN COUNT(uploads.id) > 0
  THEN jsonb_agg(distinct jsonb_build_object('id', uploads.id, 'name', uploads.name, 'sum', uploads.sum, 'type', uploads.type))::jsonb
END
as json_uploads
FROM items
LEFT JOIN groups on items.group_id = groups.id
LEFT JOIN item_has_author on items.id = item_has_author.item_id
LEFT JOIN authors on item_has_author.author_id = authors.id
LEFT JOIN item_has_upload on items.id = item_has_upload.item_id
LEFT JOIN uploads on item_has_upload.upload_id = uploads.id
WHERE items.id = 1
GROUP BY items.id, groups.name;
`

func SelectItem(id int64) (Item, error) {
	var item Item
	var row schemaItem

	err := db.Get(&row, selectItem)
	if err != nil {
		return item, err
	}

	item.ID = row.ID
	item.Name = row.Name

	if row.Description != nil {
		item.Description = *row.Description
	}

  if row.GroupID != nil && row.GroupName != nil {
    item.Group = Group{
      ID: *row.GroupID,
      Name: *row.GroupName,
    }
  }

	var authors []Author
	json.Unmarshal([]byte(row.JsonAuthors), &authors)
	item.Authors = authors

	var uploads []Upload
	json.Unmarshal([]byte(row.JsonUploads), &uploads)
	item.Uploads = uploads

	return item, nil
}
