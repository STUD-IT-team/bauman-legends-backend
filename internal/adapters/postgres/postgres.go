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
	query := `	
update "user"
	    set name=$2,
	        password=$3,
	        "group"=$4,
	        telegram=$5,
	        vk=$6,
	        phone_number=$7,
	        email=$8
	    where id = $1;
`
	_, err := r.db.Exec(query,
		userID,
		profile.Name,
		profile.Password,
		profile.Group,
		profile.Telegram,
		profile.VK,
		profile.PhoneNumber,
		profile.Email,
	)

	if err != nil {
		log.WithField(
			"origin.function", "ChangeUserProfile",
		).Errorf("Ошибка при изменении данных пользователя: %s", err.Error())
		return err
	}

	return nil
}

func (r *UserAuthStorage) CheckTeam(team request.RegisterTeam) (exists bool, err error) {
	query := `select exists (select 1 from "team" where title = $1)`

	err = r.db.Get(&exists, query, team.TeamName)
	if err != nil {
		log.WithField(
			"origin.function", "CheckTeam",
		).Errorf("Ошибка при проверке существования команды: %s", err.Error())
		return false, err
	}

	return exists, nil
}

func (r *UserAuthStorage) CreateTeam(team request.RegisterTeam) (TeamID string, err error) {
	query := `insert into "team" (
                    title
                    ) values (
                              $1
                    ) returning id;
`
	err = r.db.Get(&TeamID, query, team.TeamName)
	if err != nil {
		log.WithField(
			"origin.function", "CreateTeam",
		).Errorf("Ошибка при создании команды: %s", err.Error())
		return "", err
	}

	return TeamID, nil
}
