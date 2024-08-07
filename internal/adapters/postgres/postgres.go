package postgres

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	log "github.com/sirupsen/logrus"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
)

type UserAuthStorage struct {
	db *pgx.Conn
}

func NewUserAuthStorage(dataSource string) (repository.IUserAuthStorage, error) {
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

func (r *UserAuthStorage) CreateUser(user request.Register) (userID string, err error) {
	query := `INSERT INTO "user" (
                    PASSWORD, 
                    PHONE_NUMBER, 
                    EMAIL, 
                    TELEGRAM, 
                    VK, 
                    "group", 
                    NAME
                    ) VALUES (
                              $1, $2, $3, $4, $5, $6, $7
                    ) RETURNING ID;
`
	err = r.db.QueryRow(context.TODO(), query,
		user.Password,
		user.PhoneNumber,
		user.Email,
		user.Telegram,
		user.VK,
		user.Group,
		user.Name).Scan(&userID)
	if err != nil {
		log.WithField(
			"origin.function", "CreateUser",
		).Errorf("Ошибка при создании пользователя: %s", err.Error())
		return "", err
	}

	return userID, nil
}

func (r *UserAuthStorage) GetUserPassword(email string) (password string, err error) {
	query := `SELECT PASSWORD FROM "user" WHERE EMAIL = $1;`

	err = r.db.QueryRow(context.TODO(), query, email).Scan(&password)
	if err != nil {
		log.WithField(
			"origin.function", "GetUserPassword",
		).Errorf("Ошибка при получении пароля пользователя: %s", err.Error())
		return "", err
	}

	return password, nil
}

func (r *UserAuthStorage) GetUserID(email string) (userID string, err error) {
	query := `SELECT ID FROM "user" WHERE EMAIL = $1;`

	err = r.db.QueryRow(context.TODO(), query, email).Scan(&userID)
	if err != nil {
		log.WithField(
			"origin.function", "GetUserID",
		).Errorf("Ошибка при получении идентификатора пользователя: %s", err.Error())
		return "", err
	}

	return userID, nil
}

func (r *UserAuthStorage) CheckUser(email string) (exists bool, err error) {
	query := `SELECT EXISTS (SELECT 1 FROM "user" WHERE EMAIL = $1)`

	err = r.db.QueryRow(context.TODO(), query, email).Scan(&exists)
	if err != nil {
		log.WithField(
			"origin.function", "CheckUser",
		).Errorf("Ошибка при проверке существования пользователя: %s", err.Error())
		return false, err
	}

	return exists, nil
}

func (r *UserAuthStorage) GetUserProfile(userID string) (*response.UserProfile, error) {
	query := `	SELECT 
					NAME, 
					"group", 
					TELEGRAM, 
					VK, 
					PHONE_NUMBER, 
					EMAIL, 
					TEAM_ID
				FROM "user" 
					WHERE ID = $1;
`
	var profile response.UserProfile
	var s sql.NullString
	res := r.db.QueryRow(context.TODO(), query, userID)
	err := res.Scan(&profile.Name, &profile.Group, &profile.Telegram, &profile.VK, &profile.PhoneNumber, &profile.Email, &s)
	if s.Valid {
		profile.TeamID = s.String
	} else {
		profile.TeamID = ""
	}

	if err != nil {
		log.WithField(
			"origin.function", "GetUserProfile",
		).Errorf("Ошибка при получении данных пользователя: %s", err.Error())
		return nil, err
	}

	profile.ID = userID

	log.Printf("profile: %+v", profile)

	return &profile, nil
}

func (r *UserAuthStorage) ChangeUserProfile(userID string, profile *request.ChangeProfile) error {
	if profile.Name != "" {
		query := `UPDATE "user" SET NAME=$2 WHERE ID=$1`

		_, err := r.db.Exec(context.TODO(), query, userID, profile.Name)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении имени пользователя: %s", err.Error())
			return err
		}
	}

	if profile.Group != "" {
		query := `UPDATE "user" SET "group"=$2 WHERE ID=$1`

		_, err := r.db.Exec(context.TODO(), query, userID, profile.Group)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении группы пользователя: %s", err.Error())
			return err
		}
	}

	if profile.Password != "" {
		query := `UPDATE "user" SET PASSWORD=$2 WHERE ID=$1`

		_, err := r.db.Exec(context.TODO(), query, userID, profile.Password)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении паоля пользователя: %s", err.Error())
			return err
		}
	}

	if profile.Telegram != "" {
		query := `UPDATE "user" SET TELEGRAM=$2 WHERE ID=$1`

		_, err := r.db.Exec(context.TODO(), query, userID, profile.Telegram)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении ТГ пользователя: %s", err.Error())
			return err
		}
	}

	if profile.VK != "" {
		query := `UPDATE "user" SET VK=$2 WHERE ID=$1`

		_, err := r.db.Exec(context.TODO(), query, userID, profile.VK)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении ВК пользователя: %s", err.Error())
			return err
		}
	}

	if profile.Email != "" {
		query := `UPDATE "user" SET EMAIL=$2 WHERE ID=$1`

		_, err := r.db.Exec(context.TODO(), query, userID, profile.Email)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении почты пользователя: %s", err.Error())
			return err
		}
	}

	if profile.PhoneNumber != "" {
		query := `UPDATE "user" SET PHONE_NUMBER=$2 WHERE ID=$1`

		_, err := r.db.Exec(context.TODO(), query, userID, profile.PhoneNumber)

		if err != nil {
			log.WithField(
				"origin.function", "ChangeUserProfile",
			).Errorf("Ошибка при изменении телефона пользователя: %s", err.Error())
			return err
		}
	}

	return nil
}

func (r *UserAuthStorage) CheckUserExist(userID string) (bool, error) {
	var exist bool
	query := `SELECT EXISTS (SELECT ID FROM "user" WHERE ID = $1)`
	err := r.db.QueryRow(context.TODO(), query, userID).Scan(&exist)
	if err != nil {
		log.WithField(
			"origin.function", "CheckMembership",
		).Errorf("Ошибка при проверке существования участника в бд: %s", err.Error())
		return false, err
	}

	return exist, nil

}

func (r *UserAuthStorage) GetUserPasswordById(userID string) (password string, err error) {
	query := `SELECT PASSWORD FROM "user" WHERE ID=$1;`

	err = r.db.QueryRow(context.TODO(), query, userID).Scan(&password)
	if err != nil {
		log.WithField(
			"origin.function", "GetUserPasswordById",
		).Errorf("Ошибка при получении пароля пользователя с помощью ID: %s", err.Error())
		return "", err
	}

	return password, nil
}

func (r *UserAuthStorage) ChangeUserPassword(userID string, newPassword string) error {
	query := `UPDATE "user" SET PASSWORD=$2 WHERE ID=$1;`
	_, err := r.db.Exec(context.TODO(), query, userID, newPassword)

	if err != nil {
		log.WithField(
			"origin.function", "ChangeUserPassword",
		).Errorf("Ошибка при изменении пароля пользователя: %s", err.Error())
		return err
	}

	return nil
}
