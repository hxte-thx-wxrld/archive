package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewPresignUrl(bucket string, key string) (*string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("US"),
		Endpoint: aws.String("http://127.0.0.1:8333"),

		S3ForcePathStyle:              aws.Bool(true),
		Credentials:                   credentials.NewSharedCredentials("", "default"),
		CredentialsChainVerboseErrors: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}
	aws.NewConfig()
	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(15 * time.Minute)

	//return Presigner{PresignClient: presignClient}
	return &urlStr, err
}

func InitApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	TrackApi(rg, db)
	ArtistApi(rg, db)
	ReleaseApi(rg, db)

	rg.GET("/me", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		id := session.Get("userid")
		if id == nil {
			ctx.Status(http.StatusUnauthorized)
		} else {
			fmt.Println(id)

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
		session.Set("userid", string(*id))
		err = session.Save()
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.Status(http.StatusOK)
	})
}
