package items

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
)

type multiSelectDropdownSearchParams struct {
	SearchAuthor      string   `form:"search_author"`
	SearchUpload      string   `form:"search_upload"`
	NewSelectedAuthor string   `form:"new_selected_author"`
	SelectedAuthors   []string `form:"selected_author"`
	NewSelectedUpload string   `form:"new_selected_upload"`
	SelectedUploads   []string `form:"selected_upload"`
}

// this function is an atrocity but somehow it works, I should do better
// maybe I should split this into two controllers.
// the first (handling new selections) doesn't even need a db connection
// the second one (responding with search results) perhaps could use an interface
func HTMXMultiSelectDropdown(c echo.Context) error {
	params := multiSelectDropdownSearchParams{}
	err := c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	newSelected := ""
	var selected []string

	templDropdown := components.TemplMultiSelectDropdown{}

	if params.NewSelectedAuthor != "" {
		newSelected = params.NewSelectedAuthor
		selected = params.SelectedAuthors
		templDropdown.Name = "author"
		templDropdown.Label = "Author"
	} else if params.NewSelectedUpload != "" {
		newSelected = params.NewSelectedUpload
		selected = params.SelectedUploads
		templDropdown.Name = "upload"
		templDropdown.Label = "Upload"
	}

	if newSelected != "" {
		newSelection := strings.Split(newSelected, ":")
		if len(newSelection) < 2 {
			return echo.NewHTTPError(http.StatusNotFound, "something went wrong")
		}

		var templSelectedData []components.TemplMultiSelectDropdownItem

		for _, v := range selected {
			selectedItem := strings.Split(v, ":")
			if len(selectedItem) < 2 {
				return echo.NewHTTPError(http.StatusNotFound, "something went wrong")
			}

			templSelectedData = append(templSelectedData, components.TemplMultiSelectDropdownItem{
				ID:   selectedItem[0],
				Name: selectedItem[1],
			})
		}

		templSelectedData = append(templSelectedData, components.TemplMultiSelectDropdownItem{
			ID:   newSelection[0],
			Name: newSelection[1],
		})

		templDropdown.Selected = templSelectedData
		return controller.Render(c, http.StatusOK, components.MultiSelectDropdown(templDropdown))
	}

	var results []components.TemplMultiSelectDropdownItem

	if params.SearchAuthor != "" {
		dbData, err := db.Queries.SearchAuthors(context.Background(), "%"+params.SearchAuthor+"%")
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}

		for _, v := range dbData {
			results = append(results, components.TemplMultiSelectDropdownItem{
				ID:   strconv.FormatInt(v.ID, 10),
				Name: v.Name,
			})
		}
		return controller.Render(c, http.StatusOK, components.MultiSelectDropdownResults("author", results))

	} else if params.SearchUpload != "" {
		dbData, err := db.Queries.SearchUploads(context.Background(), "%"+params.SearchUpload+"%")
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}

		for _, v := range dbData {
			results = append(results, components.TemplMultiSelectDropdownItem{
				ID:   strconv.FormatInt(v.ID, 10),
				Name: v.Name,
			})
		}
		return controller.Render(c, http.StatusOK, components.MultiSelectDropdownResults("upload", results))
	}

	keyvals, err := c.FormParams()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if keyvals["search_author"] != nil {
		return controller.Render(c, http.StatusOK, components.MultiSelectDropdownResults("author", results))
	} else if keyvals["search_upload"] != nil {
		return controller.Render(c, http.StatusOK, components.MultiSelectDropdownResults("upload", results))
	}
	return echo.NewHTTPError(http.StatusNotFound, "reached end of function... oh well")
}
