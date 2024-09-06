package app

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"strconv"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/storage"
)

type UserService struct {
	storage storage.Storage
	auth    grpc2.AuthClient
}

func NewUserService(conn grpc.ClientConnInterface, s storage.Storage) *UserService {
	return &UserService{
		storage: s,
		auth:    grpc2.NewAuthClient(conn),
	}
}

func (s *UserService) GetUsersByFilter(session request.Session, req request.GetUsersByFilter) (response.GetUsersByFilter, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: session.Value})
	if err != nil {
		return response.GetUsersByFilter{}, err
	}

	if !res.Valid {
		return response.GetUsersByFilter{}, errors.New("invalid token")
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return response.GetUsersByFilter{}, err
	}

	isAdmin, err := s.storage.CheckUserRoleById(userId, consts.AdminRole)
	if err != nil {
		return response.GetUsersByFilter{}, err
	}

	if !isAdmin {
		return response.GetUsersByFilter{}, consts.ForbiddenError
	}

	var users []domain.Member

	if !req.WithTeam {
		users, err = s.storage.GetUserWithoutTeam()
		if err != nil {
			return response.GetUsersByFilter{}, err
		}
	} else {
		if req.CountInTeam == 0 {
			users, err = s.storage.GetAllUsers()
			if err != nil {
				return response.GetUsersByFilter{}, err
			}
		} else {
			users, err = s.storage.GetUserWithCountTeam(req.CountInTeam)
			if err != nil {
				return response.GetUsersByFilter{}, err
			}
		}
	}

	return *mapper.MakeGetUsersByFilter(users), nil
}

func (s *UserService) GetUserById(session request.Session, req request.GetUserById) (response.GetUserById, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: session.Value})
	if err != nil {
		return response.GetUserById{}, err
	}

	if !res.Valid {
		return response.GetUserById{}, errors.New("invalid token")
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return response.GetUserById{}, err
	}

	isAdmin, err := s.storage.CheckUserRoleById(userId, consts.AdminRole)
	if err != nil {
		return response.GetUserById{}, err
	}

	if !isAdmin {
		return response.GetUserById{}, consts.ForbiddenError
	}

	user, err := s.storage.GetUserById(req.Id)
	if err != nil {
		return response.GetUserById{}, err
	}

	return *mapper.MakeGetUserById(user), err
}
