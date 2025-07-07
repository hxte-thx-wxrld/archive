package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewPresignUrl(bucket string, key string) (*string, error) {
	/*sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("US"),
		Endpoint: aws.String("http://weed:8333"),

		S3ForcePathStyle:              aws.Bool(true),
		Credentials:                   credentials.NewStaticCredentials("admin", "admin", "changeme"),
		CredentialsChainVerboseErrors: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		ResponseContentType: aws.String("audio/wav"),
		Bucket:              aws.String(bucket),
		Key:                 aws.String(key),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return nil, err
	}

	//return Presigner{PresignClient: presignClient}

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u.Host = "127.0.0.1:8333"
	uu := u.String()

	return &uu, err*/
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("US"))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	svc := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	presignClient := s3.NewPresignClient(svc)
	req, err := presignClient.PresignPutObject(context.Background(), &s3.PutObjectInput{
		ContentType: aws.String("audio/x-wav"),
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Expires:     aws.Time(time.Now().Add(15 * time.Minute)),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Println(req)

	u, err := url.Parse(req.URL)
	if err != nil {
		return nil, err
	}

	u.Host = "127.0.0.1:8333"
	uu := u.String()

	return &uu, nil
}

func InitApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	TrackApi(rg, db)
	ArtistApi(rg, db)
	ReleaseApi(rg, db)

	rg.GET("/me", AuthenticatedMiddleware, func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		id := session.Get("userid")
		if id == nil {
			ctx.Status(http.StatusInternalServerError)
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
