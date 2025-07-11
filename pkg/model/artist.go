package model

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Artist struct {
	Name          string
	ArtistId      *string
	ArtistPicture string
	Description   string
}

type ArtistCreateRequest struct {
	Name string
}

type PaginatedArtistLookup struct {
	Rows       []Artist
	FullLength int
}

func (a *PaginatedArtistLookup) AllArtists(db *pgxpool.Pool, offset int) error {
	rows, err := db.Query(context.Background(), "select id, name, artist_picture, description from interpret LIMIT @limit OFFSET @offset", pgx.NamedArgs{
		"limit":  PAGINATED_COUNT,
		"offset": offset * PAGINATED_COUNT,
	})
	if err != nil {
		return err
	}

	a.getTotalCount(db)

	return a.fromRow(rows)
}
func (a *PaginatedArtistLookup) fromRow(rows pgx.Rows) error {
	a.Rows = []Artist{}
	for rows.Next() {

		artist := Artist{}
		artist.fromRow(rows)
		rows.Scan(&artist.ArtistId, &artist.Name, &artist.ArtistPicture, &artist.Description)
		a.Rows = append(a.Rows, artist)
	}

	return nil
}
func (a *PaginatedArtistLookup) getTotalCount(db *pgxpool.Pool) error {
	sql := "select count(*)/@limit from interpret"

	row := db.QueryRow(context.Background(), sql, pgx.NamedArgs{
		"limit": PAGINATED_COUNT,
	})

	return row.Scan(&a.FullLength)
}

func (a *Artist) FromId(db *pgxpool.Pool) error {
	//"select id, name, artist_picture from interpret"
	row, err := db.Query(context.Background(), "select id, name, artist_picture, description from interpret where id = @id", pgx.NamedArgs{
		"id": a.ArtistId,
	})

	if err != nil {
		return err
	}

	row.Next()
	return a.fromRow(row)
}

func (a *Artist) fromRow(rows pgx.Rows) error {
	return rows.Scan(&a.ArtistId, &a.Name, &a.ArtistPicture, &a.Description)
}

func (a *Artist) CreateArtist(db *pgxpool.Pool, parentUserId string) error {
	/*rows := db.QueryRow(context.Background(), "insert into interpret (name) values (@name) returning id::text", pgx.NamedArgs{
	"name": name,
	})*/
	rows := db.QueryRow(context.Background(), "insert into interpret (name) values (@name) returning id::text", pgx.NamedArgs{
		"name": a.Name,
	})

	err := rows.Scan(&a.ArtistId)
	if err != nil {
		return err
	}

	return a.assignArtist(db, parentUserId)
}

func (a *Artist) assignArtist(db *pgxpool.Pool, parentUserId string) error {
	_, err := db.Query(context.Background(), "insert into artists_of_user (user_id, artist_id) values (@userid, @artistid)", pgx.NamedArgs{
		"userid":   parentUserId,
		"artistid": a.ArtistId,
	})
	return err
}
