package models

import (
	"fmt"
	"math"
)

type Author struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
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
	Sum  string
	Type string
}

type TemplItemResultCard struct {
	Name          string
	ID            string
	Year          string
	ThumbnailLink string
	Authors       []TemplItemResultCardAuthors
	Uploads       []TemplItemResultCardUploads
}

type TemplItemResultCardAuthors struct {
	Name       string
	AuthorLink string
}

type TemplItemResultCardUploads struct {
	Type string
}

func PrettyByteSize(bytes int32) string {
	bytesFloat := float64(bytes)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bytesFloat) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", bytesFloat, unit)
		}
		bytesFloat /= 1024.0
	}
	return fmt.Sprintf("%.1fYiB", bytesFloat)
}
