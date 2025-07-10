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

	ag.GET("/:id", IdChecker, func(ctx *gin.Context) {
		id := ctx.Param("id")

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

	ag.PUT("/:id", AdminMiddleware, IdChecker, func(ctx *gin.Context) {
		id := ctx.Param("id")

		var r model.Release
		r.ReleaseId = id

		form, err := ctx.MultipartForm()
		if err != nil {
			fmt.Println(err)
			return
		}

		r.FromMultipartForm(db, form)
		err = r.Edit(db)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}

		fmt.Println(r)
	})

	ag.POST("/", AdminMiddleware, func(ctx *gin.Context) {
		//session := sessions.Default(ctx)
		var req model.Release

		if err := ctx.BindJSON(&req); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		//id := session.Get("userid").(string)

		err := req.Create(db)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		fmt.Println("new id: ", req.ReleaseId)

		ctx.Status(http.StatusOK)
	})
}
