package htwarchive

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDatabaseUri() string {

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s", os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_DATABASE"), os.Getenv("PG_SCHEMA"))
}

func InitDb() *pgxpool.Pool {

	conn, err := pgxpool.New(context.Background(), GetDatabaseUri())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// a
	return conn
	// defer conn.Close(context.Background())
}
