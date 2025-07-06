package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const PAGINATED_COUNT = 15

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

type TrackEditRequest struct {
	Tracktitle  string
	ReleaseDate string
	ArtistId    string
}

func GetTracksTotalCount(db *pgxpool.Pool) (*int, error) {
	sql := "select count(*)/@limit from all_tracks"

	row := db.QueryRow(context.Background(), sql, pgx.NamedArgs{
		"limit": PAGINATED_COUNT,
	})

	var count int
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetTracksPaginated(db *pgxpool.Pool, offset int) ([]MusicRow, error) {
	rows, err := db.Query(context.Background(),
		"select track_id, tracktitle, artist, artist_id, catalog_no, release_date::text, url, cover_url, length::text from all_tracks LIMIT @limit OFFSET @offset", pgx.NamedArgs{
			"limit":  PAGINATED_COUNT,
			"offset": offset * PAGINATED_COUNT,
		})
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

func GetTracks(db *pgxpool.Pool) ([]MusicRow, error) {

	sql := "select track_id, tracktitle, artist, artist_id, catalog_no, release_date::text, url, cover_url, length::text from all_tracks order by release_date is null, release_date DESC "

	rows, err := db.Query(context.Background(), sql)

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

func UpdateTrack(db *pgxpool.Pool, id string, data TrackEditRequest) error {
	row := db.QueryRow(context.Background(), "update music set title=@title, artist_id=@artistId, release_date=@releasedate WHERE id=@id", pgx.NamedArgs{
		"title":       data.Tracktitle,
		"artistId":    data.ArtistId,
		"releasedate": data.ReleaseDate,
		"id":          id,
	})

	return row.Scan()
}

func TrackApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	ag := rg.Group("/track")

	ag.GET("/", func(ctx *gin.Context) {
		offset := ctx.Query("offset")

		fullLength, err := GetTracksTotalCount(db)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		offset_int, conv_err := strconv.Atoi(offset)

		var music []MusicRow
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
		}

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
		url, err := NewPresignUrl("inbox", new_id.String())

		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		fmt.Println(url)
		ctx.JSON(http.StatusOK, TrackPresignedUploadResponse{
			Url:     *url,
			TrackId: new_id.String(),
		})
	})

	// Upload track to database
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

		rows, err := db.Query(context.Background(), "INSERT INTO public.music (id, title, artist_id, public_url) VALUES(@TrackId, @Title, @ArtistId, @PublicUrl)", pgx.NamedArgs{
			"TrackId":   id,
			"Title":     req.TrackTitle,
			"ArtistId":  req.ArtistId,
			"PublicUrl": req.PublicUrl,
		})

		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		defer rows.Close()

		//err := row.Scan()

		//if err != nil {
		//	fmt.Println(err)
		//	ctx.JSON(http.StatusInternalServerError, err)
		//} else {
		ctx.Status(http.StatusOK)
		//}

	})

	ag.PUT("/:id", AuthenticatedMiddleware, func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid track id"))
			return
		}

		var req TrackEditRequest

		if err := ctx.BindJSON(&req); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		_ = UpdateTrack(db, id, req)
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

		track, err := GetSingleTrack(db, id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, track)
	})
}
