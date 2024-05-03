package authors

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/schema"
)

type formData struct {
	Name     string `form:"name"`
	Bio      string `form:"bio"`
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
