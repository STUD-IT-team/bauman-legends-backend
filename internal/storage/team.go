package storage

import "github.com/STUD-IT-team/bauman-legends-backend/internal/domain"

type TeamStorage interface {
	CheckTeam(teamName string) (exists bool, err error)
	CreateTeam(teamName string, userId int) (teamId int, err error)
	CheckUserHasTeamById(userId int) (exists bool, err error)
	CheckUserHasTeamByEmail(email string) (exists bool, err error)
	CheckUserIsExistByEmail(email string) (id int, err error)
	SetTeamID(userId int, teamID int) error
	UpdateTeamName(userId int, teamName string) error
	DeleteTeam(userId, teamId int) error
	AddMemberToTeam(userId int, teamId int) error
	DeleteMemberFromTeam(userId int, teamId int) error
	GetTeamByUserId(userId int) (domain.Team, error)
	GetMembersTeam(teamId int) ([]domain.Member, error)
	GetTeamByFilter(count int) ([]domain.Team, error)
	GetAllTeams() ([]domain.Team, error)
	GetTotalPointsByTeamId(teamId int) (totalPoints int, err error)
	UpdateSpendPointsByTeamId(teamId, deltaPoints int) error
	UpdateGiverPointsByTeamId(teamId, deltaPoints int) error
	GetCountUserInTeam(teamId int) (count int, err error)
	CheckUserRoleById(userId int, role int) (exists bool, err error)
}

func (s *storage) CheckTeam(teamName string) (exists bool, err error) {
	return s.Team.CheckTeam(teamName)
}

func (s *storage) CreateTeam(teamName string, userId int) (teamId int, err error) {
	return s.Team.CreateTeam(teamName, userId)
}

func (s *storage) CheckUserHasTeamById(userId int) (exists bool, err error) {
	return s.Team.CheckUserHasTeamById(userId)
}

func (s *storage) CheckUserHasTeamByEmail(email string) (exists bool, err error) {
	return s.Team.CheckUserHasTeamByEmail(email)
}

func (s *storage) CheckUserIsExistByEmail(email string) (id int, err error) {
	return s.Team.CheckUserIsExistByEmail(email)
}

func (s *storage) SetTeamID(userId int, teamID int) error {
	return s.Team.SetTeamID(userId, teamID)
}

func (s *storage) UpdateTeamName(userId int, teamName string) error {
	return s.Team.UpdateTeamName(userId, teamName)
}

func (s *storage) DeleteTeam(userId, teamId int) error {
	return s.Team.DeleteTeam(userId, teamId)
}

func (s *storage) AddMemberToTeam(userId int, teamId int) error {
	return s.Team.AddMemberToTeam(userId, teamId)
}

func (s *storage) DeleteMemberFromTeam(userId int, teamId int) error {
	return s.Team.DeleteMemberFromTeam(userId, teamId)
}
func (s *storage) GetTeamByUserId(userId int) (domain.Team, error) {
	return s.Team.GetTeamByUserId(userId)
}

func (s *storage) GetMembersTeam(teamId int) ([]domain.Member, error) {
	return s.Team.GetMembersTeam(teamId)
}

func (s *storage) GetTeamByFilter(count int) ([]domain.Team, error) {
	return s.Team.GetTeamByFilter(count)
}

func (s *storage) GetAllTeams() ([]domain.Team, error) {
	return s.Team.GetAllTeams()
}

func (s *storage) GetTotalPointsByTeamId(teamId int) (totalPoints int, err error) {
	return s.Team.GetTotalPointsByTeamId(teamId)
}

func (s *storage) UpdateSpendPointsByTeamId(teamId, deltaPoints int) error {
	return s.Team.UpdateSpendPointsByTeamId(teamId, deltaPoints)
}

func (s *storage) UpdateGiverPointsByTeamId(teamId, deltaPoints int) error {
	return s.Team.UpdateGiverPointsByTeamId(teamId, deltaPoints)
}

func (s *storage) GetCountUserInTeam(teamId int) (count int, err error) {
	return s.Team.GetCountUserInTeam(teamId)
}

func (s *storage) CheckUserRoleById(userId int, role int) (exists bool, err error) {
	return s.Team.CheckUserRoleById(userId, role)
}
