package uploads

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"

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

func AdminCreateUpload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	// Read file
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer src.Close()

	// Generate md5 hash of file
	hash := md5.New()
	if _, err := io.Copy(hash, src); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	sum := hex.EncodeToString(hash.Sum(nil))

	// Create destination file
	dst, err := os.Create("uploads/" + sum)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer dst.Close()

	// Rewind the file to the beginning after hashing
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Copy to destination file
	if _, err = io.Copy(dst, src); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	// Create database entry
	ctx := context.Background()

	params := schema.CreateUploadParams{
		Sum:  sum,
		Name: file.Filename,
		Type: "pdf",
		Size: int32(file.Size),
	}
	_, err = db.Queries.CreateUpload(ctx, params)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.Redirect(http.StatusFound, "/admin/uploads")
}

func AdminGetUploadForm(c echo.Context) error {
	view := components.FormCreateUpload()

	return controller.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}
