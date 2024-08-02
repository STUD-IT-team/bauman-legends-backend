package repository

import "github.com/STUD-IT-team/bauman-legends-backend/internal/domain"

type TeamStorage interface {
	CheckTeam(teamName string) (exists bool, err error)
	CreateTeam(teamName string, userId string) (teamId int, err error)
	CheckUserHasTeamById(userId string) (exists bool, err error)
	CheckUserHasTeamByEmail(email string) (exists bool, err error)
	CheckUserIsExistByEmail(email string) (id int, err error)
	SetTeamID(userId string, teamID int) error
	UpdateTeamName(userId string, teamName string) error
	DeleteTeam(userId, teamId string) error
	AddMemberToTeam(userId string, teamId string) error
	DeleteMemberFromTeam(userId string, teamId string) error
	GetTeamByUserId(userId string) (domain.Team, error)
	GetMembersTeam(teamId string) ([]domain.Member, error)
	GetTeamByFilter(count int) ([]domain.Team, error)
	GetAllTeams() ([]domain.Team, error)
	GetTotalPointsByTeamId(teamId int) (totalPoints int, err error)
	UpdateSpendPointsByTeamId(teamId, deltaPoints int) error
	UpdateGiverPointsByTeamId(teamId, deltaPoints int) error
	GetCountUserInTeam(teamId string) (count int, err error)
	CheckUserRoleById(userId string, role int) (exists bool, err error)
}
