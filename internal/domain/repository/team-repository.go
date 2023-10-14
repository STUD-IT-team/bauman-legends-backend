package repository

import "github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"

type TeamStorage interface {
	CreateTeam(teamName string) (TeamID string, err error)
	CheckTeam(teamName string) (exists bool, err error)

	UpdateTeam(teamName string) error

	GetTeam(teamID string) (response.GetTeam, error)

	DeleteTeam(TeamID string) error

	InviteToTeam(UserID string, TeamID string) error

	DeleteFromTeam(UserID string, TeamID string) error

	UpdateMember(UserID string, RoleID int) error
}
