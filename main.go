package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/controller/authors"
	"github.com/sotirismorf/go-htmx/controller/items"
	"github.com/sotirismorf/go-htmx/db"
)

func main() {
	db.ConnectDB()

	app := echo.New()

	// app.Use(middleware.Logger())
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri}\n",
	}))

	app.GET("/", handlers.HomeHandler)

	app.GET("/admin", handlers.AdminHandler)
	app.GET("/admin/items", items.AdminItemsHandler)
	app.POST("/admin/items/create", items.AdminCreateItemHandler)
	app.GET("/admin/items/create", items.CreateItemController)
	app.GET("/admin/items/:id", items.AdminSingleItemHandler)
	app.POST("/admin/items/:id", items.AdminSingleItemHandler)
	app.DELETE("/admin/items/:id", items.AdminSingleItemDelete)
	app.GET("/htmx/admin/items/:id/edit", items.HTMXAdminItemsOneEdit)
	app.GET("/htmx/admin/items/:id/cancel", items.HTMXAdminItemsOneCancelEdit)

	app.GET("/admin/authors", authors.AdminAuthorsHandler)
	app.DELETE("/admin/authors/:id", authors.AdminSingleAuthorDelete)
	app.POST("/admin/authors", authors.Create)

	app.GET("/admin/login", handlers.LoginHandler)

	app.Logger.Fatal(app.Start(":8080"))
}
