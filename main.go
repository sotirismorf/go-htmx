package main

import (
	"context"
	"encoding/json"

	// "fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/handlers"
	"github.com/sotirismorf/go-htmx/schema"
)

func main() {
	db.ConnectDB()

	app := echo.New()

	// app.Use(middleware.Logger())
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri}\n",
	}))

	app.GET("/", HomeHandler)
	app.GET("/admin", AdminHandler)
	app.POST("/admin/items", AdminCreateItemHandler)
	app.GET("/admin/login", handlers.LoginHandler)

	app.Logger.Fatal(app.Start(":8080"))
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func HomeHandler(c echo.Context) error {

	ctx := context.Background()

	data, err := db.Queries.ListAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	items, err := db.Queries.ListItemsWithAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	type author struct {
		Id   int    `json:"f1"`
		Name string `json:"f2"`
	}

	authors := []author{}
	json.Unmarshal([]byte(items[0].AuthorIds), &authors)

	return Render(c, http.StatusOK, components.Index(data))
}

func AdminHandler(c echo.Context) error {
	ctx := context.Background()

	items, err := db.Queries.ListItems(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	authors, err := db.Queries.ListAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return Render(c, http.StatusOK, components.Admin(items, authors))
}

func AdminCreateItemHandler(c echo.Context) error {
	ctx := context.Background()

	name := c.FormValue("name")
	description := c.FormValue("description")
	author := c.FormValue("author")

	log.Println("-" + name + "- and -" + description + "-" + author + "-")

	itemParams := schema.CreateItemParams{Name: name}

	if description != "" {
		itemParams.Description = &description
	}

	createdItem, err := db.Queries.CreateItem(ctx, itemParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if author != "" {
		println("Author is ", author)

		itemHasAuthorParams := schema.CreateItemHasAuthorRelationshipParams{
			ItemID:   createdItem.ID,
			AuthorID: 1,
		}
		_, err := db.Queries.CreateItemHasAuthorRelationship(ctx, itemHasAuthorParams)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
	}

	return c.Redirect(http.StatusFound, "/admin")
}
