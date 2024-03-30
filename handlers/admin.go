package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/components"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/schema"
)

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func LoginHandler(c echo.Context) error {
	return Render(c, http.StatusOK, components.Login())
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
