package htwarchive

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitDb() *pgx.Conn {

	urlExample := os.Getenv("PG_URL")

	log.Println(urlExample)

	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
	// defer conn.Close(context.Background())
}
