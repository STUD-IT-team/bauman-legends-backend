package app

import (
	"context"
	"errors"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
	log "github.com/sirupsen/logrus"
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

		err = t.storage.SetTeamID(res.UserID, teamID)
		if err != nil {
			return response.RegisterTeam{}, err
		}
		return response.RegisterTeam{
			TeamID: teamID,
		}, nil
	}

}

func (t *TeamService) UpdateTeam(req *request.ChangeTeam) (response.UpdateTeam, error) {
	//test valid data in req

	res, err := t.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return response.UpdateTeam{}, err
	}
	if !res.Valid {
		return response.UpdateTeam{}, errors.New("valid check error")
	}

	profile, err := t.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.Session})
	if err != nil {
		return response.UpdateTeam{}, err
	}
	err = t.storage.UpdateTeam(profile.TeamID, req.TeamName)
	if err != nil {
		return response.UpdateTeam{}, err
	}
	return response.UpdateTeam{TeamName: req.TeamName}, nil
}

func (t *TeamService) GetTeam(req *request.GetTeam) (*response.GetTeam, error) {
	res, err := t.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return nil, err
	}

	if !res.Valid {
		return nil, errors.New("valid check error")
	}

	profile, err := t.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.Session})
	if err != nil {
		return nil, err
	}

	log.Info(profile.TeamID)
	team, err := t.storage.GetTeam(profile.TeamID)
	if err != nil {
		return nil, err
	}
	team.Points, err = t.storage.GetTeamPoints(profile.TeamID)
	if err != nil {
		return nil, err
	}
	log.Infof("участники команды повыше:%+v", team)
	return mapper.MakeHttpResponseGetTeam(&team), nil
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
	if err != nil {
		return err
	}

	if err = t.storage.DeleteTeam(profile.TeamID); err != nil {
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
	exist, err := t.storage.CheckUserExist(req.UserID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("user does not exist in db")
	}
	profile, err := t.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.Session})
	exists, err := t.storage.CheckMembership(req.UserID, profile.TeamID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("userId already exists on this team")
	}

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
