package items

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/schema"
)

type formData struct {
	Name        string  `form:"name"`
	Description string  `form:"description"`
	Year        int16   `form:"year"`
	AuthorID    []int64 `form:"author"`
	UploadID    []int64 `form:"upload"`
}

func AdminCreateItemHandler(c echo.Context) error {
	ctx := context.Background()

	var formData formData

	err := c.Bind(&formData)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	itemParams := schema.CreateItemParams{
		Name: formData.Name,
		Year: formData.Year,
	}

	if formData.Description != "" {
		itemParams.Description = &formData.Description
	}

	createdItem, err := db.Queries.CreateItem(ctx, itemParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	for _, v := range formData.AuthorID {
		itemHasAuthorParams := schema.CreateItemHasAuthorRelationshipParams{
			ItemID:   createdItem.ID,
			AuthorID: int64(v),
		}

		_, err = db.Queries.CreateItemHasAuthorRelationship(ctx, itemHasAuthorParams)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
	}

	for _, v := range formData.UploadID {
		itemHasUploadParams := schema.CreateItemHasUploadRelationshipParams{
			ItemID:   createdItem.ID,
			UploadID: int64(v),
		}

		_, err = db.Queries.CreateItemHasUploadRelationship(ctx, itemHasUploadParams)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
	}

	return c.Redirect(http.StatusFound, "/admin/items")
}
