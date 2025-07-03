package htwarchive

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/htw-archive/api"
	"github.com/jackc/pgx/v5/pgxpool"
)

var r = gin.Default()

//go:embed web/dist
var f embed.FS

//go:embed web/dist/index.html
var index_html []byte

func sendWebClient(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html", index_html)
}

func Serve(conn *pgxpool.Pool) {

	defer conn.Close()

	r.Use(cors.Default())
	dist, err := fs.Sub(f, "web/dist/assets")
	if err != err {
		log.Fatalln(err)
	}

	r.StaticFS("/assets", http.FS(dist))

	r.GET("/:page", sendWebClient)
	r.GET("/", sendWebClient)

	ag := r.Group("/api")

	api.InitApi(ag, conn)

	r.Run(":8080")
}
