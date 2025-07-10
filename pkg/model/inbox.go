package model

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/htw-archive/pkg/s3store"
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
	Rows      []InboxItem
	PageCount int
	Count     int
}

func GetInboxCount(db *pgxpool.Pool) (PaginatedInboxItems, error) {
	c := PaginatedInboxItems{}
	err := c.getTotalCount(db)
	return c, err
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

	err = p.getTotalCount(db)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return p.fromRow(row)
}

func (p *PaginatedInboxItems) getTotalCount(db *pgxpool.Pool) error {
	sql := "select count(*)/@limit, count(*) from uploads where status='waiting'"

	row := db.QueryRow(context.Background(), sql, pgx.NamedArgs{
		"limit": 15,
	})

	return row.Scan(&p.PageCount, &p.Count)

}

func (p *PaginatedInboxItems) fromRow(row pgx.Rows) error {
	//defer row.Close()
	p.Rows = []InboxItem{}
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
	row, err := db.Query(context.Background(), "SELECT id, uri, trackname, artistid, createdby, createdat::text, status FROM uploads where id = @id;", pgx.NamedArgs{
		"id": id,
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return item.fromRow(row)
}
func (item *InboxItem) Delete(db *pgxpool.Pool) error {
	tx, err := db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	err = s3store.DeleteTrackFromInbox(item.UploadId + ".wav")
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), "delete from uploads where id = @id", pgx.NamedArgs{
		"id": item.UploadId,
	})
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
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

	o, err := s3store.UploadTrackToS3(objId, fileobj)
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
