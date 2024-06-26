package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/schema"
	"github.com/sotirismorf/go-htmx/views"
)

type searchQueryParams struct {
	Query string `query:"query"`
	Field string `query:"field"`
	Items int32  `query:"items"`
	Page  int32  `query:"page"`
}

type RequestHeaders struct {
	HXRequest bool `header:"HX-Request"`
}

func GetSearchView(c echo.Context) error {
	queryParams := searchQueryParams{
		Items: 5,
		Page:  1,
	}

	err := c.Bind(&queryParams)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	requestHeaders := new(RequestHeaders)
	binder := &echo.DefaultBinder{}
	binder.BindHeaders(c, requestHeaders)

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

	dbParams := schema.SearchItemsParams{
		Name:   "%" + queryParams.Query + "%",
		Limit:  queryParams.Items,
		Offset: (queryParams.Page - 1) * queryParams.Items,
	}

	ctx := context.Background()
	itemsDBData, err := db.Queries.SearchItems(ctx, dbParams)
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
			ID:            strconv.FormatInt(v.ID, 10),
			Year:          strconv.Itoa(int(v.Year)),
			ThumbnailLink: thumbnailLink,
			Authors:       authorsTempl,
			Uploads:       uploadsTempl,
		})
	}

	pagination := components.TemplPagination{
		CurrentPage:  int64(queryParams.Page),
		ItemsPerPage: queryParams.Items,
	}

	// Pagination
	if len(itemsDBData) != 0 {
		pagination.TotalItems = itemsDBData[0].Count
	} else {
		pagination.TotalItems = 0
	}

	pagination.TotalPages = pagination.TotalItems / int64(queryParams.Items)
	if pagination.TotalItems%int64(queryParams.Items) != 0 {
		pagination.TotalPages += 1
	}

	if requestHeaders.HXRequest {
		return Render(c, http.StatusOK, views.SearchResults(pagination, itemsTempl))
	}

	view := views.Search(itemsTempl, pagination, sortOptions, fieldOptions)
	return Render(c, http.StatusOK, views.LayoutNormal("Home", view))
}
