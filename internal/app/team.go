package app

import (
	"context"
	"errors"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
	"google.golang.org/grpc"
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

func (t *TeamService) RegisterTeam(req *request.RegisterTeam) (response.RegisterTeam, error) {
	//test valid data in req

	res, err := t.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return response.RegisterTeam{}, err
	}
	if !res.Valid {
		return response.RegisterTeam{}, errors.New("valid check error")
	}

	exist, err := t.storage.CheckTeam(req.TeamName)

	if err != nil {
		return response.RegisterTeam{}, err
	}

	if exist {
		return response.RegisterTeam{}, errors.New("Team already exist")
	} else {
		teamID, err := t.storage.CreateTeam(req.TeamName)
		if err != nil {
			return response.RegisterTeam{}, err
		}
		return response.RegisterTeam{
			TeamID: teamID,
		}, nil
	}

}

func (t *TeamService) UpdateTeam(req *request.ChangeTeam) (response.RegisterTeam, error) {
	//test valid data in req

	res, err := t.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return response.RegisterTeam{}, err
	}
	if !res.Valid {
		return response.RegisterTeam{}, errors.New("valid check error")
	}

	exist, err := t.storage.CheckTeam(req.TeamName)

	if err != nil {
		return response.RegisterTeam{}, err
	}

	if exist {
		return response.RegisterTeam{TeamID: req.TeamName}, nil
	} else {
		return response.RegisterTeam{}, errors.New("team does not exist")
	}
}

func (t *TeamService) GetTeam(req *request.GetTeam) (response.GetTeam, error) {
	res, err := t.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return response.GetTeam{}, err
	}

	if !res.Valid {
		return response.GetTeam{}, errors.New("valid check error")
	}

	profile, err := t.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.Session})

	team, err := t.storage.GetTeam(profile.TeamID)

	if err != nil {
		return response.GetTeam{}, err
	}
	return team, nil
}

func (t *TeamService) DeleteTeam(req *request.DeleteTeam) error {
	res, err := t.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return err
	}

	if !res.Valid {
		return errors.New("valid check error")
	}

	profile, err := t.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.Session})

	err = t.storage.DeleteTeam(profile.TeamID)

	if err != nil {
		return err
	}
	return nil
}

func (t *TeamService) InviteToTeam(req *request.InviteToTeam) error {
	res, err := t.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return err
	}

	if !res.Valid {
		return errors.New("valid check error")
	}

	profile, err := t.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.Session})
	err = t.storage.InviteToTeam(req.UserID, profile.TeamID)

	if err != nil {
		return err
	}
	return nil
}

func (t *TeamService) DeleteFromTeam(req *request.DeleteFromTeam) error {
	res, err := t.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return err
	}

	if !res.Valid {
		return errors.New("valid check error")
	}

	profile, err := t.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.Session})
	err = t.storage.DeleteFromTeam(req.UserID, profile.TeamID)
	if err != nil {
		return err
	}
	return nil
}

func (t *TeamService) UpdateMember(req *request.UpdateMember) error {
	res, err := t.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return err
	}

	if !res.Valid {
		return errors.New("valid check error")
	}

	err = t.storage.UpdateMember(req.UserID, req.RoleID)
	if err != nil {
		return err
	}
	return nil
}
