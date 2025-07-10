package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Checks, if the request has given an id
func IdChecker(ctx *gin.Context) {
	objId := ctx.Param("id")
	if objId == "" {
		ctx.JSON(http.StatusBadRequest, errors.New("no id given"))
	} else {
		ctx.Next()
	}
}

func InitApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	TrackApi(rg, db)
	ArtistApi(rg, db)
	ReleaseApi(rg, db)
	FileInboxApi(rg, db)

	rg.GET("/me", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		id := session.Get("userid")
		if id == nil {
			ctx.JSON(http.StatusOK, nil)
		} else {
			id = id.(string)

			user, err := LookupUser(db, id.(string))
			if err != nil {
				fmt.Println(err)
				ctx.JSON(http.StatusInternalServerError, err)
			}
			ctx.JSON(http.StatusOK, user)
		}

	})

	rg.POST("/signup", func(ctx *gin.Context) {
		var req LoginRequest

		if err := ctx.BindJSON(&req); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		_, err := CreateUser(db, req.Username, req.Password)

		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, errors.New("username already exists"))
			return
		}

		ctx.Status(http.StatusOK)
	})

	rg.POST("/logout", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Delete("userid")
		session.Save()
		ctx.Status(http.StatusOK)
	})

	rg.POST("/login", func(ctx *gin.Context) {
		var req LoginRequest

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := ctx.BindJSON(&req); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		session := sessions.Default(ctx)
		id, err := CheckPassword(db, req.Username, req.Password)

		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusUnauthorized, errors.New("wrong username or password"))
			return
		}

		fmt.Println("Userid:", *id)

		session.Set("userid", string(id.UserId))
		session.Set("username", id.Username)
		session.Set("admin", id.Admin)

		err = session.Save()
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.Status(http.StatusOK)
	})
}
