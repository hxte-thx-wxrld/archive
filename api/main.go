package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitApi(rg *gin.RouterGroup, db *pgxpool.Pool) {
	TrackApi(rg, db)
	ArtistApi(rg, db)
	ReleaseApi(rg, db)
}
