package newdb

const selectUsers = `
SELECT authors.id, authors.name, count(item_has_author.item_id) as item_count,
COUNT(*) OVER()
FROM authors
LEFT JOIN item_has_author ON authors.id = item_has_author.author_id
GROUP BY authors.id
ORDER BY authors.id
LIMIT $2 OFFSET $1;
`

func SelectUsers(page int, limit int) (AuthorData, error) {
	rows := []schemaAuthor{}
	authors := []Author{}

	err := db.Select(&rows, selectAuthors, calcOffset(page, limit), limit)
	if err != nil {
		return AuthorData{}, err
	}

	for _, v := range rows {
		authors = append(authors, Author{
			ID:        v.ID,
			Name:      v.Name,
			ItemCount: v.ItemCount,
		})
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

	data := AuthorData{
		Results:  authors,
		Metadata: meta,
	}

	return data, nil
}
