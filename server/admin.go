package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/schema"
	"github.com/sotirismorf/go-htmx/views"
	"github.com/sotirismorf/go-htmx/views/admin/items"
)


func LoginHandler(c echo.Context) error {
  view := views.AdminLogin()

	return Render(c, http.StatusOK, views.BaseLayout("Home", view))
}

func AdminHandler(c echo.Context) error {
  view := views.Admin()

	return Render(c, http.StatusOK, views.BaseLayout("Admin Panel", view))
}

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

	return Render(c, http.StatusOK, views.BaseLayout("Admin Panel - Items", view))
}

func AdminCreateItemHandler(c echo.Context) error {
	ctx := context.Background()

	formValueName := c.FormValue("name")
	formValueDescription := c.FormValue("description")
	formValueAuthorId := c.FormValue("author")

	itemParams := schema.CreateItemParams{Name: formValueName}

	if formValueDescription != "" {
		itemParams.Description = &formValueDescription
	}

	createdItem, err := db.Queries.CreateItem(ctx, itemParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if formValueAuthorId != "" {
		authorId, err := strconv.Atoi(formValueAuthorId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}

		itemHasAuthorParams := schema.CreateItemHasAuthorRelationshipParams{
			ItemID:   createdItem.ID,
			AuthorID: int64(authorId),
		}
		_, err = db.Queries.CreateItemHasAuthorRelationship(ctx, itemHasAuthorParams)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
	}

	return c.Redirect(http.StatusFound, "/admin/items")
}