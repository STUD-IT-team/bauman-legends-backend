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

type TeamService struct {
	storage storage.Storage
	auth    grpc2.AuthClient
}

func NewTeamService(conn grpc.ClientConnInterface, s storage.Storage) *TeamService {
	return &TeamService{
		storage: s,
		auth:    grpc2.NewAuthClient(conn),
	}
}

func (s *TeamService) CreateTeam(req request.CreateTeam, ses request.Session) (response.CreateTeam, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
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

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return response.CreateTeam{}, err
	}

	exist, err = s.storage.CheckUserHasTeamById(userId)
	if err != nil {
		return response.CreateTeam{}, err
	}

	if exist {
		return response.CreateTeam{}, errors.Join(consts.ConflictError, errors.New("user already has team"))
	}

	teamID, err := s.storage.CreateTeam(req.TeamName, userId)
	if err != nil {
		return response.CreateTeam{}, err
	}

	return response.CreateTeam{
		TeamId: teamID,
	}, nil
}

func (s *TeamService) DeleteTeam(ses request.Session) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return err
	}
	if !res.Valid {
		return errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: ses.Value})
	if err != nil {
		return err
	}

	teamId, err := strconv.Atoi(profile.TeamID)
	if err != nil {
		return err
	}

	count, err := s.storage.GetCountUserInTeam(teamId)
	if err != nil {
		return err
	}

	if count > 1 {
		return errors.Join(consts.ConflictError, errors.New("team is not empty"))
	}

	id, err := strconv.Atoi(profile.Id)
	if err != nil {
		return err
	}

	err = s.storage.DeleteTeam(id, teamId)
	if err != nil {
		return err
	}

	return err
}

func (s *TeamService) UpdateTeam(req *request.UpdateTeam, ses request.Session) (response.UpdateTeam, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return response.UpdateTeam{}, err
	}
	if !res.Valid {
		return response.UpdateTeam{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: ses.Value})
	if err != nil {
		return response.UpdateTeam{}, err
	}

	id, err := strconv.Atoi(profile.Id)
	if err != nil {
		return response.UpdateTeam{}, err
	}
	isCapitan, err := s.storage.CheckUserRoleById(id, consts.CaptainRole)
	if err != nil {
		return response.UpdateTeam{}, err
	}

	if !isCapitan {
		return response.UpdateTeam{}, errors.Join(consts.ForbiddenError, errors.New("не является капитаном нет прав на изменение названия"))
	}

	err = s.storage.UpdateTeamName(id, req.NewTeamName)
	if err != nil {
		return response.UpdateTeam{}, err
	}

	return response.UpdateTeam{TeamName: req.NewTeamName}, nil
}

func (s *TeamService) GetTeam(ses request.Session) (response.GetTeam, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return response.GetTeam{}, err
	}
	if !res.Valid {
		return response.GetTeam{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: ses.Value})
	if err != nil {
		return response.GetTeam{}, err
	}

	id, err := strconv.Atoi(profile.Id)
	if err != nil {
		return response.GetTeam{}, err
	}
	team, err := s.storage.GetTeamByUserId(id)
	if err != nil {
		return response.GetTeam{}, err
	}

	teamId, err := strconv.Atoi(profile.TeamID)
	if err != nil {
		return response.GetTeam{}, err
	}

	members, err := s.storage.GetMembersTeam(teamId)
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

func (s *TeamService) GetTeamById(req *request.GetTeamById, ses request.Session) (response.GetTeamByID, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return response.GetTeamByID{}, err
	}
	if !res.Valid {
		return response.GetTeamByID{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	team, err := s.storage.GetTeamByUserId(req.Id)
	if err != nil {
		return response.GetTeamByID{}, err
	}

	members, err := s.storage.GetMembersTeam(req.Id)
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

func (s *TeamService) AddMemberToTeam(req *request.AddMemberToTeam, ses request.Session) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return err
	}
	if !res.Valid {
		return errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: ses.Value})
	if err != nil {
		return err
	}

	teamId, err := strconv.Atoi(profile.Id)
	if err != nil {
		return err
	}
	count, err := s.storage.GetCountUserInTeam(teamId)
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

	id, err := strconv.Atoi(profile.Id)
	if err != nil {
		return err
	}
	isCaptain, err := s.storage.CheckUserRoleById(id, consts.CaptainRole)
	if err != nil {
		return err
	}

	if !isCaptain {
		return errors.Join(consts.ForbiddenError, errors.New("не является капитаном"))
	}

	err = s.storage.AddMemberToTeam(userId, teamId)
	if err != nil {
		return err
	}

	return nil
}

func (s *TeamService) DeleteMemberFromTeam(req *request.DeleteMemberFromTeam, ses request.Session) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return err
	}
	if !res.Valid {
		return errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: ses.Value})
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

	id, err := strconv.Atoi(profile.Id)
	if err != nil {
		return err
	}
	exist, err = s.storage.CheckUserRoleById(id, consts.CaptainRole)
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

	teamId, err := strconv.Atoi(profile.TeamID)
	if err != nil {
		return err
	}
	err = s.storage.DeleteMemberFromTeam(req.UserID, teamId)
	if err != nil {
		return err
	}

	return nil
}

func (s *TeamService) UpdateSpendPoints(req *request.UpdateSpendPoints, ses request.Session) (response.UpdateSpendPoints, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return response.UpdateSpendPoints{}, err
	}

	if !res.Valid {
		return response.UpdateSpendPoints{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return response.UpdateSpendPoints{}, err
	}

	isSeller, err := s.storage.CheckUserRoleById(userId, consts.SellerRole)
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

func (s *TeamService) UpdateGiverPoints(req *request.UpdateGivePoints, ses request.Session) (response.UpdateGivePoints, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return response.UpdateGivePoints{}, err
	}

	if !res.Valid {
		return response.UpdateGivePoints{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return response.UpdateGivePoints{}, err
	}

	isSeller, err := s.storage.CheckUserRoleById(userId, consts.SellerRole)
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

func (s *TeamService) GetTeamsByFilter(req *request.GetTeamsByFilter, ses request.Session) (response.GetTeamsByFilter, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return response.GetTeamsByFilter{}, err
	}

	if !res.Valid {
		return response.GetTeamsByFilter{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return response.GetTeamsByFilter{}, err
	}

	isAdmin, err := s.storage.CheckUserRoleById(userId, consts.AdminRole)
	if err != nil {
		return response.GetTeamsByFilter{}, err
	}

	isSeller, err := s.storage.CheckUserRoleById(userId, consts.SellerRole)
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
		members, err := s.storage.GetMembersTeam(team.ID)
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
