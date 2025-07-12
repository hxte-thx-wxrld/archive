package model

import (
	"context"
	"fmt"

	"github.com/htw-archive/pkg/s3store"
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
	Length      *string
}

type PaginatedMusicLookup struct {
	Rows       []Music
	FullLength int
	Count      int
}

func (m *PaginatedMusicLookup) GetTracksByArtist(db *pgxpool.Pool, offset int, id string) error {
	rows, err := db.Query(context.Background(),
		"select track_id, tracktitle, artist_id, artist, catalog_no, release_date::text, url, release_id, cover_url, length::text from all_tracks where artist_id = @id order by release_date desc LIMIT @limit OFFSET @offset", pgx.NamedArgs{
			"limit":  PAGINATED_COUNT,
			"offset": offset * PAGINATED_COUNT,
			"id":     id,
		})
	if err != nil {
		return err
	}
	defer rows.Close()

	err = m.getTotalCountForArtist(db, id)
	if err != nil {
		return err
	}

	return m.fromRow(rows)
}

func (m *PaginatedMusicLookup) getTotalCountForArtist(db *pgxpool.Pool, id string) error {
	sql := "select count(*)/@limit, count(*) from all_tracks where artist_id = @id"

	row := db.QueryRow(context.Background(), sql, pgx.NamedArgs{
		"limit": PAGINATED_COUNT,
		"id":    id,
	})

	return row.Scan(&m.FullLength, &m.Count)
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
	defer rows.Close()

	err = m.getTotalCount(db)
	if err != nil {
		return err
	}

	return m.fromRow(rows)
}

func (m *PaginatedMusicLookup) getTotalCount(db *pgxpool.Pool) error {
	sql := "select count(*)/@limit, count(*) from all_tracks"

	row := db.QueryRow(context.Background(), sql, pgx.NamedArgs{
		"limit": PAGINATED_COUNT,
	})

	return row.Scan(&m.FullLength, &m.Count)
}

func (m *PaginatedMusicLookup) fromRow(rows pgx.Rows) error {
	for rows.Next() {
		entry := Music{}
		if err := entry.fromRow(rows); err != nil {
			return err
		}
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
	defer row.Close()

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
	_, err := db.Exec(context.Background(), "update music set title=@title, artist_id=@artistId, release_date=@releasedate WHERE id=@id", pgx.NamedArgs{
		"title":       m.Tracktitle,
		"artistId":    m.ArtistId,
		"releasedate": m.ReleaseDate,
		"id":          m.TrackId,
	})

	return err
}

func (item *Music) Delete(db *pgxpool.Pool) error {
	tx, err := db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	err = s3store.DeleteTrackFromArchive(item.TrackId + ".wav")
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), "delete from music where id = @id", pgx.NamedArgs{
		"id": item.TrackId,
	})
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}
