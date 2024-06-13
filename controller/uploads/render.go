package uploads

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/views"
	"github.com/sotirismorf/go-htmx/views/admin/uploads"
)

func AdminGetUploads(c echo.Context) error {
	ctx := context.Background()

	data, err := db.Queries.SelectUploads(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	templateData := []models.UploadTemplateData{}
	for _, v := range data {
		templateData = append(templateData, models.UploadTemplateData{
			ID:   v.ID,
			Name: v.Name,
			Sum:  v.Sum,
			Size: models.PrettyByteSize(v.Size),
			Type: string(v.Type),
		})
	}

	view := uploads.AdminUploads(templateData)

	return controller.Render(c, http.StatusOK, views.AdminLayout("Uploads", view))
}

func AdminGetUploadForm(c echo.Context) error {
	view := components.FormCreateUpload()

	return controller.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}
