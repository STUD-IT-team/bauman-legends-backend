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
	CheckRegisterOnSEC(secId int, time string, teamId int) (bool, error)
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
	// TODO implement me
	panic("implement me")
}

func (s *storage) DeleteRegisterOnSEC(secId int, time string, userId int) error {
	// TODO implement me
	panic("implement me")
}

func (s *storage) CheckRegisterOnSEC(secId int, time string, teamId int) (bool, error) {
	return s.SEC.CheckRegisterOnSEC(secId, time, teamId)
}
