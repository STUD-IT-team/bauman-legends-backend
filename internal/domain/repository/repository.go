package repository

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	_ "github.com/jackc/pgx"
)

type IUserAuthStorage interface {
	CreateUser(user request.Register) (userID string, err error)
	GetUserPassword(email string) (password string, err error)
	GetUserID(email string) (userID string, err error)
	CheckUser(email string) (exists bool, err error)
}
