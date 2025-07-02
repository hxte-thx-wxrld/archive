package api

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type MusicRow struct {
	TrackId     string
	Tracktitle  string
	Artist      string
	ArtistId    string
	CatalogNo   *string
	ReleaseDate *string
	PublicUrl   *string
	ReleaseId   *string
	CoverUrl    *string
}

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

type Artist struct {
	Name          string
	ArtistId      string
	ArtistPicture string
}

func GetAllReleases(db *pgx.Conn) ([]Release, error) {
	sql := "select name, release_date::text, isrc, public_cover_url, id, catalog_id from releases"
	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	var r []Release
	for rows.Next() {
		var release Release
		rows.Scan(&release.Name, &release.ReleaseDate, &release.Isrc, &release.CoverUrl, &release.ReleaseId, &release.CatalogId)

		r = append(r, release)
	}

	return r, nil
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

func GetSingleRelease(db *pgx.Conn, releaseId string) (*Release, error) {
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

func GetMusicInReleases(db *pgx.Conn, releaseId string) ([]MusicInRelease, error) {
	rows, err := db.Query(context.Background(), "select m.id, m.title, mr.order, i.name from music_in_releases mr join music m on m.id = mr.music_id join interpret i on m.artist_id = i.id where release_id = @releaseId order by mr.order", pgx.NamedArgs{
		"releaseId": releaseId,
	})

	var r []MusicInRelease
	for rows.Next() {
		var row MusicInRelease
		rows.Scan(&row.TrackId, &row.Name, &row.Order, &row.ArtistName)

		r = append(r, row)
	}

	return r, err
}

func GetMusic(db *pgx.Conn) ([]MusicRow, error) {
	args := pgx.NamedArgs{}

	sql := "select track_id, tracktitle, artist, catalog_no, release_date::text, url, cover_url from all_tracks ORDER by release_date ASC"

	rows, err := db.Query(context.Background(), sql, args)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var music []MusicRow

	for rows.Next() {
		entry := MusicRow{}
		if err := rows.Scan(&entry.TrackId, &entry.Tracktitle, &entry.Artist, &entry.CatalogNo, &entry.ReleaseDate, &entry.PublicUrl, &entry.CoverUrl); err != nil {
			return nil, err
		}

		music = append(music, entry)
	}

	return music, nil
}

func GetSingleTrack(db *pgx.Conn, id string) (*MusicRow, error) {
	row := db.QueryRow(context.Background(), "select tracktitle, artist_id, artist, catalog_no, release_date::text, url, release_id, cover_url from all_tracks where track_id = @id", pgx.NamedArgs{
		"id": id,
	})

	var track MusicRow

	err := row.Scan(&track.Tracktitle, &track.ArtistId, &track.Artist, &track.CatalogNo, &track.ReleaseDate, &track.PublicUrl, &track.ReleaseId, &track.CoverUrl)
	track.TrackId = id

	if err != nil {
		return nil, err
	}

	return &track, nil
}

func InitApi(rg *gin.RouterGroup, db *pgx.Conn) {
	rg.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Ok")
	})

	rg.GET("/music/", func(ctx *gin.Context) {
		music, err := GetMusic(db)
		log.Println(music)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
		} else {
			ctx.JSON(http.StatusOK, music)
		}
	})

	rg.GET("/track/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid track id"))
			return
		}
		track, err := GetSingleTrack(db, id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, track)
	})

	rg.GET("/release", func(ctx *gin.Context) {
		r, err := GetAllReleases(db)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, r)
	})

	rg.GET("/release/:id", func(ctx *gin.Context) {
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

	rg.GET("/artist", func(ctx *gin.Context) {
		a, err := GetAllArtists(db)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, a)
	})

	rg.GET("/artist/:id", func(ctx *gin.Context) {
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
