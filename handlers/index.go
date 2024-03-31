package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/views"
)

func HomeHandler(c echo.Context) error {

	ctx := context.Background()

	data, err := db.Queries.ListAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	items, err := db.Queries.ListItemsWithAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	type author struct {
		Id   int    `json:"f1"`
		Name string `json:"f2"`
	}

	authors := []author{}
	json.Unmarshal([]byte(items[0].Authors), &authors)

  view := views.Index(data)

	return Render(c, http.StatusOK, views.BaseLayout("Home", view))
}
