package ui

import (
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/models"
	"github.com/sotirismorf/go-htmx/newdb"
	"github.com/sotirismorf/go-htmx/views"
	"github.com/sotirismorf/go-htmx/views/admin/items"
)

func AdminSingleItem(c echo.Context) error {
	var param controller.ParamContainsID

	err := c.Bind(&param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	data, err := newdb.SelectItem(param.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	uploads := []models.UploadTemplateData{}
	for _, v := range data.Uploads {
		uploads = append(uploads, models.UploadTemplateData{
			ID:   v.ID,
			Name: v.Name,
			Sum:  v.Sum,
			Type: string(v.Type),
		})
	}

	view := items.AdminSingle([]templ.Component{
		items.AdminSingleText("ID", strconv.Itoa(int(data.ID))),
		items.AdminSingleText("Name", data.Name),
		items.AdminSingleText("Description", data.Description),
	})

	return controller.Render(c, http.StatusOK, views.AdminLayout(data.Name, view))
}

// func HTMXAdminItemsOneEdit(c echo.Context) error {
// 	var param controller.ParamContainsID
//
// 	err := c.Bind(&param)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err)
// 	}
//
// 	i, err := SingleItemPopulated(param.ID)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, err)
// 	}
//
// 	view := items.SingleItemAttributesEdit(i)
//
// 	return controller.Render(c, http.StatusOK, view)
// }

//
// func HTMXAdminItemsOneCancelEdit(c echo.Context) error {
// 	var param controller.ParamContainsID
//
// 	err := c.Bind(&param)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err)
// 	}
//
// 	i, err := SingleItemPopulated(param.ID)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, err)
// 	}
//
// 	view := items.SingleItemAttributes(i)
//
// 	return controller.Render(c, http.StatusOK, view)
// }
