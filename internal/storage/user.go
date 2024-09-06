package storage

import "github.com/STUD-IT-team/bauman-legends-backend/internal/domain"

type UserStorage interface {
	GetAllUsers() ([]domain.Member, error)
	GetUserWithoutTeam() ([]domain.Member, error)
	GetUserWithCountTeam(count int) ([]domain.Member, error)
	GetUserById(id int) (domain.Member, error)
}

func (s *storage) GetAllUsers() ([]domain.Member, error) {
	return s.User.GetAllUsers()
}

func (s *storage) GetUserWithoutTeam() ([]domain.Member, error) {
	return s.User.GetUserWithoutTeam()
}

func (s *storage) GetUserWithCountTeam(count int) ([]domain.Member, error) {
	return s.User.GetUserWithCountTeam(count)
}

func (s *storage) GetUserById(id int) (domain.Member, error) {
	return s.User.GetUserById(id)
}
