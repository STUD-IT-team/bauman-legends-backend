package app

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"strconv"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
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

func (s *TeamService) CreateTeam(req request.CreateTeam) (response.CreateTeam, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return response.CreateTeam{}, err
	}

	if !res.Valid {
		return response.CreateTeam{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	exist, err := s.storage.CheckTeam(req.TeamName)
	if err != nil {
		return response.CreateTeam{}, err
	}

	if exist {
		return response.CreateTeam{}, errors.Join(consts.ConflictError, errors.New("team already exist"))
	}

	exist, err = s.storage.CheckUserHasTeamById(res.UserID)
	if err != nil {
		return response.CreateTeam{}, err
	}

	if exist {
		return response.CreateTeam{}, errors.Join(consts.ConflictError, errors.New("user already has team"))
	}

	teamID, err := s.storage.CreateTeam(req.TeamName, res.UserID)
	if err != nil {
		return response.CreateTeam{}, err
	}

	return response.CreateTeam{
		TeamId: teamID,
	}, nil
}

func (s *TeamService) DeleteTeam(req *request.DeleteTeam) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return err
	}
	if !res.Valid {
		return errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.Session})
	if err != nil {
		return err
	}

	count, err := s.storage.GetCountUserInTeam(profile.TeamID)
	if err != nil {
		return err
	}

	if count > 1 {
		return errors.Join(consts.ConflictError, errors.New("team is not empty"))
	}

	err = s.storage.DeleteTeam(profile.Id, profile.TeamID)
	if err != nil {
		return err
	}

	return err
}

func (s *TeamService) UpdateTeam(req *request.UpdateTeam) (response.UpdateTeam, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return response.UpdateTeam{}, err
	}
	if !res.Valid {
		return response.UpdateTeam{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.Session})
	if err != nil {
		return response.UpdateTeam{}, err
	}

	isCapitan, err := s.storage.CheckUserRoleById(profile.Id, consts.CaptainRole)
	if err != nil {
		return response.UpdateTeam{}, err
	}

	if !isCapitan {
		return response.UpdateTeam{}, errors.Join(consts.ForbiddenError, errors.New("не является капитаном нет прав на изменение названия"))
	}

	err = s.storage.UpdateTeamName(profile.Id, req.NewTeamName)
	if err != nil {
		return response.UpdateTeam{}, err
	}

	return response.UpdateTeam{TeamName: req.NewTeamName}, nil
}

func (s *TeamService) GetTeam(req *request.GetTeam) (response.GetTeam, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return response.GetTeam{}, err
	}
	if !res.Valid {
		return response.GetTeam{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.Session})
	if err != nil {
		return response.GetTeam{}, err
	}

	team, err := s.storage.GetTeamByUserId(profile.Id)
	if err != nil {
		return response.GetTeam{}, err
	}

	members, err := s.storage.GetMembersTeam(profile.TeamID)
	if err != nil {
		return response.GetTeam{}, err
	}

	for _, member := range members {
		if member.RoleId == consts.CaptainRole {
			team.Captain = member
		} else {
			team.Members = append(team.Members, member)
		}
	}

	return *mapper.MakeGetTeamResponse(team), nil
}

func (s *TeamService) GetTeamById(req *request.GetTeamById) (response.GetTeamByID, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return response.GetTeamByID{}, err
	}
	if !res.Valid {
		return response.GetTeamByID{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	team, err := s.storage.GetTeamByUserId(req.TeamId)
	if err != nil {
		return response.GetTeamByID{}, err
	}

	members, err := s.storage.GetMembersTeam(req.TeamId)
	if err != nil {
		return response.GetTeamByID{}, err
	}

	for _, member := range members {
		if member.RoleId == consts.CaptainRole {
			team.Captain = member
		} else {
			team.Members = append(team.Members, member)
		}
	}

	return *mapper.MakeGetTeamByIdResponse(team), nil
}

func (s *TeamService) AddMemberToTeam(req *request.AddMemberToTeam) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return err
	}
	if !res.Valid {
		return errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.Session})
	if err != nil {
		return err
	}

	count, err := s.storage.GetCountUserInTeam(profile.TeamID)
	if err != nil {
		return err
	}

	if count >= 8 {
		return errors.Join(consts.ConflictError, errors.New("нет места в команде"))
	}

	userId, err := s.storage.CheckUserIsExistByEmail(req.UserEmail)
	if err != nil {
		return err
	}

	if userId == 0 {
		return errors.Join(consts.ConflictError, errors.New("участник с таким емейлом не существует"))
	}

	exist, err := s.storage.CheckUserHasTeamByEmail(req.UserEmail)
	if err != nil {
		return err
	}

	if exist {
		return errors.Join(consts.ConflictError, errors.New("участник уже состоит в команде"))
	}

	isCaptain, err := s.storage.CheckUserRoleById(profile.Id, consts.CaptainRole)
	if err != nil {
		return err
	}

	if !isCaptain {
		return errors.Join(consts.ForbiddenError, errors.New("не является капитаном"))
	}

	err = s.storage.AddMemberToTeam(strconv.Itoa(userId), profile.TeamID)
	if err != nil {
		return err
	}

	return nil
}

func (s *TeamService) DeleteMemberFromTeam(req *request.DeleteMemberFromTeam) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return err
	}
	if !res.Valid {
		return errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.Session})
	if err != nil {
		return err
	}

	exist, err := s.storage.CheckUserHasTeamById(req.UserID)
	if err != nil {
		return err
	}

	if !exist {
		return errors.Join(consts.ConflictError, errors.New("участник не состоит в команде"))
	}

	exist, err = s.storage.CheckUserRoleById(profile.Id, consts.CaptainRole)
	if err != nil {
		return err
	}

	if !exist {
		return errors.Join(consts.ForbiddenError, errors.New("не является капитаном нет прав на удаление"))
	}

	exist, err = s.storage.CheckUserRoleById(req.UserID, consts.CaptainRole)
	if err != nil {
		return err
	}

	if exist {
		return errors.Join(consts.ConflictError, errors.New("удаляемый участник является капитаном"))
	}

	err = s.storage.DeleteMemberFromTeam(req.UserID, profile.TeamID)
	if err != nil {
		return err
	}

	return nil
}

func (s *TeamService) UpdateSpendPoints(req *request.UpdateSpendPoints) (response.UpdateSpendPoints, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return response.UpdateSpendPoints{}, err
	}

	if !res.Valid {
		return response.UpdateSpendPoints{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	isSeller, err := s.storage.CheckUserRoleById(res.UserID, consts.SellerRole)
	if err != nil {
		return response.UpdateSpendPoints{}, err
	}

	if !isSeller {
		return response.UpdateSpendPoints{}, errors.Join(consts.ForbiddenError, errors.New("пользователь не является продавцом"))
	}

	totalPoints, err := s.storage.GetTotalPointsByTeamId(req.TeamId)
	if err != nil {
		return response.UpdateSpendPoints{}, err
	}

	if totalPoints < req.DeltaPoints {
		return response.UpdateSpendPoints{}, errors.Join(consts.LockedError, errors.New("недостаточно баллов на счету"))
	}

	err = s.storage.UpdateSpendPointsByTeamId(req.TeamId, -req.DeltaPoints)
	if err != nil {
		return response.UpdateSpendPoints{}, err
	}

	return response.UpdateSpendPoints{TotalPoints: totalPoints - req.DeltaPoints}, nil
}

func (s *TeamService) UpdateGiverPoints(req *request.UpdateGivePoints) (response.UpdateGivePoints, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return response.UpdateGivePoints{}, err
	}

	if !res.Valid {
		return response.UpdateGivePoints{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	isSeller, err := s.storage.CheckUserRoleById(res.UserID, consts.SellerRole)
	if err != nil {
		return response.UpdateGivePoints{}, err
	}

	if !isSeller {
		return response.UpdateGivePoints{}, errors.Join(consts.ForbiddenError, errors.New("пользователь не является продавцом"))
	}

	totalPoints, err := s.storage.GetTotalPointsByTeamId(req.TeamId)
	if err != nil {
		return response.UpdateGivePoints{}, err
	}

	err = s.storage.UpdateSpendPointsByTeamId(req.TeamId, req.DeltaPoints)
	if err != nil {
		return response.UpdateGivePoints{}, err
	}

	return response.UpdateGivePoints{TotalPoints: totalPoints + req.DeltaPoints}, nil
}

func (s *TeamService) GetTeamsByFilter(req *request.GetTeamsByFilter) (response.GetTeamsByFilter, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.Session})
	if err != nil {
		return response.GetTeamsByFilter{}, err
	}

	if !res.Valid {
		return response.GetTeamsByFilter{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	isAdmin, err := s.storage.CheckUserRoleById(res.UserID, consts.AdminRole)
	if err != nil {
		return response.GetTeamsByFilter{}, err
	}

	isSeller, err := s.storage.CheckUserRoleById(res.UserID, consts.SellerRole)
	if err != nil {
		return response.GetTeamsByFilter{}, err
	}

	if !isAdmin || !isSeller {
		return response.GetTeamsByFilter{}, errors.Join(consts.ForbiddenError, errors.New("не является админом или продавцом"))
	}

	var teams []domain.Team
	if req.MembersCount == 0 {
		teams, err = s.storage.GetAllTeams()
	} else {
		teams, err = s.storage.GetTeamByFilter(req.MembersCount)
	}
	if err != nil {
		return response.GetTeamsByFilter{}, err
	}

	for _, team := range teams {
		members, err := s.storage.GetMembersTeam(strconv.Itoa(team.ID))
		if err != nil {
			return response.GetTeamsByFilter{}, err
		}

		for _, member := range members {
			if member.RoleId == consts.CaptainRole {
				team.Captain = member
			} else {
				team.Members = append(team.Members, member)
			}
		}
	}

	return *mapper.MakeGetTeamsResponse(teams), nil
}
