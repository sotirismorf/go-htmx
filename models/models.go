package models

type Author struct {
	Id   int64  `json:"f1"`
	Name string `json:"f2"`
}

type ItemData struct {
	Id          int64
	Name        string
	Description *string
	Authors     []Author
}

type UploadTemplateData struct {
	ID   int64
	Name string
	Size string
	Type string
}
