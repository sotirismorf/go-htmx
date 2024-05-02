package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/views"
)

func AdminAuthorsHandler(c echo.Context) error {
	ctx := context.Background()

	authors, err := db.Queries.SelectAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	view := views.AdminAuthors(authors)

	return Render(c, http.StatusOK, views.BaseLayout("Admin Panel / Authors", view))
}

func AdminSingleAuthorDelete(c echo.Context) error {
	paramItemId := c.Param("id")

	id, err := strconv.ParseInt(paramItemId, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	err = db.Queries.DeleteAuthor(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.NoContent(http.StatusOK)
}
