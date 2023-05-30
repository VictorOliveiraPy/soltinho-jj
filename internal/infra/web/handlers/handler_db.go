package handlers

import (
	"database/sql"
	db "github.com/VictorOliveiraPy/internal/infra/database"
)

type EntityHandler struct {
	gymDB *sql.DB
	*db.Queries
}

func NewEntityHandler(dbConn *sql.DB) *EntityHandler {
	return &EntityHandler{
		gymDB:   dbConn,
		Queries: db.New(dbConn),
	}
}
