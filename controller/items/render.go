package items

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/schema"
	"github.com/sotirismorf/go-htmx/views"
	"github.com/sotirismorf/go-htmx/views/admin/items"
)

type searchParams struct {
	Limit  int32 `query:"limit"`
	Offset int32 `query:"offset"`
}

func AdminItemsHandler(c echo.Context) error {
	params := searchParams{
		Limit:  10,
		Offset: 0,
	}

	err := c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	fmt.Println(params)

	ctx := context.Background()

	itemData, err := db.Queries.SelectItemsWithAuthorsAndUploads(ctx, schema.SelectItemsWithAuthorsAndUploadsParams{
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	authors, err := db.Queries.SelectAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	itemsGenerated := []models.ItemData{}

	for _, i := range itemData {
		item := models.ItemData{}

		item.Id = i.ID
		item.Name = i.Name
		if i.Description != nil {
			item.Description = i.Description
		}
		if i.Authors != nil {
			authors := []models.Author{}
			json.Unmarshal([]byte(i.Authors), &authors)
			item.Authors = authors
		}
		itemsGenerated = append(itemsGenerated, item)
	}

	view := items.AdminItems(itemsGenerated, authors)

	return controller.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}

func CreateItemController(c echo.Context) error {
	ctx := context.Background()

	authorData, err := db.Queries.SelectAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	authorOptions := []components.SelectOption{}

	for _, v := range authorData {
		authorOptions = append(authorOptions, components.SelectOption{ID: strconv.FormatInt(v.ID, 10), Name: v.Name})
	}

	uploadData, err := db.Queries.SelectUploads(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	uploadOptions := []components.SelectOption{}

	for _, v := range uploadData {
		uploadOptions = append(uploadOptions, components.SelectOption{ID: strconv.FormatInt(v.ID, 10), Name: v.Name})
	}

	view := components.FormCreateItem(authorOptions, uploadOptions)

	return controller.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}
