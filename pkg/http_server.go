package htwarchive

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/htw-archive/pkg/api"
	"github.com/jackc/pgx/v5/pgxpool"
)

var r = gin.Default()

func Serve(conn *pgxpool.Pool, index_html []byte, f embed.FS) {

	sendWebClient := func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html", index_html)
	}

	defer conn.Close()

	r.SetTrustedProxies([]string{"172.16.0.0/12"})
	r.MaxMultipartMemory = 157286400

	store := cookie.NewStore(([]byte("secret")))
	store.Options(sessions.Options{
		HttpOnly: true,
		Secure:   false,
	})
	//r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Cookies"},

		AllowCredentials: true,
	}))
	r.Use(sessions.Sessions("ui_session", store))

	dist, err := fs.Sub(f, "web/dist/assets")
	if err != err {
		log.Fatalln(err)
	}

	r.StaticFS("/assets", http.FS(dist))

	r.GET("/:page", sendWebClient)
	r.GET("/", sendWebClient)
	/*r.GET("/login", func(c *gin.Context) {

	})*/

	ag := r.Group("/api")

	api.InitApi(ag, conn)

	r.Run(":8080")
}
