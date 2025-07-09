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

		fmt.Println(lookup)
		/*var music []MusicRow
		if offset != "" && conv_err == nil {

			music, err = GetTracksPaginated(db, offset_int)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusInternalServerError, err)
			}
		} else {
			music, err = GetTracks(db)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusInternalServerError, err)
			}
		}*/

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
		} else {
			ctx.JSON(http.StatusOK, lookup)
		}
	})

	// Upload track to database
	/*ag.POST("/", AuthenticatedMiddleware, func(ctx *gin.Context) {
		sess := sessions.Default(ctx)
		userid := sess.Get("userid").(string)
		admin := sess.Get("admin").(bool)

		fmt.Println("is admin", admin)

		form, err := ctx.MultipartForm()
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		var req UploadedTrack

		err = ctx.Bind(&req)

		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		fmt.Println(req)
		var tmpaudio multipart.File

		for _, file := range form.File["AudioFile"] {

			//if file.Header.Get("Content-Type") == "audio/x-wav" {

			tmpaudio, err = file.Open()
			if err != nil {
				fmt.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}

			break
			//}

		}

		res, err := registerUpload(db, req, tmpaudio, userid, admin)
		if err != nil {
			fmt.Println("db", err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		//err := row.Scan()

		//if err != nil {
		//	fmt.Println(err)
		//	ctx.JSON(http.StatusInternalServerError, err)
		//} else {

		ctx.JSON(http.StatusOK, res)
		//}

	})*/

	ag.PUT("/:id", AuthenticatedMiddleware, func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid track id"))
			return
		}

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

	ag.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid track id"))
			return
		}

		track := model.Music{}
		err := track.FromId(db, id)

		fmt.Println(track)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, track)
	})
}
