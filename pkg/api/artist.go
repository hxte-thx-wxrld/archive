package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/htw-archive/pkg/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ArtistApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	ag := rg.Group("/artist")

	ag.GET("/", func(ctx *gin.Context) {
		//a, err := GetAllArtists(db)
		offset := ctx.Query("offset")

		offset_int, err := strconv.Atoi(offset)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		a := model.PaginatedArtistLookup{}
		err = a.AllArtists(db, offset_int)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, a)
	})

	ag.POST("/", AuthenticatedMiddleware, func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		var req model.Artist

		if err := ctx.BindJSON(&req); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		id := session.Get("userid").(string)

		err := req.CreateArtist(db, id)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		fmt.Println("new id: ", id)

		ctx.Status(http.StatusOK)
	})

	ag.GET("/:id", IdChecker, func(ctx *gin.Context) {
		id := ctx.Param("id")

		a := model.Artist{
			ArtistId: &id,
		}

		err := a.FromId(db)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, a)
	})

	ag.GET("/:id/tracks", IdChecker, func(ctx *gin.Context) {
		id := ctx.Param("id")
		offset := ctx.Query("offset")

		offset_int, err := strconv.Atoi(offset)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		lookup := model.PaginatedMusicLookup{}
		err = lookup.GetTracksByArtist(db, offset_int, id)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, lookup)
	})
}
