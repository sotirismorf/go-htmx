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

func calcOffset(page int, limit int) int {
	return limit * (page - 1)
}

func calcPageCount(results int, limit int) int {
	pageCount := results / limit
	if results%limit != 0 {
		pageCount += 1
	}

	return pageCount
}

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
    Results: items,
    Metadata: meta,
  }

	return data, nil
}
