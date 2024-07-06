package newdb

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

