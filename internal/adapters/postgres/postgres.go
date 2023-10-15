package postgres

import (
	"database/sql"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
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

func NewTeamStorage(dataSource string) (repository.TeamStorage, error) {
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
					email, 
					team_id
				from "user" 
					where id = $1;
`
	var profile response.UserProfile
	var s sql.NullString
	res := r.db.QueryRow(query, userID)
	err := res.Scan(&profile.Name, &profile.Group, &profile.Telegram, &profile.VK, &profile.PhoneNumber, &profile.Email, &s)
	if s.Valid {
		profile.TeamID = s.String
	} else {
		profile.TeamID = ""
	}
	log.Println(profile)
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

func (r *UserAuthStorage) CheckTeam(teamName string) (exists bool, err error) {
	query := `select exists (select 1 from "team" where title = $1)`

	err = r.db.Get(&exists, query, teamName)
	if err != nil {
		log.WithField(
			"origin.function", "CheckTeam",
		).Errorf("Ошибка при проверке существования команды: %s", err.Error())
		return false, err
	}

	return exists, nil
}

func (r *UserAuthStorage) CreateTeam(teamName string) (string, error) {
	query := `insert into team (
                    title
                    ) values (
                              $1
                    ) returning id;`
	var teamID string
	err := r.db.QueryRow(query, teamName).Scan(&teamID)
	if err != nil {
		log.WithField(
			"origin.function", "CreateTeam",
		).Errorf("Ошибка при создании команды: %s", err.Error())
		return "", err
	}
	return teamID, nil
}

func (r *UserAuthStorage) UpdateTeam(teamName string) error {
	query := `	
		update "team"
	    set title=$2
	    where title = $1;
`
	_, err := r.db.Exec(query,
		teamName,
		teamName,
	)

	if err != nil {
		log.WithField(
			"origin.function", "ChangeUserProfile",
		).Errorf("Ошибка при изменении данных пользователя: %s", err.Error())
		return err
	}

	return nil
}

func (r *UserAuthStorage) GetTeam(teamID string) (domain.Team, error) {
	log.Infof("team from GetTeam: %s", teamID)
	var team domain.Team
	var members []domain.Member
	query := `select id, title from "team" where id = $1;`
	err := r.db.QueryRow(query, teamID).Scan(&team.TeamId, &team.Title)
	log.Infof("team from db: %s:%s", team.TeamId, team.Title)
	if err != nil {
		log.WithField(
			"origin.function", "GetTeamPG",
		).Errorf("Ошибка при попытке достать данные команды: %s", err.Error())
		return domain.Team{}, err
	}
	mems, err := r.db.Query(`select id, name, role_id from "user" where team_id = $1;`, teamID)
	for mems.Next() {
		var mem domain.Member
		err = mems.Scan(&mem.Id, &mem.Name, &mem.Role)
		if err != nil {
			return domain.Team{}, err
		}
		members = append(members, mem)
	}
	copy(team.Members, members)
	return team, nil
}

func (r *UserAuthStorage) DeleteTeam(TeamID string) error {
	_, err := r.db.Exec(`update "user" set team_id=null, role_id=null where team_id = $1;`, TeamID)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(`delete from "team" where team_id=$1;`, TeamID)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserAuthStorage) InviteToTeam(UserID string, TeamID string) error {
	_, err := r.db.Exec(`update "user" set team_id=$1, where id = $2;`, TeamID, UserID)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserAuthStorage) DeleteFromTeam(UserID string, TeamID string) error {
	_, err := r.db.Exec(`update "user" set team_id=null, where id = $2 and team_id =$1;`, TeamID, UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserAuthStorage) UpdateMember(UserID string, RoleID int) error {
	_, err := r.db.Exec(`update "user" set role_id=$1, where id = $2;`, RoleID, UserID)
	if err != nil {
		return err
	}
	return nil

}

func (r *UserAuthStorage) SetTeamID(UserID string, teamID string) error {
	_, err := r.db.Exec(`update "user" set team_id = $1 where id = $2;`, teamID, UserID)
	return err
}
