package model

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Analysis struct {
	TrackId  string
	Tempo    float64
	ZCR      float64
	RMS      float64
	Centroid float64
	Rolloff  float64
	Flatness float64
	Duration int
}

func (a *Analysis) ForTrackId(db *pgxpool.Pool) error {
	row := db.QueryRow(context.Background(), "select trackid, tempo, zcr, rms, centroid, rolloff, flatness, duration from analysis where trackid = @id", pgx.NamedArgs{
		"id": a.TrackId,
	})
	return a.fromRow(row)
}

func (a *Analysis) fromRow(row pgx.Row) error {
	return row.Scan(&a.TrackId, &a.Tempo, &a.ZCR, &a.RMS, &a.Centroid, &a.Rolloff, &a.Flatness, &a.Duration)
}
