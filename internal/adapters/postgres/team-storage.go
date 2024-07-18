package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
)

type TeamStorage struct {
	db *pgx.Conn
}

func NewTeamStorage(dataSource string) (repository.IUserAuthStorage, error) {
	db, err := pgx.Connect(context.Background(), dataSource)
	if err != nil {
		return nil, err
	}

	// db.DB.SetMaxOpenConns(1000) // The default is 0 (unlimited)
	// db.DB.SetMaxIdleConns(10)   // defaultMaxIdleConns = 2
	// db.DB.SetConnMaxLifetime(0) // 0, connections are reused forever.
	return &UserAuthStorage{
		db: db,
	}, err
}
