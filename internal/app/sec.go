package app

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"strconv"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/storage"
)

type SECService struct {
	storage storage.Storage
	auth    grpc2.AuthClient
}

func NewSECService(conn grpc.ClientConnInterface, storage storage.Storage) *SECService {
	return &SECService{
		storage: storage,
		auth:    grpc2.NewAuthClient(conn),
	}
}

func (s *SECService) GetSECByFilter(filter request.GetSecByFilter, ses request.Session) (response.GetSecByFilter, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return response.GetSecByFilter{}, err
	}

	if !res.Valid {
		return response.GetSecByFilter{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	secs, err := s.storage.GetSECByFilter()
	if err != nil {
		return response.GetSecByFilter{}, err
	}

	return *mapper.MakeGetSECByFilter(secs), nil
}

func (s *SECService) GetSECById(filter request.GetSecByIdFilter, ses request.Session) (response.GetSecById, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return response.GetSecById{}, err
	}

	if !res.Valid {
		return response.GetSecById{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	secs, err := s.storage.GetSecByID(filter.Id)
	if err != nil {
		return response.GetSecById{}, err
	}

	return *mapper.MakeGetSECById(secs), nil
}

func (s *SECService) GetSECByTeamId(filter request.GetSecByTeamId, ses request.Session) (response.GetSecByTeamId, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return response.GetSecByTeamId{}, err
	}

	if !res.Valid {
		return response.GetSecByTeamId{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: ses.Value})
	if err != nil {
		return response.GetSecByTeamId{}, err
	}

	teamId, err := strconv.Atoi(profile.TeamID)
	if err != nil {
		return response.GetSecByTeamId{}, err
	}

	secs, err := s.storage.GetSecByTeamId(teamId)
	if err != nil {
		return response.GetSecByTeamId{}, err
	}

	return *mapper.MakeGetSECByTeamId(secs), nil
}

func (s *SECService) GetSECAdminByFilter(filter request.GetSecAdminByFilter, ses request.Session) (response.GetSecAdminByFilter, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return response.GetSecAdminByFilter{}, err
	}

	if !res.Valid {
		return response.GetSecAdminByFilter{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return response.GetSecAdminByFilter{}, err
	}

	isAdmin, err := s.storage.CheckUserRoleById(userId, consts.AdminRole)
	if err != nil {
		return response.GetSecAdminByFilter{}, err
	}

	if !isAdmin {
		return response.GetSecAdminByFilter{}, errors.Join(consts.ForbiddenError, errors.New("invalid user"))
	}

	secs, err := s.storage.GetSECAdmin()
	if err != nil {
		return response.GetSecAdminByFilter{}, err
	}

	return *mapper.MakeGetSECAdminByFilter(secs), nil
}

func (s *SECService) GetSECAdminById(filter request.GetSecAdminByIdFilter, ses request.Session) (response.GetSecAdminById, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return response.GetSecAdminById{}, err
	}

	if !res.Valid {
		return response.GetSecAdminById{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return response.GetSecAdminById{}, err
	}

	isAdmin, err := s.storage.CheckUserRoleById(userId, consts.AdminRole)
	if err != nil {
		return response.GetSecAdminById{}, err
	}

	if !isAdmin {
		return response.GetSecAdminById{}, errors.Join(consts.ForbiddenError, errors.New("invalid user"))
	}

	secs, err := s.storage.GetSECAdminById(filter.Id)
	if err != nil {
		return response.GetSecAdminById{}, err
	}

	return *mapper.MakeGetSECAdminById(secs), nil
}

func (s *SECService) DeleteRegisterOnSEC(filter request.DeleteRegisterOnSecFilter, ses request.Session) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return err
	}

	if !res.Valid {
		return errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return err
	}

	isCaptain, err := s.storage.CheckUserRoleById(userId, consts.CaptainRole)
	if err != nil {
		return err
	}

	if !isCaptain {
		return errors.Join(consts.UnAuthorizedError, errors.New("invalid user"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: ses.Value})
	if err != nil {
		return err
	}

	teamId, err := strconv.Atoi(profile.TeamID)
	if err != nil {
		return err
	}

	exist, err := s.storage.CheckRegisterOnMasterClass(filter.MasterClassId, teamId)
	if err != nil {
		return err
	}

	if !exist {
		return errors.Join(consts.ForbiddenError, errors.New("invalid user"))
	}

	err = s.storage.DeleteRegisterOnSEC(filter.MasterClassId, teamId)
	if err != nil {
		return err
	}

	return nil
}

func (s *SECService) CreateRegisterOnSEC(filter request.CreateRegisterOnSecFilter, ses request.Session) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return err
	}

	if !res.Valid {
		return errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return err
	}

	isCaptain, err := s.storage.CheckUserRoleById(userId, consts.CaptainRole)
	if err != nil {
		return err
	}

	if !isCaptain {
		return errors.Join(consts.ForbiddenError, errors.New("invalid user"))
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: ses.Value})
	if err != nil {
		return err
	}

	teamId, err := strconv.Atoi(profile.TeamID)
	if err != nil {
		return err
	}

	exist, err := s.storage.CheckRegisterOnSec(filter.MasterClassId, teamId)
	if err != nil {
		return err
	}

	if exist {
		return consts.ConflictError
	}

	exist, err = s.storage.CheckIntersectionTimeInterval(filter.MasterClassId, teamId)
	if err != nil {
		return err
	}

	if exist {
		return consts.ConflictError
	}

	exist, err = s.storage.CheckMasterClassIsExist(filter.MasterClassId)
	if err != nil {
		return err
	}

	if !exist {
		return consts.NotFoundError
	}

	exist, err = s.storage.CheckMasterClassTime(filter.MasterClassId)
	if err != nil {
		return err
	}

	if exist {
		return consts.LockedError
	}

	count, err := s.storage.CheckMasterClassBusyPlaceById(filter.MasterClassId, teamId)
	if err != nil {
		return err
	}

	if count < 0 {
		return consts.ConflictError
	}

	err = s.storage.CreateRegisterOnSEC(filter.MasterClassId, teamId)
	if err != nil {
		return err
	}

	return nil
}

func (s *SECService) GetMasterClassById(filter request.GetMasterClassByID, ses request.Session) (response.GetMasterClassByID, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: ses.Value})
	if err != nil {
		return response.GetMasterClassByID{}, err
	}

	if !res.Valid {
		return response.GetMasterClassByID{}, errors.Join(consts.UnAuthorizedError, errors.New("valid check error"))
	}

	secs, err := s.storage.GetMasterClassByID(filter.MasterClassId)
	if err != nil {
		return response.GetMasterClassByID{}, err
	}

	return *mapper.MakeGetMasterClassById(secs), nil
}
