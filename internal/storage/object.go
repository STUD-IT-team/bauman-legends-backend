package storage

import "github.com/STUD-IT-team/bauman-legends-backend/internal/domain"

type ObjectStorage interface {
	GetObject(object domain.Object) (domain.Object, error)
	PutObject(object domain.Object) (string, error)
	DeleteObject(object domain.Object) error
}

func (s *storage) GetObject(object domain.Object) (domain.Object, error) {
	return s.Object.GetObject(object)
}

func (s *storage) PutObject(object domain.Object) (string, error) {
	return s.Object.PutObject(object)
}

func (s *storage) DeleteObject(object domain.Object) error {
	return s.Object.DeleteObject(object)
}
