// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package schema

import ()

type Author struct {
	ID   int64
	Name string
	Bio  *string
}

type Group struct {
	ID   int64
	Name string
}

type Item struct {
	ID          int64
	Name        string
	Description *string
}

type ItemHasAuthor struct {
	ItemID   int64
	AuthorID int64
}

type Publisher struct {
	ID          int64
	Name        string
	Description *string
}
