package ui

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/newdb"
	"github.com/sotirismorf/go-htmx/views"
)

type searchQueryParams struct {
	Query        string `query:"query"`
	Field        string `query:"field"`
	ItemsPerPage int    `query:"items"`
	PageIndex    int    `query:"page"`
}

type RequestHeaders struct {
	HXRequest bool `header:"HX-Request"`
}

func Search(c echo.Context) error {
	queryParams := searchQueryParams{
		ItemsPerPage: 5,
		PageIndex:    1,
	}

	err := c.Bind(&queryParams)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	requestHeaders := new(RequestHeaders)
	binder := &echo.DefaultBinder{}
	binder.BindHeaders(c, requestHeaders)

	sortOptions := []components.SelectOption{
		{Name: "Date ascending", ID: "date-asc"},
		{Name: "Date descending", ID: "date-desc"},
	}

	fieldOptions := []components.SelectOption{
		{Name: "Title", ID: "title"},
		{Name: "Author", ID: "author"},
	}

	data, err := newdb.SearchItems(queryParams.PageIndex, queryParams.ItemsPerPage, queryParams.Query)

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
	for _, v := range data.Results {
		// Many-to-many relations
		var authorsTempl []models.TemplItemResultCardAuthors
		var uploadsTempl []models.TemplItemResultCardUploads

		// Authors
		if v.Authors != nil {
			for _, k := range v.Authors {
				authorsTempl = append(authorsTempl, models.TemplItemResultCardAuthors{
					Name:       k.Name,
					AuthorLink: "/authors/" + strconv.FormatInt(k.ID, 10),
				})
			}
		}

		// Thumbnail link
		var thumbnailLink string

		// Uploads
		if len(v.Uploads) != 0 {
			thumbnailLink = "/static/thumbnails/" + v.Uploads[0].Sum + ".jpg"

			for _, k := range v.Uploads {
				uploadsTempl = append(uploadsTempl, models.TemplItemResultCardUploads{
					Type: k.Name,
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
		CurrentPage:  int64(queryParams.PageIndex),
		ItemsPerPage: int32(queryParams.ItemsPerPage),
		TotalItems:   int64(data.Metadata.TotalResults),
		TotalPages:   int64(data.Metadata.TotalPages),
		Endpoint:     "/search",
	}

	if requestHeaders.HXRequest {
		return controller.Render(c, http.StatusOK, views.SearchResults(pagination, itemsTempl))
	}

	view := views.Search(itemsTempl, pagination, sortOptions, fieldOptions)
	return controller.Render(c, http.StatusOK, views.LayoutNormal("Home", view, []components.NavItem{
		{TranslationID: "nav.home", Active: false, Href: "/"},
		{TranslationID: "nav.search", Active: true, Href: "/search"},
		{TranslationID: "nav.about", Active: false, Href: "/about"},
	}))
}

func calcPageCount(totalItems int64, pageItems int64) int64 {
	pageCount := totalItems / int64(pageItems)
	if totalItems%int64(pageItems) != 0 {
		pageCount += 1
	}

	return pageCount
}
