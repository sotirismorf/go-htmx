package ui

import (
	"context"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/views"
)

func AdminCreateItemForm(c echo.Context) error {
	ctx := context.Background()

	authorData, err := db.Queries.SelectAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	authorOptions := []components.SelectOption{}

	for _, v := range authorData {
		authorOptions = append(authorOptions, components.SelectOption{ID: strconv.FormatInt(v.ID, 10), Name: v.Name})
	}

	uploadData, err := db.Queries.SelectUploads(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	uploadOptions := []components.SelectOption{}

	for _, v := range uploadData {
		uploadOptions = append(uploadOptions, components.SelectOption{ID: strconv.FormatInt(v.ID, 10), Name: v.Name})
	}

	// view := components.FormCreateItem(authorOptions, uploadOptions)
	view := components.AdminCreateForm(components.CreateForm{
		CTName: "items",
		Inputs: []templ.Component{
			components.DahliaInput(components.TemplInput{
				Name:     "name",
				Label:    "Name",
				Type:     "text",
				Required: true,
			}),
			components.DahliaInput(components.TemplInput{
				Name:  "description",
				Label: "Description",
				Type:  "text",
			}),
			components.DahliaInput(components.TemplInput{
				Name:     "year",
				Label:    "Year Published",
				Type:     "number",
				Required: true,
			}),
			components.MultiSelectDropdown(components.TemplMultiSelectDropdown{
				Name:     "author",
				Label:    "Author",
				Selected: []components.TemplMultiSelectDropdownItem{},
			}),
			components.MultiSelectDropdown(components.TemplMultiSelectDropdown{
				Name:     "upload",
				Label:    "Upload",
				Selected: []components.TemplMultiSelectDropdownItem{},
			}),
		},
	})

	return controller.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}

func AdminCreateAuthorForm(c echo.Context) error {
	view := components.AdminCreateForm(components.CreateForm{
		CTName: "authors",
		Inputs: []templ.Component{
			components.DahliaInput(components.TemplInput{
				Name:     "name",
				Label:    "Name",
				Type:     "text",
				Required: true,
			}),
			components.DahliaInput(components.TemplInput{
				Name:  "bio",
				Label: "Biography",
				Type:  "text",
			}),
		},
	})

	return controller.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}

func AdminCreateGroupForm(c echo.Context) error {

	dbCities, err := db.Queries.SelectPlaces(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	templPlaces := []components.SelectOption{}
	for _, v := range dbCities {
		templPlaces = append(templPlaces, components.SelectOption{
			Name: v.Name,
			ID:   strconv.FormatInt(int64(v.ID), 10),
		})
	}

	view := components.AdminCreateForm(components.CreateForm{
		CTName: "groups",
		Inputs: []templ.Component{
			components.DahliaInput(components.TemplInput{
				Label: "Name",
				Name:  "name",
			}),
			components.Select(components.TemplSelect{
				Name:       "location",
				Label:      "City",
				IsDisabled: false,
				Options:    templPlaces,
			}),
		},
	})

	return controller.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}

func AdminCreateUploadForm(c echo.Context) error {
	view := components.FormCreateUpload()

	return controller.Render(c, http.StatusOK, views.AdminLayout("Admin Panel - Items", view))
}
