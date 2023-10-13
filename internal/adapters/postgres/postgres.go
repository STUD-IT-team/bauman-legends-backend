package postgres

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type UserAuthStorage struct {
	db *sqlx.DB
}

func NewUserAuthStorage(dataSource string) (repository.IUserAuthStorage, error) {
	db, err := sqlx.Open("pgx", dataSource)
	if err != nil {
		return nil, err
	}
	return &UserAuthStorage{
		db: db,
	}, err
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
	err = r.db.Get(&userID, query,
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

	err = r.db.Get(&password, query, email)
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

	err = r.db.Get(&userID, query, email)
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

	err = r.db.Get(&exists, query, email)
	if err != nil {
		log.WithField(
			"origin.function", "CheckUser",
		).Errorf("Ошибка при проверке существования пользователя: %s", err.Error())
		return false, err
	}

	return exists, nil
}

func (r *UserAuthStorage) GetUserProfile(userID string) (*response.UserProfile, error) {
	query := `	select 
					name, 
					"group", 
					telegram, 
					vk, 
					phone_number, 
					email 
				from "user" 
					where id = $1;
`
	var profile response.UserProfile
	err := r.db.Get(&profile, query, userID)

	if err != nil {
		log.WithField(
			"origin.function", "GetUserProfile",
		).Errorf("Ошибка при получении данных пользователя: %s", err.Error())
		return nil, err
	}

	return &profile, nil
}

func (r *UserAuthStorage) ChangeUserProfile(userID string, profile *request.ChangeProfile) error {
	if profile.Name != "" {
		query := `update "user" set name=$2 where id=$1`

		_, err := r.db.Exec(query, userID, profile.Name)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении имени пользователя: %s", err.Error())
			return err
		}
	}

	if profile.Group != "" {
		query := `update "user" set "group"=$2 where id=$1`

		_, err := r.db.Exec(query, userID, profile.Group)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении группы пользователя: %s", err.Error())
			return err
		}
	}

	if profile.Password != "" {
		query := `update "user" set password=$2 where id=$1`

		_, err := r.db.Exec(query, userID, profile.Password)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении паоля пользователя: %s", err.Error())
			return err
		}
	}

	if profile.Telegram != "" {
		query := `update "user" set telegram=$2 where id=$1`

		_, err := r.db.Exec(query, userID, profile.Telegram)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении ТГ пользователя: %s", err.Error())
			return err
		}
	}

	if profile.VK != "" {
		query := `update "user" set vk=$2 where id=$1`

		_, err := r.db.Exec(query, userID, profile.VK)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении ВК пользователя: %s", err.Error())
			return err
		}
	}

	if profile.Email != "" {
		query := `update "user" set email=$2 where id=$1`

		_, err := r.db.Exec(query, userID, profile.Email)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении почты пользователя: %s", err.Error())
			return err
		}
	}

	if profile.PhoneNumber != "" {
		query := `update "user" set phone_number=$2 where id=$1`

		_, err := r.db.Exec(query, userID, profile.PhoneNumber)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении телефона пользователя: %s", err.Error())
			return err
		}
	}

	return nil
}
