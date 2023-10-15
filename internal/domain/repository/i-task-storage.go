package repository

import "github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"

type ITaskStorage interface {
	GetTaskTypes() (*response.TaskTypes, error)
}
