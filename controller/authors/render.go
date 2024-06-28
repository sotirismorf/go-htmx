package authors

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/views"
)

func AdminAuthorsHandler(c echo.Context) error {
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

func CreateAuthorForm(c echo.Context) error {
	view := components.FormCreateAuthor()

	return controller.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}
