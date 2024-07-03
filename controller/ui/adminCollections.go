package ui

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/schema"
	"github.com/sotirismorf/go-htmx/views"
	"github.com/sotirismorf/go-htmx/views/admin/uploads"
)

type searchParams struct {
	Limit  int32 `query:"limit"`
	Offset int32 `query:"offset"`
}

func AdminItems(c echo.Context) error {
	params := searchParams{
		Limit:  10,
		Offset: 0,
	}

	err := c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	ctx := context.Background()

	dbAuthorRows, err := db.Queries.SelectItemsWithAuthorsAndUploads(ctx, schema.SelectItemsWithAuthorsAndUploadsParams{
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	rows := []components.TemplAdminTableRow{}

	for _, i := range dbAuthorRows {

		// authors
		templAuthors := []components.TemplAdminTableCell{}
		if i.Authors != nil {
			authors := []models.Author{}
			json.Unmarshal([]byte(i.Authors), &authors)
			for _, v := range authors {
				templAuthors = append(templAuthors, components.TemplAdminTableCell{
					Text:       v.Name,
					IsRelation: true,
					CTName:     "authors",
					ID:         strconv.FormatInt(v.ID, 10),
				})
			}
		}

		rows = append(rows, components.TemplAdminTableRow{
			ID: strconv.FormatInt(i.ID, 10),
			Cells: [][]components.TemplAdminTableCell{
				{{Text: strconv.FormatInt(i.ID, 10)}},
				{{Text: i.Name}},
				templAuthors,
			},
		})
	}

	view := components.AdminPage(components.TemplAdminPage{
		Title:       "Items",
		CTName:      "items",
		CanDelete:   true,
		CanDownload: false,
		Columns:     []string{"ID", "Name", "Authors"},
		Rows:        rows,
	})

	return controller.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}

func AdminAuthors(c echo.Context) error {
	ctx := context.Background()

	authorData, err := db.Queries.SelectAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	rows := []components.TemplAdminTableRow{}
	for _, i := range authorData {
		rows = append(rows, components.TemplAdminTableRow{
			ID: strconv.FormatInt(i.ID, 10),
			Cells: [][]components.TemplAdminTableCell{
				{{Text: strconv.FormatInt(i.ID, 10)}},
				{{Text: i.Name}},
				{{Text: "example bio"}},
				{{Text: func() string {
					if i.Bio != nil {
						return *i.Bio
					} else {
						return ""
					}
				}()}},
			}})
	}

	view := components.AdminPage(components.TemplAdminPage{
		Title:       "Authors",
		CTName:      "authors",
		CanDelete:   true,
		CanDownload: false,
		Columns:     []string{"ID", "Name", "Bio"},
		Rows:        rows,
	})

	return controller.Render(c, http.StatusOK, views.AdminLayout("Authors", view))
}

func AdminUploads(c echo.Context) error {
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
