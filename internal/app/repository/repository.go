package repository

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	CreateUser(user domain.User) (userID string)
	// TODO
	// 	GetUser(id)
}

type Repository struct {
	db *sqlx.DB
}

func (r *Repository) CreateUser(userID domain.User) string {
	// TODO
	// 	 Написать запрос добавления пользователя с возвращением его id,
	//	 отредактировать тип первичного ключа в sql файле
	return ""
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
