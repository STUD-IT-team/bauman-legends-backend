package repository

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
