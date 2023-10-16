package app

import (
	"context"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
)

type Task struct {
	Repository repository.ITaskStorage
	grpc2.UnimplementedTaskServer
}

func NewTask(r repository.ITaskStorage) *Task {
	return &Task{
		Repository: r,
	}
}

func (t *Task) GetTaskTypes(_ context.Context, _ *grpc2.GetTaskTypesRequest) (*grpc2.TaskTypesResponse, error) {
	t.Repository.GetTaskTypes()
	return nil, nil
}
