package postgres

import (
	"database/sql"
	"fmt"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
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

func NewTaskStorage(dataSource string) (repository.TaskStorage, error) {
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

func (r *UserAuthStorage) UpdateTeam(teamID string, teamName string) error {
	log.Infof("teamID and TeamName in UpdateTeam: %s, %s", teamID, teamName)
	query := `	
		update "team"
	    set title=$1
	    where id = $2;
`
	_, err := r.db.Exec(query,
		teamName,
		teamID,
	)

	if err != nil {
		log.WithField(
			"origin.function", "ChangeUserProfile",
		).Errorf("Ошибка при изменении данных команды: %s", err.Error())
		return err
	}

	return nil
}

func (r *UserAuthStorage) GetTeam(teamID string) (domain.Team, error) {
	log.Infof("team from GetTeam: %s", teamID)
	var team domain.Team
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
		log.Infof("members: %+v", mems)
		if err != nil {
			return domain.Team{}, err
		}
		team.Members = append(team.Members, mem)
	}
	//log.Infof("участники команды:%+v", members)
	//copy(team.Members, members)
	log.Infof("участники команды пониже:%+v", team.Members)
	return team, nil
}

func (r *UserAuthStorage) GetTeamPoints(teamID string) (int, error) {
	var points int
	var res = true
	err := r.db.QueryRow(`select SUM(task.points) from team_task left join task on team_task.task_id = task.id 
                        where team_task.team_id = $1 and result = $2`, teamID, res).Scan(&points)
	if err != nil {
		return 0, err
	}
	return points, nil
}

func (r *UserAuthStorage) DeleteTeam(TeamID string) error {
	_, err := r.db.Exec(`update "user" set team_id=null, role_id=null where team_id = $1;`, TeamID)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(`delete from "team" where id=$1;`, TeamID)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserAuthStorage) InviteToTeam(UserID string, TeamID string) error {

	_, err := r.db.Exec(`update "user" set team_id=$1, role_id = 1 where id = $2;`, TeamID, UserID)

	if err != nil {
		return err
	}
	return nil
}
func (r *UserAuthStorage) DeleteFromTeam(UserID string, TeamID string) error {
	_, err := r.db.Exec(`update "user" set team_id=null where id = $2 and team_id =$1;`, TeamID, UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserAuthStorage) UpdateMember(UserID string, RoleID int) error {
	_, err := r.db.Exec(`update "user" set role_id=$1 where id = $2;`, RoleID, UserID)
	if err != nil {
		return err
	}
	return nil

}

func (r *UserAuthStorage) SetTeamID(UserID string, teamID string) error {

	_, err := r.db.Exec(`update "user" set team_id = $1, role_id = 3 where id = $2;`, teamID, UserID)
	return err
}

func (r *UserAuthStorage) CheckMembership(userId string, teamID string) (bool, error) {
	var exists bool
	query := `select exists (select id from "user" where team_id = $1 and id = $2)`
	err := r.db.Get(&exists, query, teamID, userId)
	if err != nil {
		log.WithField(
			"origin.function", "CheckMembership",
		).Errorf("Ошибка при проверке существования участника команды: %s", err.Error())
		return false, err
	}

	return exists, nil
}

func (r *UserAuthStorage) CheckUserExist(userID string) (bool, error) {
	var exist bool
	query := `select exists (select id from "user" where id = $1)`
	err := r.db.Get(&exist, query, userID)
	if err != nil {
		log.WithField(
			"origin.function", "CheckMembership",
		).Errorf("Ошибка при проверке существования участника в бд: %s", err.Error())
		return false, err
	}

	return exist, nil

}

const queryGetTaskTypes = `
	select task.type_id, count(task.type_id) 
    from team_task 
    left join task on team_task.task_id = task.id 
    where team_id=$1 
    group by team_id, task.type_id;
`

func (r *UserAuthStorage) GetTaskTypes(teamID string) (domain.TaskTypes, error) {
	var taskTypes domain.TaskTypes
	rows, err := r.db.Query(queryGetTaskTypes, teamID)
	if err != nil {
		return nil, fmt.Errorf("can't db.Query on GetTaskTypes: %w", err)
	}

	for rows.Next() {
		var out domain.TaskType
		err = rows.Scan(&out.ID, &out.Count)
		if err != nil {
			return nil, err
		}
		err = r.db.QueryRow(`select title from team_task where id = $1`, out.ID).Scan(&out.Title)
		if err != nil {
			return nil, err
		}
		out.IsActive = true
		taskTypes = append(taskTypes, out)
	}
	return taskTypes, nil
}

func (r *UserAuthStorage) GetTaskAmount(taskTypeID int) (int, error) {
	var amount int

	err := r.db.QueryRow(`select count(*) from task where type_id = $1`, taskTypeID).Scan(&amount)
	if err != nil {
		return 0, err
	}
	return amount, nil
}

func (r *UserAuthStorage) GetBusyNocPlaceses() (int, error) {
	var busyNoc int
	err := r.db.QueryRow(`select count(*) from team_task where type_id = 2;`).Scan(&busyNoc)
	if err != nil {
		return 0, err
	}
	return busyNoc, nil
}

func (r *UserAuthStorage) GetTeamTaskAmount(taskTypeID int) (int, error) {
	var teamAmount int
	err := r.db.QueryRow(`select count(*) from team_task where type_id = $1`, taskTypeID).Scan(&teamAmount)
	if err != nil {
		return 0, err
	}
	return teamAmount, nil
}

func (r *UserAuthStorage) GetAvailableTaskID(teamID string, taskTypeID int) ([]string, error) {
	var taskIDs []string
	rows, err := r.db.Query(`select task.id from team_task 
    				  right join task on team_task.task_id = task.id 
               		  where team_id = $1 and task.type_id = $2 and team_task.task_id is NULL`, teamID, taskTypeID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var taskID string
		err = rows.Scan(&taskID)
		if err != nil {
			return nil, err
		}
		taskIDs = append(taskIDs, taskID)
	}

	return taskIDs, nil
}

const querySetTaskToTeam = `insert into team_task`

func (r *UserAuthStorage) SetTaskToTeam(taskID string, teamID string) error {
	_, err := r.db.Exec(`insert into team_task(task_id, team_id) values ($1, $2)`, taskID, teamID)
	return err
}

func (r *UserAuthStorage) CheckActiveTaskExist(teamID string) (bool, error) {
	var exists bool
	query := `select exists (select task_id from team_task where team_id = $1 and
              answer_text is null and answerImageBase64 is null and result is null)` //answerText, answerImageUrl, result
	err := r.db.Get(&exists, query, teamID)
	if err != nil {
		log.WithField(
			"origin.function", "CheckActiveTaskExist",
		).Errorf("Ошибка при проверке существования у команды активной задачи: %s", err.Error())
		return false, err
	}

	return exists, nil
}

func (r *UserAuthStorage) GetActiveTaskID(teamID string) (string, error) {
	var taskID string
	err := r.db.QueryRow(`select task_id from team_task where team_id = $1 and
		answer_text is null and answerImageBase64 is null and result is null`, teamID).Scan(&taskID)
	if err != nil {
		return "", err
	}
	return taskID, nil
}

func (r *UserAuthStorage) GetTask(taskID string) (domain.Task, error) {
	var out domain.Task
	err := r.db.Get(&out, `select id, title, description, type_id, 
       					  max_points, min_points, answer_type_id from task where id=$1`)
	if err != nil {
		return domain.Task{}, err
	}
	return domain.Task{}, nil
}

func (r *UserAuthStorage) GetTaskStartedTime(taskID string, TeamID string) (time.Time, error) {
	var out time.Time

	err := r.db.Get(&out, `select start_time from team_task where task_id = $1 and team_id = $2`, taskID, TeamID)
	if err != nil {
		return time.Time{}, err
	}

	return out, nil
}

func (r *UserAuthStorage) GetTaskTypeName(taskID string) (string, error) {
	var out string

	err := r.db.Get(&out, `select task_type.title from task left join task_type on task.type_id = task_type.id where task_id = $1`, taskID)
	if err != nil {
		return "", err
	}

	return out, nil
}

func (r *UserAuthStorage) SetActiveTaskExpired(taskID string, TeamID string) error {
	_, err := r.db.Exec(`update team_task set result = false where task_id = $1 and team_id = $2`, taskID, TeamID)
	return err
}

func (r *UserAuthStorage) SetAnswerText(text string, teamID string, taskID string) error {
	_, err := r.db.Exec(`update team_task set answer_text = $1 where team_id = $2 and task_id = $3`, text)
	return err
}
func (r *UserAuthStorage) SetAnswerImageBase64(url string, teamID string, taskID string) error {
	_, err := r.db.Exec(`update team_task set answerImageBase64 = $1 where team_id = $2 and task_id = $3`, url)
	return err
}

func (r *UserAuthStorage) GetAnswers(teamID string) ([]domain.Answer, error) {
	var answers []domain.Answer
	err := r.db.Select(&answers, `select * from team_task where team_id = $1`, teamID)
	if err != nil {
		return make([]domain.Answer, 0), err
	}
	return answers, nil
}
