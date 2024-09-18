package storage

import "github.com/STUD-IT-team/bauman-legends-backend/internal/domain"

type SECStorage interface {
	GetSECByFilter() ([]domain.Sec, error)
	GetSecByID(id int) ([]domain.Sec, error)
	GetSecByTeamId(teamId int) ([]domain.Sec, error)
	CreateRegisterOnSEC(masterClassId, teamId int) error
	DeleteRegisterOnSEC(masterClassId, teamId int) error
	GetSECAdmin() ([]domain.Sec, error)
	GetSECAdminById(secId int) ([]domain.Sec, error)
	CheckRegisterOnMasterClass(masterClassId, teamId int) (bool, error)
	CheckRegisterOnSec(secId, teamId int) (bool, error)
	CheckIntersectionTimeInterval(masterClassId, teamId int) (bool, error)
	CheckMasterClassIsExist(masterClassId int) (bool, error)
	CheckMasterClassBusyPlaceById(masterClassId, teamId int) (int, error)
	CheckMasterClassTime(masterClass int) (bool, error)
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

func (s *storage) CreateRegisterOnSEC(masterClassId, userId int) error {
	return s.SEC.CreateRegisterOnSEC(masterClassId, userId)
}

func (s *storage) DeleteRegisterOnSEC(masterClassId, teamId int) error {
	return s.SEC.DeleteRegisterOnSEC(masterClassId, teamId)
}

func (s *storage) CheckRegisterOnMasterClass(masterClassId, teamId int) (bool, error) {
	return s.SEC.CheckRegisterOnMasterClass(masterClassId, teamId)
}

func (s *storage) CheckRegisterOnSec(secId, teamId int) (bool, error) {
	return s.SEC.CheckRegisterOnSec(secId, teamId)
}

func (s *storage) CheckIntersectionTimeInterval(masterClassId, teamId int) (bool, error) {
	return s.SEC.CheckIntersectionTimeInterval(masterClassId, teamId)
}

func (s *storage) CheckMasterClassIsExist(masterClassId int) (bool, error) {
	return s.SEC.CheckMasterClassIsExist(masterClassId)
}

func (s *storage) CheckMasterClassBusyPlaceById(masterClassId, teamId int) (int, error) {
	return s.SEC.CheckMasterClassBusyPlaceById(masterClassId, teamId)
}

func (s *storage) CheckMasterClassTime(masterClass int) (bool, error) {
	return s.SEC.CheckMasterClassTime(masterClass)
}
