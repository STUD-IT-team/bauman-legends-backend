package app

import (
	"google.golang.org/grpc"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
)

type TeamService struct {
	storage repository.TeamStorage
	auth    grpc2.AuthClient
}

func NewTeamService(conn grpc.ClientConnInterface, r repository.TeamStorage) *TeamService {
	return &TeamService{
		storage: r,
		auth:    grpc2.NewAuthClient(conn),
	}
}
