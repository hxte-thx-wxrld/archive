package main

import (
	"embed"
	"fmt"

	htwarchive "github.com/htw-archive/pkg"
	"github.com/htw-archive/pkg/s3store"
)

//go:embed web/dist
var f embed.FS

//go:embed assets/*
var s3assets embed.FS

//go:embed scripts/*
var scripts embed.FS

//go:embed web/dist/index.html
var index_html []byte

func main() {
	_ = htwarchive.DefaultDaemon(&scripts)

	//config := htwarchive.DefaultServerConfig()
	s3store.SetupS3Store(s3assets)
	conn := htwarchive.InitDb()
	htwarchive.Serve(conn, index_html, f)
	fmt.Println("post serve")
}
