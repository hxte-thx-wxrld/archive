package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const PAGINATED_COUNT = 15

type UploadedTrack struct {
	TrackTitle  string
	ArtistId    string
	ReleaseDate string
}

type UploadedTrackResponse struct {
	UploadId string
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

func registerUpload(db *pgxpool.Pool, req UploadedTrack, fileobj multipart.File, userid string) (*UploadedTrackResponse, error) {
	tx, err := db.BeginTx(context.Background(), pgx.TxOptions{})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer tx.Rollback(context.Background())
	rows := db.QueryRow(context.Background(), "INSERT INTO uploads (trackname, release_date, artistId, createdby) VALUES (@Title, @ReleaseDate::date, @ArtistId, @CreatedBy) returning id::text", pgx.NamedArgs{
		"Title":       req.TrackTitle,
		"ArtistId":    req.ArtistId,
		"ReleaseDate": req.ReleaseDate,
		"CreatedBy":   userid,
	})

	var objId string
	err = rows.Scan(&objId)
	if err != nil {
		fmt.Println("sql", err)
		return nil, err
	}

	o, err := uploadTrackToS3(objId, fileobj)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(o)

	_ = db.QueryRow(context.Background(), "update uploads set uri='/inbox/track/' || id where id = @id", pgx.NamedArgs{
		"id": objId,
	})

	tx.Commit(context.Background())
	fmt.Println("Created Id: ", objId)
	return &UploadedTrackResponse{
		UploadId: objId,
	}, nil
}

func uploadTrackToS3(trackName string, data multipart.File) (*s3.PutObjectOutput, error) {
	svc, err := NewS3Config()
	if err != nil {
		return nil, err
	}

	o, err := svc.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String("inbox"),
		Key:    aws.String("track/" + trackName),
		Body:   data,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return o, nil
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
		"select track_id, tracktitle, artist, artist_id, catalog_no, release_date::text, url, cover_url, length::text from all_tracks order by release_date desc LIMIT @limit OFFSET @offset", pgx.NamedArgs{
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

	/*ag.POST("/presign", AuthenticatedMiddleware, func(ctx *gin.Context) {
		hash := ctx.Query("hash")
		new_id := uuid.New()
		url, err := NewPresignUrl(hash, "inbox", new_id.String())

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
	})*/

	// Upload track to database
	ag.POST("/", AuthenticatedMiddleware, func(ctx *gin.Context) {
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

		var req UploadedTrack

		err = ctx.Bind(&req)

		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		fmt.Println(req)
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

		res, err := registerUpload(db, req, tmpaudio, userid)
		if err != nil {
			fmt.Println("db", err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		//err := row.Scan()

		//if err != nil {
		//	fmt.Println(err)
		//	ctx.JSON(http.StatusInternalServerError, err)
		//} else {

		ctx.JSON(http.StatusOK, res)
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
