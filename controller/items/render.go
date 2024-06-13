package items

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/views"
	"github.com/sotirismorf/go-htmx/views/admin/items"
)

func AdminItemsHandler(c echo.Context) error {
	ctx := context.Background()

	itemData, err := db.Queries.SelectItemsWithAuthors(ctx)
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
		authorOptions = append(authorOptions, components.SelectOption{ID: v.ID, Name: v.Name})
	}

	uploadData, err := db.Queries.SelectUploads(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	uploadOptions := []components.SelectOption{}

	for _, v := range uploadData {
		uploadOptions = append(uploadOptions, components.SelectOption{ID: v.ID, Name: v.Name})
	}

	view := components.FormCreateItem(authorOptions, uploadOptions)

	return controller.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}
