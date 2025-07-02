package main

import (
	"log"

	htwarchive "github.com/htw-archive"
	"github.com/joho/godotenv"
)

func main() {
	// close database on return
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	conn := htwarchive.InitDb()
	htwarchive.Serve(conn)
}
