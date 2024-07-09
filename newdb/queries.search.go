package newdb

import (
	"encoding/json"
	"fmt"
)

const searchItems = `
SELECT items.id, items.name, items.year, items.group_id, groups.name as group_name,
CASE
	WHEN COUNT(authors.id) > 0
	THEN jsonb_agg(distinct jsonb_build_object('id', authors.id, 'name', authors.name))
END
as json_authors,
CASE
  WHEN COUNT(uploads.id) > 0
  THEN jsonb_agg(distinct jsonb_build_object('id', uploads.id, 'name', uploads.name, 'sum', uploads.sum, 'type', uploads.type))
END
as json_uploads,
COUNT(*) OVER()
FROM items
LEFT JOIN groups on items.group_id = groups.id
LEFT JOIN item_has_author ON items.id = item_has_author.item_id
LEFT JOIN authors         ON item_has_author.author_id = authors.id
LEFT JOIN item_has_upload ON items.id = item_has_upload.item_id
LEFT JOIN uploads         ON item_has_upload.upload_id = uploads.id
WHERE unaccent(lower(items.name)) LIKE unaccent(lower($1))
GROUP BY items.id, groups.name
ORDER BY items.year DESC
LIMIT $3 OFFSET $2;
`

func SearchItems(page int, limit int, search string) (ItemData, error) {
	rows := []schemaItem{}
	items := []Item{}

  fmt.Println(search)

	err := db.Select(&rows, searchItems, fmt.Sprintf("%%%s%%", search), calcOffset(page, limit), limit)
	if err != nil {
		return ItemData{}, err
	}

	for _, v := range rows {
		item := Item{ID: v.ID, Name: v.Name, Year: v.Year}

		json.Unmarshal([]byte(v.JsonUploads), &item.Uploads)
		json.Unmarshal([]byte(v.JsonAuthors), &item.Authors)

		if v.GroupID != nil && v.GroupName != nil {
			item.Group = Group{ID: *v.GroupID, Name: *v.GroupName}
		}

		items = append(items, item)
	}

	meta := Metadata{
		PageIndex: page,
		Limit:     limit,
	}

	if len(rows) > 0 {
		meta.TotalResults = rows[0].Count
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
