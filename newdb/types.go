package newdb

type Metadata struct {
	PageIndex    int
	Limit        int
	TotalPages   int
	TotalResults int
}

type ItemData struct {
	Metadata Metadata
	Results  []Item
}

type Item struct {
	ID          int64
	Name        string
	Description *string
	Group       Group
	Authors     []Author
}

type Group struct {
	ID   int32
	Name string
}

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Bio  string
}

type schemaItem struct {
	ID          int64
	Name        string
	Description *string
	GroupID     *int32  `db:"group_id"`
	GroupName   *string `db:"group_name"`
	JsonAuthors []byte  `db:"json_authors"`
	Count       int     `db:"count"`
}
