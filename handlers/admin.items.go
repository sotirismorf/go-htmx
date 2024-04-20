package handlers

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

func AdminSingleItemHandler(c echo.Context) error {
	paramItemId := c.Param("id")

	id, err := strconv.ParseInt(paramItemId, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	i, err := SingleItemPopulated(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	view := views.AdminSingleItem(i)

	return Render(c, http.StatusOK, views.BaseLayout(i.Name, view))
}

func AdminSingleItemDelete(c echo.Context) error {
	paramItemId := c.Param("id")

	id, err := strconv.ParseInt(paramItemId, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	err = db.Queries.DeleteItem(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.NoContent(http.StatusOK)
}

func HTMXAdminItemsOneEdit(c echo.Context) error {
	paramItemId := c.Param("id")

	id, err := strconv.ParseInt(paramItemId, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	i, err := SingleItemPopulated(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	view := views.SingleItemAttributesEdit(i)

	return Render(c, http.StatusOK, view)
}

func HTMXAdminItemsOneCancelEdit(c echo.Context) error {
	paramItemId := c.Param("id")

	id, err := strconv.ParseInt(paramItemId, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	i, err := SingleItemPopulated(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	view := views.SingleItemAttributes(i)

	return Render(c, http.StatusOK, view)
}
