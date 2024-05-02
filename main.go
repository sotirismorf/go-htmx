package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/handlers"
)

func main() {
	db.ConnectDB()

	app := echo.New()

	// app.Use(middleware.Logger())
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri}\n",
	}))

	app.GET("/", handlers.HomeHandler)

	app.GET ("/admin", handlers.AdminHandler)
	app.GET ("/admin/items", handlers.AdminItemsHandler)
	app.POST("/admin/items", handlers.AdminCreateItemHandler)
	app.GET ("/admin/authors", handlers.AdminAuthorsHandler)

  app.GET   ("/admin/items/:id", handlers.AdminSingleItemHandler)
  app.POST  ("/admin/items/:id", handlers.AdminSingleItemHandler)
  app.DELETE("/admin/items/:id", handlers.AdminSingleItemDelete)
  app.DELETE("/admin/authors/:id", handlers.AdminSingleAuthorDelete)

  app.GET ("/htmx/admin/items/:id/edit", handlers.HTMXAdminItemsOneEdit)
  app.GET ("/htmx/admin/items/:id/cancel", handlers.HTMXAdminItemsOneCancelEdit)

	app.GET("/admin/login", handlers.LoginHandler)

	app.Logger.Fatal(app.Start(":8080"))
}
