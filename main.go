package main

import (
	"embed"
	"fmt"

	htwarchive "github.com/htw-archive/pkg"
)

//go:embed web/dist
var f embed.FS

//go:embed scripts/*
var scripts embed.FS

//go:embed web/dist/index.html
var index_html []byte

func main() {
	_ = htwarchive.DefaultDaemon(&scripts)

	//config := htwarchive.DefaultServerConfig()
	conn := htwarchive.InitDb()
	htwarchive.Serve(conn, index_html, f)
	fmt.Println("post serve")
}
