package uploads

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/schema"
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

	// Determine filetype
	extension := filepath.Ext(file.Filename)
	ft := mime.TypeByExtension(extension)
	var fileType schema.Filetype

	switch ft {
	case "image/jpeg":
		fileType = schema.FiletypeImageJpeg
	case "image/png":
		fileType = schema.FiletypeImagePng
	case "application/pdf":
		fileType = schema.FiletypeApplicationPdf
	default:
		return echo.NewHTTPError(http.StatusNotFound, "Filetype \""+ft+"\" is not supported")
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

	// Generate thumbnail
	magickInputFile := "uploads/" + sum
	if fileType == "application/pdf" {
		magickInputFile += "[0]"
	}
	cmd := exec.Command("magick", magickInputFile, "-thumbnail", "256x256>", "uploads/thumbnails/"+sum+".jpg")
	if err := cmd.Run(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Create database entry
	params := schema.CreateUploadParams{
		Sum:  sum,
		Name: file.Filename,
		Type: fileType,
		Size: int32(file.Size),
	}
	_, err = db.Queries.CreateUpload(context.Background(), params)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.Redirect(http.StatusFound, "/admin/uploads")
}

func DeleteUpload(c echo.Context) error {
	var param controller.ParamContainsID

	err := c.Bind(&param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	data, err := db.Queries.SelectSingleUpload(context.Background(), param.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	err = db.Queries.DeleteSingleUpload(context.Background(), param.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	err = os.Remove("uploads/" + data[0].Sum)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	err = os.Remove("uploads/thumbnails/" + data[0].Sum + ".jpg")
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.NoContent(http.StatusOK)
}
