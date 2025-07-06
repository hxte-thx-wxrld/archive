package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Artist struct {
	Name          string
	ArtistId      *string
	ArtistPicture string
}

func AssignArtistToUser(db *pgxpool.Pool, userid string, artistid string) error {
	_, err := db.Query(context.Background(), "insert into artists_of_user (user_id, artist_id) values (@userid, @artistid)", pgx.NamedArgs{
		"userid":   userid,
		"artistid": artistid,
	})

	if err != nil {
		return err
	}

	return nil
}

func GetAssignedArtists(db *pgxpool.Pool, userid string) ([]Artist, error) {
	rows, err := db.Query(context.Background(), "select i.id as ArtistId, i.name from artists_of_user aou join interpret i on i.id = aou.artist_id where aou.user_id::text = @userId", pgx.NamedArgs{
		"userId": userid,
	})

	if err != nil {
		return nil, err
	}

	var a []Artist
	for rows.Next() {
		var artist Artist

		err := rows.Scan(&artist.ArtistId, &artist.Name)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		fmt.Println(artist)

		a = append(a, artist)
	}
	return a, nil
}

func CreateArtist(db *pgxpool.Pool, name string, for_user string) (*string, error) {
	rows := db.QueryRow(context.Background(), "insert into interpret (name) values (@name) returning id::text", pgx.NamedArgs{
		"name": name,
	})

	var id string
	err := rows.Scan(&id)
	if err != nil {
		return nil, err
	}

	err = AssignArtistToUser(db, for_user, id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func GetAllArtists(db *pgxpool.Pool) ([]Artist, error) {
	rows, err := db.Query(context.Background(), "select id, name, artist_picture from interpret")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var a []Artist

	for rows.Next() {

		var artist Artist
		rows.Scan(&artist.ArtistId, &artist.Name, &artist.ArtistPicture)
		a = append(a, artist)
	}

	return a, nil
}

func GetSingleArtist(db *pgxpool.Pool, artistId string) (*Artist, error) {
	row := db.QueryRow(context.Background(), "select id, name, artist_picture from interpret where id = @artistId", pgx.NamedArgs{
		"artistId": artistId,
	})

	var artist Artist

	err := row.Scan(&artist.ArtistId, &artist.Name, &artist.ArtistPicture)
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func ArtistApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	ag := rg.Group("/artist")

	ag.GET("/", func(ctx *gin.Context) {
		a, err := GetAllArtists(db)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, a)
	})

	ag.POST("/", AuthenticatedMiddleware, func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		var req Artist

		if err := ctx.BindJSON(&req); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		id := session.Get("userid")

		id, err := CreateArtist(db, req.Name, id.(string))
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		fmt.Println("new id: ", id)

		ctx.Status(http.StatusOK)
	})

	ag.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid release id"))
			return
		}

		a, err := GetSingleArtist(db, id)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, a)
	})
}
