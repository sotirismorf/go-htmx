package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/invopop/ctxi18n"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sotirismorf/go-htmx/controller"
	"github.com/sotirismorf/go-htmx/controller/crud"
	"github.com/sotirismorf/go-htmx/controller/items"
	"github.com/sotirismorf/go-htmx/controller/ui"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/newdb"
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

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

    fmt.Println(sess.Values)

		auth, ok := sess.Values["authenticated"].(bool)
		if !auth || !ok {
	    return c.Redirect(http.StatusFound, "/login")
		}

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

	err := newdb.GetDB("postgresql://username:password@127.0.0.1:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri}\n",
	}))
	e.Use(i18nMiddleware)
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/", controller.HomeHandler)
	e.GET("/downloads/:id", crud.GetUpload)
	e.GET("/item/:id", controller.Item)
	e.GET("/login", controller.LoginHandler)
	e.GET("/search", ui.Search)
	e.POST("/htmx/multi-select-dropdown", items.HTMXMultiSelectDropdown)
	e.POST("/login", controller.CreateSession)
	e.POST("/logout", controller.Logout)
	e.Static("/assets", "static")
	e.Static("/static/thumbnails", "uploads/thumbnails")

	// Admin Group
	g := e.Group("/admin")
	g.Use(AuthMiddleware)

	g.DELETE("/authors/:id", crud.DeleteAuthor)
	g.DELETE("/items/:id", crud.DeleteItem)
	g.DELETE("/uploads/:id", crud.DeleteUpload)
	g.GET("", controller.AdminHandler)
	g.GET("/authors", ui.AdminAuthors)
	g.GET("/authors/create", ui.AdminCreateAuthorForm)
	g.GET("/groups", ui.AdminGroups)
	g.GET("/groups/create", ui.AdminCreateGroupForm)
	g.GET("/items/:id", ui.AdminSingleItem)
	g.GET("/items/create", ui.AdminCreateItemForm)
	g.GET("/uploads", ui.AdminUploads)
	g.GET("/uploads/create", ui.AdminCreateUploadForm)
	g.GET("/items", ui.AdminItems)
	g.POST("/authors/create", crud.CreateAuthor)
	g.POST("/groups/create", crud.CreateGroup)
	g.POST("/items/create", crud.CreateItem)
	g.POST("/uploads/create", crud.CreateUpload)
	// e.GET("/htmx/admin/items/:id/edit", ui.HTMXAdminItemsOneEdit)
	// e.GET("/htmx/admin/items/:id/cancel", items.HTMXAdminItemsOneCancelEdit)

	e.Logger.Fatal(e.Start(":8080"))
}
