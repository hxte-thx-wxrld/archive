package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	htwarchive "github.com/htw-archive/pkg"
	"github.com/htw-archive/pkg/api"
	"github.com/htw-archive/pkg/s3store"
	"github.com/joho/godotenv"
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
	godotenv.Load(".env")

	// SIGINT handler
	sigs := make(chan os.Signal, 1)
	//done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancelDaemon := context.WithCancel(context.Background())

	htwarchive.DefaultDaemon(ctx, &scripts)

	//config := htwarchive.DefaultServerConfig()
	s3store.SetupS3Store(s3assets)
	conn := htwarchive.InitDb()
	err := api.UpdateAdminAccess(conn, os.Getenv("ROOT_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}

	go htwarchive.Serve(conn, index_html, f)

	<-sigs
	// ctrl-c was presser
	cancelDaemon()

	fmt.Println("waiting for daemon to end...")
	<-ctx.Done()
	time.Sleep(1 * time.Second)
}
