package items

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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

type multiSelectDropdownSearchParams struct {
	Search       string   `form:"search"`
	NewSelection string   `form:"newSelection"`
	Selected     []string `form:"selected"`
}

func HTMXMultiSelectDropdown(c echo.Context) error {
	params := multiSelectDropdownSearchParams{}
	err := c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if params.NewSelection != "" {
		newSelection := strings.Split(params.NewSelection, ":")
		if len(newSelection) < 2 {
			return echo.NewHTTPError(http.StatusNotFound, "something went wrong")
		}

		var selected []components.TemplMultiSelectDropdownItem

		for _, v := range params.Selected {
			selectedItem := strings.Split(v, ":")
			if len(selectedItem) < 2 {
				return echo.NewHTTPError(http.StatusNotFound, "something went wrong")
			}

			selected = append(selected, components.TemplMultiSelectDropdownItem{
				ID:   selectedItem[0],
				Name: selectedItem[1],
			})
		}

		selected = append(selected, components.TemplMultiSelectDropdownItem{
			ID:   newSelection[0],
			Name: newSelection[1],
		})

		return controller.Render(c, http.StatusOK, components.MultiSelectDropdown(
			components.TemplMultiSelectDropdown{
				Name:     "author",
				Label:    "Author",
				Selected: selected,
			}))
	}

	var results []components.TemplMultiSelectDropdownItem

	if params.Search == "" {
		return controller.Render(c, http.StatusOK, components.MultiSelectDropdownResults("author", results))
	}

	dbData, err := db.Queries.SearchAuthors(context.Background(), "%"+params.Search+"%")
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	for _, v := range dbData {
		results = append(results, components.TemplMultiSelectDropdownItem{
			ID:   strconv.FormatInt(v.ID, 10),
			Name: v.Name,
		})
	}

	return controller.Render(c, http.StatusOK, components.MultiSelectDropdownResults("author", results))
}
