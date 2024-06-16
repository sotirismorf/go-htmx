package controller

import (
	"context"
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/views"
)

func GetSearchView(c echo.Context) error {

	ctx := context.Background()

	itemsDBData, err := db.Queries.SelectItemsWithAuthorsAndUploads(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	type author struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}

	type upload struct {
		Id       int64  `json:"id"`
		Filename string `json:"filename"`
		Sum      string `json:"sum"`
	}

	uploads := []upload{}
	json.Unmarshal([]byte(itemsDBData[0].Uploads), &uploads)

	itemsTempl := []models.TemplItemResultCard{}
	for _, v := range itemsDBData {
		// Many-to-many relations
		var authorsTempl []models.TemplItemResultCardAuthors
		var uploadsTempl []models.TemplItemResultCardUploads

		// Authors
		if v.Authors != nil {
			authorsJson := []author{}
			json.Unmarshal([]byte(v.Authors), &authorsJson)

			for _, k := range authorsJson {
				authorsTempl = append(authorsTempl, models.TemplItemResultCardAuthors{
					Name:       k.Name,
					AuthorLink: "/authors/" + strconv.FormatInt(k.Id, 10),
				})
			}
		}

		// Thumbnail link
		var thumbnailLink string

		// Uploads
		if v.Uploads != nil {
			uploadsJson := []upload{}
			json.Unmarshal([]byte(v.Uploads), &uploadsJson)

			thumbnailLink = "/static/thumbnails/" + uploadsJson[0].Sum + ".jpg"

			for _, k := range uploadsJson {
				uploadsTempl = append(uploadsTempl, models.TemplItemResultCardUploads{
					Type: k.Filename,
				})
			}
		}

		itemsTempl = append(itemsTempl, models.TemplItemResultCard{
			Name:          v.Name,
			Year:          strconv.Itoa(int(v.Year)),
			ThumbnailLink: thumbnailLink,
			Authors:       authorsTempl,
			Uploads:       uploadsTempl,
		})
	}

	view := views.Search(itemsTempl)

	return Render(c, http.StatusOK, views.AdminLayout("Home", view))
}
