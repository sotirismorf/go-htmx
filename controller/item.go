package controller

import (
	"context"
	"encoding/json"
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

	dbItem, err := db.Queries.SelectItemPopulated(ctx, param.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	props := views.ItemPage{
		Name:    dbItem.Name,
		Authors: []views.Author{},
		Uploads: []views.Upload{},
	}

	// Authors
	if dbItem.Authors != nil {
		authors := []models.Author{}
		json.Unmarshal([]byte(dbItem.Authors), &authors)

		for _, v := range authors {
			props.Authors = append(props.Authors, views.Author{
				ID:   strconv.FormatInt(v.ID, 10),
				Name: v.Name,
			})
		}
	}

	// Uploads
	if dbItem.Uploads != nil {
		uploads := []models.Upload{}
		json.Unmarshal([]byte(dbItem.Uploads), &uploads)

		for _, v := range uploads {
			props.Uploads = append(props.Uploads, views.Upload{
				ID:   strconv.FormatInt(v.ID, 10),
				Name: v.Filename,
			})
		}
    props.ThumbnailLink = fmt.Sprintf("/static/thumbnails/%s.jpg", uploads[0].Sum)
	}

	if dbItem.Description != nil {
		props.Description = *dbItem.Description
	}

	return Render(c, http.StatusOK, views.LayoutNormal("Home", views.Item(props)))
}
