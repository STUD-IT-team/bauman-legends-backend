package storage

import "github.com/STUD-IT-team/bauman-legends-backend/internal/domain"

type SECStorage interface {
	GetSECByFilter() ([]domain.Sec, error)
	GetSecByID(id int) ([]domain.Sec, error)
	GetSecByTeamId(teamId int) ([]domain.Sec, error)
	CreateRegisterOnSEC(secId int, time string, teamId int) error
	DeleteRegisterOnSEC(secId int, time string, teamId int) error
	GetSECAdmin() ([]domain.Sec, error)
	GetSECAdminById(secId int) ([]domain.Sec, error)
	CheckRegisterOnMasterClass(secId int, time string, teamId int) (bool, error)
	CheckRegisterOnSec(secId, teamId int) (bool, error)
	CheckIntersectionTimeInterval(secId int, time string, teamId int) (bool, error)
	CheckMasterClassIsExist(secId int, time string) (bool, error)
	CheckMasterClassBusyPlaceById(secId int, time string, teamId int) (int, error)
}

func (s *storage) GetSECByFilter() ([]domain.Sec, error) {
	return s.SEC.GetSECByFilter()
}

func (s *storage) GetSecByID(secId int) ([]domain.Sec, error) {
	return s.SEC.GetSecByID(secId)
}

func (s *storage) GetSecByTeamId(teamId int) ([]domain.Sec, error) {
	return s.SEC.GetSecByTeamId(teamId)
}

func (s *storage) GetSECAdmin() ([]domain.Sec, error) {
	return s.SEC.GetSECAdmin()
}

func (s *storage) GetSECAdminById(secId int) ([]domain.Sec, error) {
	return s.SEC.GetSECAdminById(secId)
}

func (s *storage) CreateRegisterOnSEC(secId int, time string, userId int) error {
	return s.SEC.CreateRegisterOnSEC(secId, time, userId)
}

func (s *storage) DeleteRegisterOnSEC(secId int, time string, teamId int) error {
	return s.SEC.DeleteRegisterOnSEC(secId, time, teamId)
}

func (s *storage) CheckRegisterOnMasterClass(secId int, time string, teamId int) (bool, error) {
	return s.SEC.CheckRegisterOnMasterClass(secId, time, teamId)
}

func (s *storage) CheckRegisterOnSec(secId, teamId int) (bool, error) {
	return s.SEC.CheckRegisterOnSec(secId, teamId)
}

func (s *storage) CheckIntersectionTimeInterval(secId int, time string, teamId int) (bool, error) {
	return s.SEC.CheckIntersectionTimeInterval(secId, time, teamId)
}

func (s *storage) CheckMasterClassIsExist(secId int, time string) (bool, error) {
	return s.SEC.CheckMasterClassIsExist(secId, time)
}

func (s *storage) CheckMasterClassBusyPlaceById(secId int, time string, teamId int) (int, error) {
	return s.SEC.CheckMasterClassBusyPlaceById(secId, time, teamId)
}
