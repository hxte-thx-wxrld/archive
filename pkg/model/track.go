package model

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Music struct {
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

type PaginatedMusicLookup struct {
	Rows       []Music
	FullLength int
}

func (m *PaginatedMusicLookup) AllTracks(db *pgxpool.Pool, offset int) error {
	rows, err := db.Query(context.Background(),
		"select track_id, tracktitle, artist_id, artist, catalog_no, release_date::text, url, release_id, cover_url, length::text from all_tracks order by release_date desc LIMIT @limit OFFSET @offset", pgx.NamedArgs{
			"limit":  PAGINATED_COUNT,
			"offset": offset * PAGINATED_COUNT,
		})
	if err != nil {
		return err
	}

	err = m.getTotalCount(db)
	if err != nil {
		return err
	}

	return m.fromRow(rows)
}

func (m *PaginatedMusicLookup) getTotalCount(db *pgxpool.Pool) error {
	sql := "select count(*)/@limit from all_tracks"

	row := db.QueryRow(context.Background(), sql, pgx.NamedArgs{
		"limit": PAGINATED_COUNT,
	})

	return row.Scan(&m.FullLength)
}

func (m *PaginatedMusicLookup) fromRow(rows pgx.Rows) error {
	for rows.Next() {
		entry := Music{}
		entry.fromRow(rows)
		m.Rows = append(m.Rows, entry)
	}
	return nil
}

func (m *Music) FromId(db *pgxpool.Pool, id string) error {
	row, err := db.Query(context.Background(), "select track_id, tracktitle, artist_id, artist, catalog_no, release_date::text, url, release_id, cover_url, length::text from all_tracks where track_id = @id", pgx.NamedArgs{
		"id": id,
	})

	if err != nil {
		return err
	}

	row.Next()
	return m.fromRow(row)
}

func (m *Music) fromRow(row pgx.Rows) error {
	if err := row.Scan(&m.TrackId, &m.Tracktitle, &m.ArtistId, &m.Artist, &m.CatalogNo, &m.ReleaseDate, &m.PublicUrl, &m.ReleaseId, &m.CoverUrl, &m.Length); err != nil {
		fmt.Println("Music.fromRow:", err)
		return err
	}
	return nil
}

func (m *Music) EditTrack(db *pgxpool.Pool) error {
	row := db.QueryRow(context.Background(), "update music set title=@title, artist_id=@artistId, release_date=@releasedate WHERE id=@id", pgx.NamedArgs{
		"title":       m.Tracktitle,
		"artistId":    m.ArtistId,
		"releasedate": m.ReleaseDate,
		"id":          m.TrackId,
	})

	return row.Scan()
}
