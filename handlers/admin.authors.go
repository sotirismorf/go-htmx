package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/views"
	"github.com/sotirismorf/go-htmx/views/admin/authors"
)

func AdminAuthorsHandler(c echo.Context) error {
	ctx := context.Background()

	authorData, err := db.Queries.SelectAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	view := authors.AdminAuthors(authorData)

	return Render(c, http.StatusOK, views.BaseLayout("Admin Panel / Authors", view))
}

type ParamContainsID struct {
	ID int64 `param:"id"`
}

func AdminSingleAuthorDelete(c echo.Context) error {
	var param ParamContainsID

	err := c.Bind(&param); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err = db.Queries.DeleteAuthor(context.Background(), param.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.NoContent(http.StatusOK)
}
