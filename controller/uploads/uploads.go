package uploads

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/views"
	"github.com/sotirismorf/go-htmx/views/admin/uploads"
)

func AdminGetUploads(c echo.Context) error {
	ctx := context.Background()

	data, err := db.Queries.SelectUploads(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	view := uploads.AdminUploads(data)

	return handlers.Render(c, http.StatusOK, views.AdminLayout("Authors", view))
}
