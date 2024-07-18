package app

import (
	"google.golang.org/grpc"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
)

type TaskService struct {
	storage repository.TaskStorage
	auth    grpc2.AuthClient
}

func NewTaskService(conn grpc.ClientConnInterface, r repository.TaskStorage) *TaskService {
	return &TaskService{
		storage: r,
		auth:    grpc2.NewAuthClient(conn),
	}
}
