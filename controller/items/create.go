package items

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/schema"
)

type formData struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	AuthorID    int64  `form:"author"`
	UploadID    int64  `form:"upload"`
}

func AdminCreateItemHandler(c echo.Context) error {
	ctx := context.Background()

	var formData formData

	err := c.Bind(&formData)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	itemParams := schema.CreateItemParams{Name: formData.Name}

	if formData.Description != "" {
		itemParams.Description = &formData.Description
	}

	createdItem, err := db.Queries.CreateItem(ctx, itemParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if formData.AuthorID != 0 {
		itemHasAuthorParams := schema.CreateItemHasAuthorRelationshipParams{
			ItemID:   createdItem.ID,
			AuthorID: int64(formData.AuthorID),
		}

		_, err = db.Queries.CreateItemHasAuthorRelationship(ctx, itemHasAuthorParams)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
	}

	if formData.UploadID != 0 {
		itemHasUploadParams := schema.CreateItemHasUploadRelationshipParams{
			ItemID:   createdItem.ID,
			UploadID: int64(formData.UploadID),
		}

		_, err = db.Queries.CreateItemHasUploadRelationship(ctx, itemHasUploadParams)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
	}

	return c.Redirect(http.StatusFound, "/admin/items")
}
