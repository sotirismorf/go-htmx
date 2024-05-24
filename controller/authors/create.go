package authors

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/schema"
	"github.com/sotirismorf/go-htmx/views"
)

type formData struct {
	Name string `form:"name"`
	Bio  string `form:"bio"`
}

func Create(c echo.Context) error {
	ctx := context.Background()

	var formData formData

	err := c.Bind(&formData)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	createAuthorParams := schema.CreateAuthorParams{Name: formData.Name}

	if formData.Bio != "" {
		createAuthorParams.Bio = &formData.Bio
	}

	_, err = db.Queries.CreateAuthor(ctx, createAuthorParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.Redirect(http.StatusFound, "/admin/authors")
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
