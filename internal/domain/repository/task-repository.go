package repository

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"time"
)

type TaskStorage interface {
	GetTaskTypes(teamID string) (domain.TaskTypes, error)
	GetTaskAmount(taskTypeID int) (int, error)
	GetBusyNocPlaceses() (int, error)
	GetAvailableTaskID(teamID string, taskTypeID int) ([]string, error)
	GetTeamTaskAmount(taskTypeID int) (int, error)
	SetTaskToTeam(taskID string, teamID string) error
	CheckActiveTaskExist(teamID string) (bool, error)
	GetActiveTaskID(teamID string) (string, error)
	GetTask(taskID string) (domain.Task, error)
	GetTaskStartedTime(taskID string, TeamID string) (time.Time, error)
	GetTaskTypeName(taskID string) (string, error)
	SetActiveTaskExpired(taskID string, TeamID string) error
	SetAnswerText(text string, teamID string, taskID string) error
	SetAnswerImageBase64(url string, teamID string, taskID string) error
	GetAnswers(teamID string) ([]domain.Answer, error)
}
