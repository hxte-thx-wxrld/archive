package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SELECT uri, trackname, artistid, createdby, createdat, status FROM uploads where id = @UploadId and status = 'waiting';

type InboxItem struct {
	UploadId  string
	Uri       string
	Trackname string
	ArtistId  string
	CreatedBy string
	CreatedAt string
	Status    string
}

type PaginatedInboxItems struct {
	Rows       []InboxItem
	FullLength int
}

func acceptUpload(db *pgxpool.Pool, UploadId string) {
	_ = db.QueryRow(context.Background(), "update uploads set status='accepted' where id = @id", pgx.NamedArgs{
		"id": UploadId,
	})
}

func GetInbox(db *pgxpool.Pool) ([]InboxItem, error) {
	row, err := db.Query(context.Background(), "SELECT id, uri, trackname, artistid, createdby, createdat::text, status FROM uploads where status = 'waiting';", pgx.NamedArgs{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer row.Close()

	var items []InboxItem

	for row.Next() {
		var item InboxItem
		err = row.Scan(&item.UploadId, &item.Uri, &item.Trackname, &item.ArtistId, &item.CreatedBy, &item.CreatedAt, &item.Status)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func FileInboxApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	ag := rg.Group("/inbox", AuthenticatedMiddleware, AdminMiddleware)

	ag.GET("/", func(ctx *gin.Context) {
		inbox, err := GetInbox(db)

		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, PaginatedInboxItems{
			Rows:       inbox,
			FullLength: 0,
		})
	})

	ag.GET("/:id/accept", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		acceptUpload(db, id)
		ctx.Status(http.StatusOK)
	})
}
