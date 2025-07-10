package htwarchive

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDb() *pgxpool.Pool {

	urlExample := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s", os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_DATABASE"), os.Getenv("PG_SCHEMA"))

	log.Println(urlExample)

	conn, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// a
	return conn
	// defer conn.Close(context.Background())
}
