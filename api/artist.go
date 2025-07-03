package api

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Artist struct {
	Name          string
	ArtistId      string
	ArtistPicture string
}

func GetAllArtists(db *pgx.Conn) ([]Artist, error) {
	rows, err := db.Query(context.Background(), "select id, name, artist_picture from interpret")
	if err != nil {
		return nil, err
	}

	var a []Artist

	for rows.Next() {

		var artist Artist
		rows.Scan(&artist.ArtistId, &artist.Name, &artist.ArtistPicture)
		a = append(a, artist)
	}

	return a, nil
}

func GetSingleArtist(db *pgx.Conn, artistId string) (*Artist, error) {
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

func ArtistApi(rg *gin.RouterGroup, db *pgx.Conn) {
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
