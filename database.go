package htwarchive

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDb() *pgxpool.Pool {

	urlExample := os.Getenv("PG_URL")

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
