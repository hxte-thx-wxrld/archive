package api

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/htw-archive/pkg/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func FileInboxApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	ag := rg.Group("/inbox", AuthenticatedMiddleware)

	ag.POST("/", func(ctx *gin.Context) {
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

		var t model.InboxItem
		err = ctx.Bind(&t)

		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		t.CreatedBy = userid

		fmt.Println(t)

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

		err = t.RegisterUpload(db, tmpaudio, admin)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}

		ctx.JSON(http.StatusOK, t)
	})

	ag.GET("/", AdminMiddleware, func(ctx *gin.Context) {
		i := model.PaginatedInboxItems{}
		err := i.AllWaitingInboxItems(db, 0)
		fmt.Println(i)

		//inbox, err := GetInbox(db)

		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, i)
	})

	ag.POST("/:id/accept", AdminMiddleware, IdChecker, func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		item := model.InboxItem{
			UploadId: id,
		}
		item.Accept(db)

		ctx.Status(http.StatusOK)
	})

	ag.DELETE("/:id/delete", AdminMiddleware, IdChecker, func(ctx *gin.Context) {
		id := ctx.Param("id")

		item := model.InboxItem{
			UploadId: id,
		}

		if err := item.Delete(db); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		} else {
			ctx.Status(http.StatusOK)
		}
	})

	ag.GET("/count", AdminMiddleware, func(ctx *gin.Context) {
		c, err := model.GetInboxCount(db)
		if err != nil {
			fmt.Println(err)
		}

		ctx.JSON(http.StatusOK, c)
	})
}
