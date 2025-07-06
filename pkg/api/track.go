package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TrackPresignedUploadResponse struct {
	Url     string
	TrackId string
}

type UploadedTrack struct {
	TrackTitle string
	ArtistId   string
	PublicUrl  string
}

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
	Length      string
}

type MusicLookup struct {
	Rows       []MusicRow
	FullLength int
}

func GetTracksTotalCount(db *pgxpool.Pool) (*int, error) {
	sql := "select count(*)/15 from all_tracks"

	row := db.QueryRow(context.Background(), sql)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetTracks(db *pgxpool.Pool, limit int, offset int) ([]MusicRow, error) {
	args := pgx.NamedArgs{
		"limit":  "ALL",
		"offset": 0,
	}

	if limit != 0 {
		args["limit"] = limit
	}

	if offset != 0 {
		args["offset"] = offset
	}

	sql := "select track_id, tracktitle, artist, artist_id, catalog_no, release_date::text, url, cover_url, length::text from all_tracks order by release_date is null, release_date DESC "

	rows, err := db.Query(context.Background(), sql, args)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var music []MusicRow

	for rows.Next() {
		entry := MusicRow{}
		if err := rows.Scan(&entry.TrackId, &entry.Tracktitle, &entry.Artist, &entry.ArtistId, &entry.CatalogNo, &entry.ReleaseDate, &entry.PublicUrl, &entry.CoverUrl, &entry.Length); err != nil {
			return nil, err
		}

		music = append(music, entry)
	}

	return music, nil
}

func GetSingleTrack(db *pgxpool.Pool, id string) (*MusicRow, error) {
	row := db.QueryRow(context.Background(), "select tracktitle, artist_id, artist, catalog_no, release_date::text, url, release_id, cover_url, length::text from all_tracks where track_id = @id", pgx.NamedArgs{
		"id": id,
	})

	var track MusicRow

	err := row.Scan(&track.Tracktitle, &track.ArtistId, &track.Artist, &track.CatalogNo, &track.ReleaseDate, &track.PublicUrl, &track.ReleaseId, &track.CoverUrl, &track.Length)
	track.TrackId = id

	if err != nil {
		return nil, err
	}

	return &track, nil
}

func TrackApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	ag := rg.Group("/track")

	ag.GET("/", func(ctx *gin.Context) {
		fullLength, err := GetTracksTotalCount(db)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		music, err := GetTracks(db, 0, 0)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
		} else {
			ctx.JSON(http.StatusOK, MusicLookup{
				Rows:       music,
				FullLength: *fullLength,
			})
		}
	})

	ag.POST("/presign", AuthenticatedMiddleware, func(ctx *gin.Context) {
		new_id := uuid.New()
		url, err := NewPresignUrl("tracks", new_id.String()+".wav")

		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		fmt.Println(url)
		ctx.JSON(http.StatusOK, TrackPresignedUploadResponse{
			Url:     url,
			TrackId: new_id.String(),
		})
	})

	ag.POST("/:id", AuthenticatedMiddleware, func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid track id"))
			return
		}

		var req UploadedTrack

		if err := ctx.BindJSON(&req); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		db.QueryRow(context.Background(), "INSERT INTO public.music (id, title, artist_id, public_url) VALUES(@TrackId, @Title, @ArtistId, @PublicUrl)", pgx.NamedArgs{
			"TrackId":   id,
			"Title":     req.TrackTitle,
			"ArtistId":  req.ArtistId,
			"PublicUrl": req.PublicUrl,
		})

		//err := row.Scan()

		//if err != nil {
		//	fmt.Println(err)
		//	ctx.JSON(http.StatusInternalServerError, err)
		//} else {
		ctx.Status(http.StatusOK)
		//}

	})
	ag.GET("/:id", func(ctx *gin.Context) {
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
}
