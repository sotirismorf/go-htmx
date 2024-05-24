package authors

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/controller"
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

	return handlers.Render(c, http.StatusOK, views.AdminLayout("Authors", view))
}

func AdminSingleAuthorDelete(c echo.Context) error {
	var param handlers.ParamContainsID

	err := c.Bind(&param)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err = db.Queries.DeleteAuthor(context.Background(), param.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.NoContent(http.StatusOK)
}

func CreateAuthorForm(c echo.Context) error {
	ctx := context.Background()

	authors, err := db.Queries.SelectAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	options := []components.SelectOption{}

	for _, v := range authors {
		options = append(options, components.SelectOption{ID: v.ID, Name: v.Name})
	}

	view := components.FormCreateAuthor()

	return handlers.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}
