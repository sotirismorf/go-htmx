package items

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/views"
	"github.com/sotirismorf/go-htmx/views/admin/items"
)

func SingleItemPopulated(id int64) (models.ItemData, error) {
	ctx := context.Background()

	i := models.ItemData{}

	item, err := db.Queries.SelectSingleItemWithAuthors(ctx, id)
	if err != nil {
		return i, err
	}

	i.Id = item.ID
	i.Name = item.Name
	if item.Description != nil {
		i.Description = item.Description
	}
	if item.Authors != nil {
		authors := []models.Author{}
		json.Unmarshal([]byte(item.Authors), &authors)
		i.Authors = authors
	}

	return i, nil
}

func AdminGetSingleItem(c echo.Context) error {
	var param handlers.ParamContainsID

	err := c.Bind(&param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	i, err := SingleItemPopulated(param.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	view := items.AdminSingleItem(i)

	return handlers.Render(c, http.StatusOK, views.AdminLayout(i.Name, view))
}

func AdminDeleteSingleItem(c echo.Context) error {
	var param handlers.ParamContainsID

	err := c.Bind(&param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = db.Queries.DeleteItem(context.Background(), param.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.NoContent(http.StatusOK)
}

func HTMXAdminItemsOneEdit(c echo.Context) error {
	var param handlers.ParamContainsID

	err := c.Bind(&param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	i, err := SingleItemPopulated(param.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	view := items.SingleItemAttributesEdit(i)

	return handlers.Render(c, http.StatusOK, view)
}

func HTMXAdminItemsOneCancelEdit(c echo.Context) error {
	var param handlers.ParamContainsID

	err := c.Bind(&param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	i, err := SingleItemPopulated(param.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	view := items.SingleItemAttributes(i)

	return handlers.Render(c, http.StatusOK, view)
}
