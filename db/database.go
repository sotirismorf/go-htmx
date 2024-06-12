package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sotirismorf/go-htmx/schema"
)

var Queries *schema.Queries

func ConnectDB() {
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, "postgresql://username:password@127.0.0.1:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close(ctx)

	queries := schema.New(conn)
	Queries = queries
}
