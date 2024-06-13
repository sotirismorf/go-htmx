package uploads

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
)

func GetUpload(c echo.Context) error {

	var param controller.ParamContainsID
	err := c.Bind(&param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	data, err := db.Queries.SelectSingleUpload(context.Background(), param.ID)
	if err != nil || len(data) != 1 {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.Attachment("uploads/"+data[0].Sum, data[0].Name)
}
