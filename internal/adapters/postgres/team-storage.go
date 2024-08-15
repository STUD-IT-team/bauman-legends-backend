package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/storage"
)

type TeamStorage struct {
	db *pgx.Conn
}

func NewTeamStorage(dataSource string) (storage.TeamStorage, error) {
	config, err := pgx.ParseConfig(dataSource)
	if err != nil {
		return nil, err
	}

	db, err := pgx.ConnectConfig(context.Background(), config)

	// db.DB.SetMaxOpenConns(1000) // The default is 0 (unlimited)
	// db.DB.SetMaxIdleConns(10)   // defaultMaxIdleConns = 2
	// db.DB.SetConnMaxLifetime(0) // 0, connections are reused forever.
	return &TeamStorage{
		db: db,
	}, err
}

func (s *TeamStorage) CheckTeam(teamName string) (exists bool, err error) {
	err = s.db.QueryRow(context.TODO(), checkTeamQuery, teamName).Scan(&exists)
	if err != nil {
		log.WithField(
			"origin.function", "CheckTeam",
		).Errorf("Ошибка при проверке существования команды: %s", err.Error())
		return false, errors.Join(consts.InternalServerError, err)
	}

	return exists, nil
}

func (s *TeamStorage) CreateTeam(teamName string, userId int) (teamId int, err error) {
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		log.WithField(
			"origin.function", "CreateTeam",
		).Errorf("Ошибка при создании транзакции: %s", err.Error())
		return -1, errors.Join(consts.InternalServerError, err)
	}

	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.TODO(), createTeamQuery, teamName).Scan(&teamId)
	if err != nil {
		log.WithField(
			"origin.function", "CreateTeam",
		).Errorf("Ошибка при создании команды: %s", err.Error())
		return -1, errors.Join(consts.InternalServerError, err)
	}

	_, err = tx.Exec(context.TODO(), setTeamIdQuery, teamId, userId)
	if err != nil {
		log.WithField(
			"origin.function", "CreateTeam",
		).Errorf("Ошибка при назначении команды участнику: %s", err.Error())
		return -1, errors.Join(consts.InternalServerError, err)
	}

	_, err = tx.Exec(context.TODO(), setRoleByUserIdQuery, 3, userId)
	if err != nil {
		log.WithField(
			"origin.function", "CreateTeam",
		).Errorf("Ошибка при выдаче роли капитанаё: %s", err.Error())
		return -1, errors.Join(consts.InternalServerError, err)
	}

	err = tx.Commit(context.TODO())
	if err != nil {
		log.WithField(
			"origin.function", "CreateTeam",
		).Errorf("Ошибка при завершении транзакции: %s", err.Error())
		return -1, errors.Join(consts.InternalServerError, err)
	}

	return teamId, nil
}

func (s *TeamStorage) CheckUserHasTeamById(userId int) (exists bool, err error) {
	err = s.db.QueryRow(context.TODO(), checkUserHasTeamByIdQuery, userId).Scan(&exists)
	if err != nil {
		log.WithField(
			"origin.function", "CheckUserHasTeam",
		).Errorf("Ошибка при проверке существования команды у участника: %s", err.Error())
		return false, err
	}

	return exists, nil
}

func (s *TeamStorage) CheckUserIsExistByEmail(email string) (id int, err error) {
	err = s.db.QueryRow(context.TODO(), checkUserIsExistByEmailQuery, email).Scan(&id)
	if err != nil {
		log.WithField(
			"origin.function", "CheckUserIsExistByEmail",
		).Errorf("Ошибка при проверке существования участника по эмейлу: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (s *TeamStorage) CheckUserHasTeamByEmail(email string) (exists bool, err error) {
	err = s.db.QueryRow(context.TODO(), checkUserHasTeamByEmailQuery, email).Scan(&exists)
	if err != nil {
		log.WithField(
			"origin.function", "CheckUserHasTeam",
		).Errorf("Ошибка при проверке существования команды у участника: %s", err.Error())
		return false, err
	}

	return exists, nil
}

func (s *TeamStorage) SetTeamID(userId int, teamID int) error {
	_, err := s.db.Exec(context.TODO(), setTeamIdQuery, userId, teamID)
	if err != nil {
		log.WithField(
			"origin.function", "SetTeamID",
		).Errorf("Ошибка при назначении команды участнику: %s", err.Error())
		return err
	}

	return err
}

func (s *TeamStorage) UpdateTeamName(userId int, teamName string) error {
	_, err := s.db.Exec(context.TODO(), updateTeamNameQuery, teamName, userId)
	if err != nil {
		log.WithField(
			"origin.function", "UpdateTeamName",
		).Errorf("Ошибка при смене названия команды: %s", err.Error())
		return err
	}

	return nil
}

func (s *TeamStorage) CheckUserRoleById(userId int, role int) (exists bool, err error) {
	err = s.db.QueryRow(context.TODO(), checkUserRoleQuery, userId, role).Scan(&exists)
	if err != nil {
		log.WithField(
			"origin.function", "UserIsCapitanTeam",
		).Errorf("Ошибка при проверке роли в команде: %s", err.Error())
		return false, err
	}

	return exists, err
}

func (s *TeamStorage) DeleteMemberFromTeam(userId int, teamId int) error {
	_, err := s.db.Exec(context.TODO(), deleteMemberFromTeamQuery, userId)
	if err != nil {
		log.WithField(
			"origin.function", "DeleteMemberFromTeam",
		).Errorf("Ошибка при удалении участника из команды: %s", err.Error())
		return err
	}
	return nil
}

func (s *TeamStorage) AddMemberToTeam(userId int, teamId int) error {
	_, err := s.db.Exec(context.TODO(), addMemberToTeamQuery, teamId, userId)
	if err != nil {
		log.WithField(
			"origin.function", "DeleteMemberFromTeam",
		).Errorf("Ошибка при удалении участника из команды: %s", err.Error())
		return err
	}
	return nil
}

func (s *TeamStorage) GetTeamByUserId(userId int) (domain.Team, error) {
	var team domain.Team
	err := s.db.QueryRow(context.TODO(), getTeamByIdQuery, userId).Scan(
		&team.ID,
		&team.Name,
		&team.Points,
	)
	if err != nil {
		log.WithField(
			"origin.function", "GetTeamByUserId",
		).Errorf("Ошибка при получении данных о команде: %s", err.Error())
		return domain.Team{}, err
	}

	return team, nil
}

func (s *TeamStorage) GetMembersTeam(teamId int) ([]domain.Member, error) {
	var members []domain.Member
	rows, err := s.db.Query(context.TODO(), getMembersTeamQuery, teamId)
	if err != nil {
		log.WithField(
			"origin.function", "GetMembersTeam",
		).Errorf("Ошибка при получении данных об участниках команды: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		var member domain.Member
		err = rows.Scan(
			&member.ID,
			&member.RoleId,
			&member.Name,
			&member.Group,
			&member.Email,
		)
		if err != nil {
			log.WithField(
				"origin.function", "GetMembersTeam",
			).Errorf("Ошибка при получении данных об участнике команды: %s", err.Error())
			return nil, err
		}

		members = append(members, member)
	}

	return members, nil
}

func (s *TeamStorage) DeleteTeam(userId, teamId int) error {
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		log.WithField(
			"origin.function", "DeleteTeam",
		).Errorf("Ошибка при создании транзакции: %s", err.Error())
		return err
	}

	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.TODO(), setTeamIdQuery, nil, userId)
	if err != nil {
		log.WithField(
			"origin.function", "DeleteTeam",
		).Errorf("Ошибка при  удалении team_id: %s", err.Error())
		return err
	}

	_, err = tx.Exec(context.TODO(), deleteTeamQuery, teamId)
	if err != nil {
		log.WithField(
			"origin.function", "DeleteTeam",
		).Errorf("Ошибка при удалении команды: %s", err.Error())
		return err
	}

	_, err = tx.Exec(context.TODO(), setRoleByUserIdQuery, 1, userId)
	if err != nil {
		log.WithField(
			"origin.function", "DeleteTeam",
		).Errorf("Ошибка при смене роли участника команды: %s", err.Error())
		return err
	}

	err = tx.Commit(context.TODO())
	if err != nil {
		log.WithField(
			"origin.function", "DeleteTeam",
		).Errorf("Ошибка при завершении транзакции: %s", err.Error())
		return err
	}

	return nil
}

func (s *TeamStorage) GetTeamByFilter(count int) ([]domain.Team, error) {
	var teams []domain.Team
	rowsTeams, err := s.db.Query(context.TODO(), getTeamByCountUsersQuery, count)
	if err != nil {
		log.WithField(
			"origin.function", "GetTeamByFilter",
		).Errorf("Ошибка при получении данных о команде: %s", err.Error())
		return nil, err
	}

	for rowsTeams.Next() {
		var team domain.Team
		err = rowsTeams.Scan(&team.ID, &team.Name, &team.Points, &team)
		if err != nil {
			log.WithField(
				"origin.function", "GetTeamByFilter",
			).Errorf("Ошибка при получении данных о команде: %s", err.Error())
			return nil, err
		}

		teams = append(teams, team)
	}

	return teams, nil
}

func (s *TeamStorage) GetAllTeams() ([]domain.Team, error) {
	var teams []domain.Team
	rowsTeams, err := s.db.Query(context.TODO(), getAllTeamQuery)
	if err != nil {
		log.WithField(
			"origin.function", "GetTeamByFilter",
		).Errorf("Ошибка при получении данных о команде: %s", err.Error())
		return nil, err
	}

	for rowsTeams.Next() {
		var team domain.Team
		err = rowsTeams.Scan(&team.ID, &team.Name, &team.Points)
		if err != nil {
			log.WithField(
				"origin.function", "GetTeamByFilter",
			).Errorf("Ошибка при получении данных о команде: %s", err.Error())
			return nil, err
		}

		teams = append(teams, team)
	}

	return teams, nil
}

func (s *TeamStorage) GetTotalPointsByTeamId(teamId int) (totalPoints int, err error) {
	err = s.db.QueryRow(context.TODO(), getTotalPointsByTeamIdQuery, teamId).Scan(&totalPoints)
	if err != nil {
		log.WithField(
			"origin.function", "GetTotalPointsByTeamId",
		).Errorf("Ошибка при получении баллов команды: %s", err.Error())
		return 0, err
	}
	return totalPoints, nil
}

func (s *TeamStorage) UpdateSpendPointsByTeamId(teamId, deltaPoints int) error {
	tx, err := s.db.Begin(context.TODO())
	if err != nil {
		log.WithField(
			"origin.function", "GetTotalPointsByTeamId",
		).Errorf("Ошибка создания транзакцииё: %s", err.Error())
		return err
	}
	defer tx.Rollback(context.TODO())

	_, err = tx.Exec(context.TODO(), setDeltaPointsByTeamIdQuery, deltaPoints, teamId)
	if err != nil {
		log.WithField(
			"origin.function", "UpdateSpendPointsByTeamId",
		).Errorf("Ошибка при списании баллов у команды: %s", err.Error())
		return err
	}

	_, err = tx.Exec(context.TODO(), setVideoFlagInTeam, teamId)
	if err != nil {
		log.WithField(
			"origin.function", "GetTotalPointsByTeamId",
		).Errorf("Ошибка проставленияя флага: %s", err.Error())
		return err
	}

	err = tx.Commit(context.TODO())
	if err != nil {
		log.WithField(
			"origin.function", "GetTotalPointsByTeamId",
		).Errorf("Ошибка закрытия транзакции: %s", err.Error())
		return err
	}

	return nil
}

func (s *TeamStorage) UpdateGiverPointsByTeamId(teamId, deltaPoints int) error {
	_, err := s.db.Exec(context.TODO(), setDeltaPointsByTeamIdQuery, deltaPoints, teamId)
	if err != nil {
		log.WithField(
			"origin.function", "UpdateGiverPointsByTeamId",
		).Errorf("Ошибка при добавлении баллов: %s", err.Error())
	}

	return err
}

func (s *TeamStorage) GetCountUserInTeam(teamId int) (count int, err error) {
	err = s.db.QueryRow(context.TODO(), getCountUserInTeamQuery, teamId).Scan(&count)
	if err != nil {
		log.WithField(
			"origin.function", "GetCountUserInTeam",
		).Errorf("Ошибка при подсчете количества участников команды: %s", err.Error())
		return 0, err
	}

	return count, nil
}
