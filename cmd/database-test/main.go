package main

import (
	"context"
	"fmt"
	"log"

	htwarchive "github.com/htw-archive/pkg"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	// database connection
	conn := htwarchive.InitDb()
	for _ = range 100 {
		DoQuery(conn)
		DoQuery(conn)
		DoQuery(conn)
		DoQuery(conn)
		showStats(conn)
	}
}

func showStats(db *pgxpool.Pool) {
	stats := db.Stat()
	fmt.Println("--------")
	fmt.Println("Database Stats")
	fmt.Println("Acquired Count:", stats.AcquireCount())
	fmt.Println("Acquired Duration:", stats.AcquireDuration())
	fmt.Println("Acquired Conns:", stats.AcquiredConns())
	fmt.Println("Contructing Conns:", stats.ConstructingConns())
	fmt.Println("Empty Acquire Conns:", stats.EmptyAcquireCount())
	fmt.Println("Canceled Acquire Conns:", stats.CanceledAcquireCount())
	fmt.Println("Idle Conns:", stats.IdleConns())
	fmt.Println("Max Conns:", stats.MaxConns())
	fmt.Println("Total Conns:", stats.TotalConns())
	fmt.Println("--------")
}

func DoQuery(conn *pgxpool.Pool) {
	rows, err := conn.Query(context.Background(), "SELECT 1+1 as result")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {

		var result int
		err = rows.Scan(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Query Finished:", result)
	}
}
