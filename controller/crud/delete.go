package crud

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
)

func DeleteItem(c echo.Context) error {
	var param controller.ParamContainsID

	err := c.Bind(&param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = db.Queries.DeleteItem(context.Background(), param.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.NoContent(http.StatusOK)
}

func DeleteAuthor(c echo.Context) error {
	var param controller.ParamContainsID

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
