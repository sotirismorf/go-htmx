package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/views"
)

func LoginHandler(c echo.Context) error {
	view := views.AdminLogin()

	return Render(c, http.StatusOK, views.LayoutNormal("Home", view, []components.NavItem{
		{TranslationID: "nav.home", Active: false, Href: "/"},
		{TranslationID: "nav.search", Active: false, Href: "/search"},
		{TranslationID: "nav.about", Active: false, Href: "/about"},
	}))
}

func AdminHandler(c echo.Context) error {
	view := views.Admin()

	return Render(c, http.StatusOK, views.AdminLayout("Admin Panel", view))
}
