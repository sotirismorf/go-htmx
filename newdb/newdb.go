package newdb

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func GetDB(uri string) error {
	DB, err := PgxCreateDB(uri)
	if err != nil {
		return err
	}
	DB.SetMaxIdleConns(2)
	DB.SetMaxOpenConns(4)
	DB.SetConnMaxLifetime(time.Duration(30) * time.Minute)

	db = DB

	return nil
}

func PgxCreateDB(uri string) (*sqlx.DB, error) {
	connConfig, _ := pgx.ParseConfig(uri)
	afterConnect := stdlib.OptionAfterConnect(func(ctx context.Context, conn *pgx.Conn) error {
		_, err := conn.Exec(ctx, `
			 SET SESSION "some.key" = 'somekey';
			 CREATE TEMP TABLE IF NOT EXISTS sometable AS SELECT 212 id;
		`)
		if err != nil {
			return err
		}
		return nil
	})

	pgxdb := stdlib.OpenDB(*connConfig, afterConnect)
	return sqlx.NewDb(pgxdb, "pgx"), nil
}
