package uploads

import (
	"context"
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/views"
	"github.com/sotirismorf/go-htmx/views/admin/uploads"
	"github.com/sotirismorf/go-htmx/models"
)

func prettyByteSize(bytes int32) string {
	bytesFloat := float64(bytes)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bytesFloat) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", bytesFloat, unit)
		}
		bytesFloat /= 1024.0
	}
	return fmt.Sprintf("%.1fYiB", bytesFloat)
}

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
			Size: prettyByteSize(v.Size),
			Type: string(v.Type),
		})
		fmt.Println(prettyByteSize(v.Size))
	}

	view := uploads.AdminUploads(templateData)

	return controller.Render(c, http.StatusOK, views.AdminLayout("Uploads", view))
}
