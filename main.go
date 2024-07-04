package main

import (
	"embed"
	"log"

	"github.com/invopop/ctxi18n"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/controller/items"
	"github.com/sotirismorf/go-htmx/controller/ui"
	"github.com/sotirismorf/go-htmx/controller/crud"
	"github.com/sotirismorf/go-htmx/db"
)

func i18nMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Initialize the i18n context
		// You can configure the i18n options as needed
		i18nCtx, err := ctxi18n.WithLocale(c.Request().Context(), "el")
		if err != nil {
      log.Println("hey")
		}

		// Attach the i18n context to the echo context
		c.SetRequest(c.Request().WithContext(i18nCtx))

		return next(c)
	}
}

//go:embed translations
var locales embed.FS

func main() {
	if err := ctxi18n.Load(locales); err != nil {
		log.Fatalf("error loading locales: %v", err)
	}

	db.ConnectDB()

	e := echo.New()

	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri}\n",
	}))
	e.Use(i18nMiddleware)

	e.GET("/", controller.HomeHandler)
	e.GET("/search", ui.Search)
	e.GET("/item/:id", controller.Item)

	e.GET("/downloads/:id", crud.GetUpload)
	e.Static("/static/thumbnails", "uploads/thumbnails")
	e.Static("/assets", "assets")

	e.GET("/admin", controller.AdminHandler)
	e.GET("/admin/items", ui.AdminItems)
	e.POST("/admin/items/create", crud.CreateItem)
	e.GET("/admin/items/create", ui.AdminCreateItemForm)
	e.GET("/admin/items/:id", ui.AdminSingleItem)
	e.DELETE("/admin/items/:id", crud.DeleteItem)
	e.GET("/htmx/admin/items/:id/edit", ui.HTMXAdminItemsOneEdit)
	// e.GET("/htmx/admin/items/:id/cancel", items.HTMXAdminItemsOneCancelEdit)

	e.GET("/admin/groups", ui.AdminGroups)
	e.POST("/admin/groups/create", crud.CreateGroup)
	e.GET("/admin/groups/create", ui.AdminCreateGroupForm)

	e.GET("/admin/authors", ui.AdminAuthors)
	e.DELETE("/admin/authors/:id", crud.DeleteAuthor)
	e.GET("/admin/authors/create", ui.AdminCreateAuthorForm)
	e.POST("/admin/authors/create", crud.CreateAuthor)

	e.GET("/admin/uploads", ui.AdminUploads)
	e.GET("/admin/uploads/create", ui.AdminCreateUploadForm)
	e.POST("/admin/uploads/create", crud.CreateUpload)
	e.DELETE("/admin/uploads/:id", crud.DeleteUpload)

	e.GET("/admin/login", controller.LoginHandler)

	e.POST("/htmx/multi-select-dropdown", items.HTMXMultiSelectDropdown)

	e.Logger.Fatal(e.Start(":8080"))
}
