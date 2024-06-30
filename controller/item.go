package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/views"
)

func Item(c echo.Context) error {
	var param ParamContainsID
	err := c.Bind(&param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctx := context.Background()

	dbItem, err := db.Queries.SelectItem(ctx, param.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	dbUploads, err := db.Queries.SelectUploadsOfItemByItemID(
		context.WithValue(c.Request().Context(), 1, c),
		dbItem.ID)

	props := views.ItemPage{
		Name:    dbItem.Name,
		Uploads: []views.Upload{},
	}

	templUploads := []views.Upload{}
	for _, v := range dbUploads {
		templUploads = append(templUploads, views.Upload{
			ID:       strconv.FormatInt(v.ID, 10),
			Name:     v.Name,
			Size:     models.PrettyByteSize(v.Size),
			Filetype: string(v.Type),
		})
	}

	if len(dbUploads) > 0 {
		props.Uploads = templUploads
		props.ThumbnailLink = fmt.Sprintf("/static/thumbnails/%s.jpg", dbUploads[0].Sum)
	}

	if dbItem.Description != nil {
		props.Description = *dbItem.Description
	}

	view := views.Item(props)

	return Render(c, http.StatusOK, views.LayoutNormal("Home", view))
}
