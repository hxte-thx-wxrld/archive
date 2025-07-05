package api

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Release struct {
	Name         string
	ReleaseDate  string
	Isrc         *string
	ReleaseId    string
	CatalogId    string
	CoverUrl     string
	RelatedMusic []MusicInRelease
}

type MusicInRelease struct {
	TrackId    string
	Name       string
	Order      int
	ArtistName string
}

func GetMusicInReleases(db *pgxpool.Pool, releaseId string) ([]MusicInRelease, error) {
	rows, err := db.Query(context.Background(), "select m.id, m.title, mr.order, i.name from music_in_releases mr join music m on m.id = mr.music_id join interpret i on m.artist_id = i.id where release_id = @releaseId order by mr.order", pgx.NamedArgs{
		"releaseId": releaseId,
	})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var r []MusicInRelease
	for rows.Next() {
		var row MusicInRelease
		rows.Scan(&row.TrackId, &row.Name, &row.Order, &row.ArtistName)

		r = append(r, row)
	}

	return r, err
}

func GetAllReleases(db *pgxpool.Pool) ([]Release, error) {
	sql := "select name, release_date::text, isrc, public_cover_url, id, catalog_id from releases order by release_date desc"
	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var r []Release
	for rows.Next() {
		var release Release
		rows.Scan(&release.Name, &release.ReleaseDate, &release.Isrc, &release.CoverUrl, &release.ReleaseId, &release.CatalogId)

		r = append(r, release)
	}

	return r, nil
}

func GetSingleRelease(db *pgxpool.Pool, releaseId string) (*Release, error) {
	row := db.QueryRow(context.Background(), "select name, release_date::text, isrc, public_cover_url, id, catalog_id from releases where id = @releaseId", pgx.NamedArgs{
		"releaseId": releaseId,
	})

	var release Release

	err := row.Scan(&release.Name, &release.ReleaseDate, &release.Isrc, &release.CoverUrl, &release.ReleaseId, &release.CatalogId)
	if err != nil {
		return nil, err
	}

	music, err := GetMusicInReleases(db, releaseId)
	if err != nil {
		return nil, err
	}
	release.RelatedMusic = music

	return &release, nil
}

func ReleaseApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	ag := rg.Group("/release")
	ag.GET("/", func(ctx *gin.Context) {
		r, err := GetAllReleases(db)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, r)
	})

	ag.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid release id"))
			return
		}

		r, err := GetSingleRelease(db, id)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, r)
	})
}
