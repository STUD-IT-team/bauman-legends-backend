package repository

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
)

type TeamStorage interface {
	CreateTeam(teamName string) (TeamID string, err error)
	CheckTeam(teamName string) (exists bool, err error)

	UpdateTeam(teamID string, teamName string) error

	GetTeam(teamID string) (domain.Team, error)

	DeleteTeam(TeamID string) error

	InviteToTeam(UserID string, TeamID string) error

	DeleteFromTeam(UserID string, TeamID string) error

	UpdateMember(UserID string, RoleID int) error

	SetTeamID(userID string, teamID string) error

	CheckMembership(userId string, teamID string) (bool, error)

	CheckUserExist(userID string) (bool, error)

}
