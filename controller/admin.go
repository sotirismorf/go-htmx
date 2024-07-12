package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/views"
)

func LoginHandler(c echo.Context) error {
	view := views.AdminLogin()

	return Render(c, http.StatusOK, views.LayoutNormal("Home", view))
}

func AdminHandler(c echo.Context) error {
	view := views.Admin()

	return Render(c, http.StatusOK, views.AdminLayout("Admin Panel", view))
}
