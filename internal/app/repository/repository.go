package repository

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	CreateUser(user request.Register) (userID string, err error)
	GetUserPassword(email string) (password string, err error)
	GetUserID(email string) (userID string, err error)
	CheckUser(email string) (exists bool, err error)
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(dataSource string) IRepository {
	db, err := sqlx.Open("pgx", dataSource)
	if err != nil {
		return nil
	}
	return &Repository{
		db: db,
	}
}
