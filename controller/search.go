package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/views"
)

type searchQueryParams struct {
	Query string `query:"query"`
	Field string `query:"field"`
}

type RequestHeaders struct {
	HXRequest bool `header:"HX-Request"`
}

func GetSearchView(c echo.Context) error {
	var queryParams searchQueryParams

	err := c.Bind(&queryParams)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	requestHeaders := new(RequestHeaders)
	binder := &echo.DefaultBinder{}
	binder.BindHeaders(c, requestHeaders)

	fmt.Println(requestHeaders.HXRequest)

	sortOptions := []components.SelectOption{
		components.SelectOption{
			Name: "Date ascending",
			ID:   "date-asc",
		},
		components.SelectOption{
			Name: "Date descending",
			ID:   "date-desc",
		},
	}

	fieldOptions := []components.SelectOption{
		components.SelectOption{
			Name: "Title",
			ID:   "title",
		},
		components.SelectOption{
			Name: "Author",
			ID:   "author",
		},
	}

	ctx := context.Background()
	itemsDBData, err := db.Queries.SearchItems(ctx, "%"+queryParams.Query+"%")
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

	if requestHeaders.HXRequest {
		return Render(c, http.StatusOK, views.SearchResults(itemsTempl))
	}

	view := views.Search(itemsTempl, sortOptions, fieldOptions)
	return Render(c, http.StatusOK, views.AdminLayout("Home", view))
}
