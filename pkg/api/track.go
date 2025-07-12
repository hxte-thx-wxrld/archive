package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/htw-archive/pkg/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

const PAGINATED_COUNT = 15

func TrackApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	ag := rg.Group("/track")

	ag.GET("/", func(ctx *gin.Context) {
		offset := ctx.Query("offset")

		offset_int, err := strconv.Atoi(offset)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		lookup := model.PaginatedMusicLookup{}
		err = lookup.AllTracks(db, offset_int)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
		} else {
			ctx.JSON(http.StatusOK, lookup)
		}
	})

	ag.PUT("/:id", AuthenticatedMiddleware, IdChecker, func(ctx *gin.Context) {
		id := ctx.Param("id")

		var t model.Music
		if err := ctx.BindJSON(&t); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		t.TrackId = id

		t.EditTrack(db)

		/*if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}*/

		ctx.Status(http.StatusOK)
	})

	ag.GET("/:id", IdChecker, func(ctx *gin.Context) {
		id := ctx.Param("id")

		track := model.Music{}
		err := track.FromId(db, id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, track)
	})

	ag.DELETE("/:id", IdChecker, func(ctx *gin.Context) {
		id := ctx.Param("id")

		track := model.Music{
			TrackId: id,
		}
		err := track.FromId(db, id)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		err = track.Delete(db)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.Status(http.StatusOK)
	})

	ag.GET("/:id/analysis", IdChecker, func(ctx *gin.Context) {
		id := ctx.Param("id")

		t := ctx.Query("type")

		a := model.Analysis{
			TrackId: id,
		}
		err := a.ForTrackId(db)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		if t == "json" {

			fmt.Println("json export")
			ctx.Header("Content-Disposition", "attachment; filename="+id+"-analysis.json")
			ctx.Header("Content-Type", "application/json")
		}

		ctx.JSON(http.StatusOK, a)

	})
}
