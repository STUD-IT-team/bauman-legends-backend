package postgres

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type UserAuthStorage struct {
	db *sqlx.DB
}

func NewUserAuthStorage(dataSource string) repository.IUserAuthStorage {
	db, err := sqlx.Open("pgx", dataSource)
	if err != nil {
		return nil
	}
	return &UserAuthStorage{
		db: db,
	}
}

func (r *UserAuthStorage) CreateUser(user request.Register) (userID string, err error) {
	query := `insert into "user" (
                    password, 
                    phone_number, 
                    email, 
                    telegram, 
                    vk, 
                    "group", 
                    name
                    ) values (
                              $1, $2, $3, $4, $5, $6, $7
                    ) returning id;
`
	err = r.db.Select(&userID, query,
		user.Password,
		user.PhoneNumber,
		user.Email,
		user.Telegram,
		user.VK,
		user.Group,
		user.Name)
	if err != nil {
		log.WithField(
			"origin.function", "CreateUser",
		).Errorf("Ошибка при создании пользователя: %s", err.Error())
		return "", err
	}

	return userID, nil
}

func (r *UserAuthStorage) GetUserPassword(email string) (password string, err error) {
	query := `select password from "user" where email = $1;`

	err = r.db.Select(&password, query, email)
	if err != nil {
		log.WithField(
			"origin.function", "GetUserPassword",
		).Errorf("Ошибка при получении пароля пользователя: %s", err.Error())
		return "", err
	}

	return password, nil
}

func (r *UserAuthStorage) GetUserID(email string) (userID string, err error) {
	query := `select id from "user" where email = $1;`

	err = r.db.Select(&userID, query, email)
	if err != nil {
		log.WithField(
			"origin.function", "GetUserID",
		).Errorf("Ошибка при получении идентификатора пользователя: %s", err.Error())
		return "", err
	}

	return userID, nil
}

func (r *UserAuthStorage) CheckUser(email string) (exists bool, err error) {
	query := `select exists (select 1 from "user" where email = $1)`

	err = r.db.Select(&exists, query, email)
	if err != nil {
		log.WithField(
			"origin.function", "CheckUser",
		).Errorf("Ошибка при проверке существования пользователя: %s", err.Error())
		return false, err
	}

	return exists, nil
}
