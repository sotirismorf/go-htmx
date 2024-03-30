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
	app.GET("/admin/items", handlers.AdminHandler)
	app.POST("/admin/items", handlers.AdminCreateItemHandler)
  app.GET("/admin/items/:id", handlers.AdminSingleItemHandler)
	app.GET("/admin/login", handlers.LoginHandler)

	app.Logger.Fatal(app.Start(":8080"))
}
