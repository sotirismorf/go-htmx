package newdb

type Item struct {
	ID          int64
	Name        string
	Description string
	Group       Group
	Authors     []Author
	Uploads     []Upload
}

type Upload struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Sum  string `json:"sum"`
	Type string `json:"type"`
	Size string `json:"size"`
}

type Group struct {
	ID   int32
	Name string
}

type Author struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Bio       string
	ItemCount int
}

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

type AuthorData struct {
	Metadata Metadata
	Results  []Author
}

type schemaItem struct {
	ID          int64
	Name        string
	Description *string
	GroupID     *int32  `db:"group_id"`
	GroupName   *string `db:"group_name"`
	JsonAuthors []byte  `db:"json_authors"`
	JsonUploads []byte  `db:"json_uploads"`
	Count       int     `db:"count"`
}

type schemaAuthor struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Bio       string
	ItemCount int `db:"item_count"`
	Count     int `db:"count"`
}
