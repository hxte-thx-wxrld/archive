package htwarchive

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/htw-archive/api"
	"github.com/jackc/pgx/v5"
)

var r = gin.Default()

//go:embed web/dist
var f embed.FS

//go:embed web/dist/index.html
var index_html []byte

func Serve(conn *pgx.Conn) {

	defer conn.Close(context.Background())

	dist, err := fs.Sub(f, "web/dist/assets")
	if err != err {
		log.Fatalln(err)
	}

	r.Use(cors.Default())

	r.StaticFS("/assets", http.FS(dist))

	r.GET("/", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html", index_html)
	})

	ag := r.Group("/api")
	api.InitApi(ag, conn)

	r.Run(":8080")
}
