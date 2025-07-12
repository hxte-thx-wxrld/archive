package model

import (
	"context"
	"mime/multipart"

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
	Position   int
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

	defer rows.Close()

	err = r.getTotalCount(db)
	if err != nil {
		return err
	}

	err = r.fromRow(db, rows)
	if err != nil {
		return err
	}
	rows.Close()
	return err
}

func (r *ReleasesLookup) fromRow(db *pgxpool.Pool, rows pgx.Rows) error {
	for rows.Next() {
		entry := Release{}
		entry.fromRow(rows)
		entry.getAssociatedMusic(db)
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

func (r *Release) FromMultipartForm(db *pgxpool.Pool, form *multipart.Form) {
	r.Name = form.Value["Name"][0]
	r.ReleaseDate = form.Value["ReleaseDate"][0]
	r.Isrc = &form.Value["Isrc"][0]

	tl := form.Value["tracklist"]
	for i, v := range tl {
		var m MusicInRelease
		m.TrackId = v
		m.Position = i
		r.RelatedMusic = append(r.RelatedMusic, m)
	}

}

func (r *Release) Create(db *pgxpool.Pool) error {
	row := db.QueryRow(context.Background(), "INSERT INTO public.releases (name)VALUES(@name) returning id", pgx.NamedArgs{
		"name": r.Name,
	})

	return row.Scan(&r.ReleaseId)
}

func (r *Release) Edit(db *pgxpool.Pool) error {
	// UPDATE releases SET "name"='test', release_date=@ReleaseDate, isrc=@isrc WHERE id=@ReleaseId;
	tx, err := db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), "UPDATE releases SET name=@title, release_date=@releasedate, isrc=@isrc WHERE id=@id", pgx.NamedArgs{
		"title":       r.Name,
		"isrc":        r.Isrc,
		"releasedate": r.ReleaseDate,
		"id":          r.ReleaseId,
	})
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	_, err = tx.Exec(context.Background(), "DELETE FROM public.music_in_releases WHERE release_id=@id", pgx.NamedArgs{
		"id": r.ReleaseId,
	})
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	for _, row := range r.RelatedMusic {
		_, err = tx.Exec(context.Background(), "INSERT INTO music_in_releases (music_id, position, release_id) VALUES(@music_id, @position, @release_id);", pgx.NamedArgs{
			"music_id":   row.TrackId,
			"position":   row.Position,
			"release_id": r.ReleaseId,
		})
		if err != nil {
			tx.Rollback(context.Background())
			return err
		}
	}

	return tx.Commit(context.Background())

	//return row.Scan()
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

	row.Close()

	return err
}

func (r *Release) fromRow(row pgx.Rows) error {
	return row.Scan(&r.Name, &r.ReleaseDate, &r.Isrc, &r.CoverUrl, &r.ReleaseId, &r.CatalogId)
}

func (r *Release) getAssociatedMusic(db *pgxpool.Pool) error {
	rows, err := db.Query(context.Background(),
		"select m.id, m.title, mr.position, i.name from music_in_releases mr join music m on m.id = mr.music_id join interpret i on m.artist_id = i.id where release_id = @releaseId order by mr.position", pgx.NamedArgs{
			"releaseId": r.ReleaseId,
		})

	if err != nil {
		return err
	}

	r.RelatedMusic = []MusicInRelease{}
	for rows.Next() {
		t := MusicInRelease{}
		err = t.fromRow(rows)
		if err != nil {
			return err
		}

		r.RelatedMusic = append(r.RelatedMusic, t)
	}

	rows.Close()
	return nil
}

func (r *MusicInRelease) fromRow(rows pgx.Rows) error {
	return rows.Scan(&r.TrackId, &r.Name, &r.Position, &r.ArtistName)
}
