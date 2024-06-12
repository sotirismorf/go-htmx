package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/controller/authors"
	"github.com/sotirismorf/go-htmx/controller/items"
	"github.com/sotirismorf/go-htmx/controller/uploads"
	"github.com/sotirismorf/go-htmx/db"
)

func main() {
	db.ConnectDB()

	app := echo.New()

	// app.Use(middleware.Logger())
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri}\n",
	}))

	app.GET("/", controller.HomeHandler)

	app.GET("/downloads/:id", uploads.GetUpload)
	app.GET("/static/:id", uploads.GetFileAsInline)

	app.GET("/admin", controller.AdminHandler)
	app.GET("/admin/items", items.AdminItemsHandler)
	app.POST("/admin/items/create", items.AdminCreateItemHandler)
	app.GET("/admin/items/create", items.CreateItemController)
	app.GET("/admin/items/:id", items.AdminGetSingleItem)
	app.DELETE("/admin/items/:id", items.AdminDeleteSingleItem)
	app.GET("/htmx/admin/items/:id/edit", items.HTMXAdminItemsOneEdit)
	app.GET("/htmx/admin/items/:id/cancel", items.HTMXAdminItemsOneCancelEdit)

	app.GET("/admin/authors", authors.AdminAuthorsHandler)
	app.DELETE("/admin/authors/:id", authors.AdminSingleAuthorDelete)
	app.GET("/admin/authors/create", authors.CreateAuthorForm)
	app.POST("/admin/authors/create", authors.Create)

	app.GET("/admin/uploads", uploads.AdminGetUploads)
	app.GET("/admin/uploads/create", uploads.AdminGetUploadForm)
	app.POST("/admin/uploads/create", uploads.AdminCreateUpload)

	app.GET("/admin/login", controller.LoginHandler)

	app.Logger.Fatal(app.Start(":8080"))
}
