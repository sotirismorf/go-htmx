package crud

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sotirismorf/go-htmx/db"
	"github.com/sotirismorf/go-htmx/schema"
)

type itemFormData struct {
	Name        string  `form:"name"`
	Description string  `form:"description"`
	Year        int16   `form:"year"`
	AuthorID    []int64 `form:"author"`
	UploadID    []int64 `form:"upload"`
}

func CreateItem(c echo.Context) error {
	ctx := context.Background()

	var formData itemFormData

	err := c.Bind(&formData)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	itemParams := schema.CreateItemParams{
		Name: formData.Name,
		Year: formData.Year,
	}

	if formData.Description != "" {
		itemParams.Description = &formData.Description
	}

	createdItem, err := db.Queries.CreateItem(ctx, itemParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	for _, v := range formData.AuthorID {
		itemHasAuthorParams := schema.CreateItemHasAuthorRelationshipParams{
			ItemID:   createdItem.ID,
			AuthorID: int64(v),
		}

		_, err = db.Queries.CreateItemHasAuthorRelationship(ctx, itemHasAuthorParams)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
	}

	for _, v := range formData.UploadID {
		itemHasUploadParams := schema.CreateItemHasUploadRelationshipParams{
			ItemID:   createdItem.ID,
			UploadID: int64(v),
		}

		_, err = db.Queries.CreateItemHasUploadRelationship(ctx, itemHasUploadParams)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
	}

	return c.Redirect(http.StatusFound, "/admin/items")
}

type authorFormData struct {
	Name string `form:"name"`
	Bio  string `form:"bio"`
}

func CreateAuthor(c echo.Context) error {
	ctx := context.Background()

	var formData authorFormData

	err := c.Bind(&formData)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	createAuthorParams := schema.CreateAuthorParams{Name: formData.Name}

	if formData.Bio != "" {
		createAuthorParams.Bio = &formData.Bio
	}

	_, err = db.Queries.CreateAuthor(ctx, createAuthorParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.Redirect(http.StatusFound, "/admin/authors")
}

type groupFormData struct {
	Name string `form:"name"`
	City int16  `form:"location"`
}

func CreateGroup(c echo.Context) error {
	ctx := context.Background()

	var formData groupFormData

	err := c.Bind(&formData)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	dbParams := schema.CreateGroupParams{
		Name:     formData.Name,
		Location: formData.City,
	}

  fmt.Println(dbParams)

	err = db.Queries.CreateGroup(ctx, dbParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.Redirect(http.StatusFound, "/admin/groups")
}
