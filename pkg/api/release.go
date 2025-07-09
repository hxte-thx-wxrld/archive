package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/htw-archive/pkg/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ReleaseApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	ag := rg.Group("/release")
	ag.GET("/", func(ctx *gin.Context) {
		offset := ctx.Query("offset")

		offset_int, err := strconv.Atoi(offset)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		lookup := model.ReleasesLookup{}
		err = lookup.AllReleases(db, offset_int)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			fmt.Println(err)
			return
		}

		ctx.JSON(http.StatusOK, lookup)
	})

	ag.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid release id"))
			return
		}

		//r, err := GetSingleRelease(db, id)
		r := model.Release{}
		err := r.FromId(db, id)

		//fmt.Println(r)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, r)
	})
}
