package notification

import (
	"context"
	"database/sql"
	db "go-metrics-service/db/sqlc"
)

type RepositoryInterface interface {
	InsertNotification(ctx context.Context, arg db.InsertNotificationParams) error
}

type Repository struct {
	Conn    *sql.DB
	DBtx    db.DBTX
	Queries *db.Queries
	SqlConn *sql.DB
}

func NewRepository(conn *sql.DB) *Repository {
	q := db.New(conn)
	return &Repository{
		Conn:    conn,
		DBtx:    conn,
		Queries: q,
		SqlConn: conn,
	}
}

func (r *Repository) InsertNotification(ctx context.Context, arg db.InsertNotificationParams) error {
	return r.Queries.InsertNotification(ctx, arg)
}
