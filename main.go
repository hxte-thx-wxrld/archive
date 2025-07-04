package main

import (
	"embed"
	"log"

	htwarchive "github.com/htw-archive/pkg"
	"github.com/joho/godotenv"
)

//go:embed web/dist
var f embed.FS

//go:embed web/dist/index.html
var index_html []byte

func main() {
	// close database on return
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}
	conn := htwarchive.InitDb()
	htwarchive.Serve(conn, index_html, f)
}
