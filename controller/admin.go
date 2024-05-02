package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/views"
)

func LoginHandler(c echo.Context) error {
	view := views.AdminLogin()

	return Render(c, http.StatusOK, views.BaseLayout("Home", view))
}

func AdminHandler(c echo.Context) error {
	view := views.Admin()

	return Render(c, http.StatusOK, views.BaseLayout("Admin Panel", view))
}
