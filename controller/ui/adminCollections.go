package ui

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/newdb"
	"github.com/sotirismorf/go-htmx/views"
	"github.com/sotirismorf/go-htmx/views/admin/uploads"
)

func AdminItems(c echo.Context) error {
	queryParams := searchQueryParams{
		ItemsPerPage: 20,
		PageIndex:    1,
	}

	err := c.Bind(&queryParams)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	data, err := newdb.SelectItemsPopulated(queryParams.PageIndex, queryParams.ItemsPerPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	rows := []components.TemplAdminTableRow{}

	for _, i := range data.Results {

		// authors
		var templAuthors []components.TemplAdminTableCell
		for _, v := range i.Authors {
			templAuthors = append(templAuthors, components.TemplAdminTableCell{
				Text:       v.Name,
				IsRelation: true,
				CTName:     "authors",
				ID:         strconv.FormatInt(v.ID, 10),
			})
		}

		rows = append(rows, components.TemplAdminTableRow{
			ID: strconv.FormatInt(i.ID, 10),
			Cells: [][]components.TemplAdminTableCell{
				{{Text: i.Name}},
				{{Text: i.Group.Name, IsRelation: true, CTName: "groups", ID: strconv.Itoa(int(i.Group.ID))}},
				templAuthors,
			},
		})
	}

	view := components.AdminPage(components.TemplAdminPage{
		Title:       "Items",
		CTName:      "items",
		CanDelete:   true,
		CanDownload: false,
		Columns:     []string{"Name", "Group", "Authors"},
		Rows:        rows,
		Pagination: components.TemplPagination{
			CurrentPage: int64(queryParams.PageIndex),
			TotalPages:  int64(data.Metadata.TotalPages),
			Endpoint:    "/admin/items",
		},
	})

	requestHeaders := new(RequestHeaders)
	binder := &echo.DefaultBinder{}
	binder.BindHeaders(c, requestHeaders)
	if requestHeaders.HXRequest {
		return controller.Render(c, http.StatusOK, view)
	}

	return controller.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}

func AdminAuthors(c echo.Context) error {
	qp := searchQueryParams{
		ItemsPerPage: 20,
		PageIndex:    1,
	}

	err := c.Bind(&qp)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	data, err := newdb.SelectAuthors(qp.PageIndex, qp.ItemsPerPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	rows := []components.TemplAdminTableRow{}
	for _, i := range data.Results {
		rows = append(rows, components.TemplAdminTableRow{
			ID: strconv.FormatInt(i.ID, 10),
			Cells: [][]components.TemplAdminTableCell{
				{{Text: strconv.FormatInt(i.ID, 10)}},
				{{Text: i.Name}},
				{{Text: strconv.Itoa(i.ItemCount)}},
			}})
	}

	view := components.AdminPage(components.TemplAdminPage{
		Title:       "Authors",
		CTName:      "authors",
		CanDelete:   true,
		CanDownload: false,
		Columns:     []string{"ID", "Name", "Items"},
		Rows:        rows,
		Pagination: components.TemplPagination{
			CurrentPage: int64(qp.PageIndex),
			TotalPages:  int64(data.Metadata.TotalPages),
			Endpoint:    "/admin/items",
		},
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

func AdminGroups(c echo.Context) error {
	ctx := context.Background()

	dbData, err := db.Queries.SelectGroups(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	rows := []components.TemplAdminTableRow{}
	for _, i := range dbData {
		rows = append(rows, components.TemplAdminTableRow{
			ID: strconv.FormatInt(int64(i.ID), 10),
			Cells: [][]components.TemplAdminTableCell{
				{{Text: strconv.FormatInt(int64(i.ID), 10)}},
				{{Text: i.Name}},
				{{Text: *i.City}},
			}})
	}

	view := components.AdminPage(components.TemplAdminPage{
		Title:       "Groups",
		CTName:      "groups",
		CanDelete:   true,
		CanDownload: false,
		Columns:     []string{"ID", "Name", "City"},
		Rows:        rows,
	})

	return controller.Render(c, http.StatusOK, views.AdminLayout("Groups", view))
}
