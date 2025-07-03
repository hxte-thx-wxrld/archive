package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func InitApi(rg *gin.RouterGroup, db *pgx.Conn) {
	TrackApi(rg, db)
	ArtistApi(rg, db)
	ReleaseApi(rg, db)
}
