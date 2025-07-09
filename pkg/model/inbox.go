package model

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InboxItem struct {
	UploadId    string
	Uri         string
	Trackname   string
	ArtistId    string
	CreatedBy   string
	CreatedAt   string
	Status      string
	ReleaseDate string
}

type PaginatedInboxItems struct {
	Rows       []InboxItem
	FullLength int
}

func (p *PaginatedInboxItems) AllWaitingInboxItems(db *pgxpool.Pool, page int) error {
	row, err := db.Query(context.Background(), "SELECT id, uri, trackname, artistid, createdby, createdat::text, status, release_date::text FROM uploads where status = 'waiting' limit @limit offset @offset;", pgx.NamedArgs{
		"limit":  15,
		"offset": 15 * page,
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	c, err := p.getTotalCount(db)
	if err != nil {
		fmt.Println(err)
		return err
	}
	p.FullLength = *c

	return p.fromRow(row)
}

func (p *PaginatedInboxItems) getTotalCount(db *pgxpool.Pool) (*int, error) {
	sql := "select count(*)/@limit from uploads"

	row := db.QueryRow(context.Background(), sql, pgx.NamedArgs{
		"limit": 15,
	})

	var count int
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil

}

func (p *PaginatedInboxItems) fromRow(row pgx.Rows) error {
	for row.Next() {
		item := InboxItem{}
		if err := item.fromRow(row); err != nil {
			fmt.Println(err)
			return err
		}
		p.Rows = append(p.Rows, item)
	}

	return nil
}

func (item *InboxItem) FromId(db *pgxpool.Pool, id string) error {
	row, err := db.Query(context.Background(), "SELECT id, uri, trackname, artistid, createdby, createdat::text, status FROM uploads where status = 'waiting' and id = @id;", pgx.NamedArgs{
		"id": id,
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return item.fromRow(row)
}
func (item *InboxItem) Accept(db *pgxpool.Pool) error {
	_, err := db.Query(context.Background(), "update uploads set status='accepted' where id = @id", pgx.NamedArgs{
		"id": item.UploadId,
	})

	return err
}

func (item *InboxItem) fromRow(row pgx.Rows) error {
	err := row.Scan(&item.UploadId, &item.Uri, &item.Trackname, &item.ArtistId, &item.CreatedBy, &item.CreatedAt, &item.Status, &item.ReleaseDate)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (item *InboxItem) RegisterUpload(db *pgxpool.Pool, fileobj multipart.File, admin bool) error {
	tx, err := db.BeginTx(context.Background(), pgx.TxOptions{})

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer tx.Rollback(context.Background())

	var status string
	if admin {
		status = "waiting"
	} else {
		status = "waiting"
	}

	rows := db.QueryRow(context.Background(), "INSERT INTO uploads (trackname, release_date, artistId, createdby, status) VALUES (@Title, @ReleaseDate::date, @ArtistId, @CreatedBy, @Status) returning id::text", pgx.NamedArgs{
		"Title":       item.Trackname,
		"ArtistId":    item.ArtistId,
		"ReleaseDate": item.ReleaseDate,
		"CreatedBy":   item.CreatedBy,
		"Status":      status,
	})

	var objId string
	err = rows.Scan(&objId)
	if err != nil {
		fmt.Println("sql", err)
		return err
	}

	o, err := UploadTrackToS3(objId, fileobj)
	if err != nil {
		fmt.Println(err)
		return err
	}

	item.UploadId = objId

	fmt.Println(o)

	_ = db.QueryRow(context.Background(), "update uploads set uri='/inbox/track/' || id where id = @id", pgx.NamedArgs{
		"id": objId,
	})

	tx.Commit(context.Background())
	fmt.Println("Created Id: ", objId)
	return nil
}
