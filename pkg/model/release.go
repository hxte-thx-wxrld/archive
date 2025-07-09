package model

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

type ReleasesLookup struct {
	Rows       []Release
	FullLength int
}

func (r *ReleasesLookup) AllReleases(db *pgxpool.Pool, offset int) error {
	rows, err := db.Query(context.Background(), "select name, release_date::text, isrc, public_cover_url, id, catalog_id from releases order by release_date desc limit @limit offset @offset", pgx.NamedArgs{
		"limit":  PAGINATED_COUNT,
		"offset": offset * PAGINATED_COUNT,
	})

	if err != nil {
		return err
	}

	err = r.getTotalCount(db)
	if err != nil {
		return err
	}

	return r.fromRow(db, rows)
}

func (r *ReleasesLookup) fromRow(db *pgxpool.Pool, rows pgx.Rows) error {
	for rows.Next() {
		entry := Release{}
		entry.fromRow(rows)
		//entry.getAssociatedMusic(db)
		r.Rows = append(r.Rows, entry)
	}
	return nil
}

func (m *ReleasesLookup) getTotalCount(db *pgxpool.Pool) error {
	sql := "select count(*)/@limit from releases"

	row := db.QueryRow(context.Background(), sql, pgx.NamedArgs{
		"limit": PAGINATED_COUNT,
	})

	return row.Scan(&m.FullLength)
}

func (r *Release) FromId(db *pgxpool.Pool, id string) error {
	row, err := db.Query(context.Background(), "select name, release_date::text, isrc, public_cover_url, id, catalog_id from releases where id = @id", pgx.NamedArgs{
		"id": id,
	})

	if err != nil {
		return err
	}

	row.Next()
	err = r.fromRow(row)
	if err != nil {
		return err
	}

	err = r.getAssociatedMusic(db)
	if err != nil {
		return err
	}

	return err
}

func (r *Release) fromRow(row pgx.Rows) error {
	return row.Scan(&r.Name, &r.ReleaseDate, &r.Isrc, &r.CoverUrl, &r.ReleaseId, &r.CatalogId)
}

func (r *Release) getAssociatedMusic(db *pgxpool.Pool) error {
	rows, err := db.Query(context.Background(),
		"select m.id, m.title, mr.order, i.name from music_in_releases mr join music m on m.id = mr.music_id join interpret i on m.artist_id = i.id where release_id = @releaseId order by mr.order", pgx.NamedArgs{
			"releaseId": r.ReleaseId,
		})

	if err != nil {
		return err
	}

	for rows.Next() {
		t := MusicInRelease{}
		err = t.fromRow(rows)
		if err != nil {
			return err
		}

		r.RelatedMusic = append(r.RelatedMusic, t)
	}
	return nil
}

func (r *MusicInRelease) fromRow(rows pgx.Rows) error {
	return rows.Scan(&r.TrackId, &r.Name, &r.Order, &r.ArtistName)
}
