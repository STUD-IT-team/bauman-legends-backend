package storage

import "github.com/STUD-IT-team/bauman-legends-backend/internal/domain"

type TextTaskStorage interface {
	GetNewTextTask(teamId int) (domain.TextTask, error)
	GetLastTextTask(teamId int) (domain.TextTask, error)
	GetStatusLastTextTask(teamId int) (status bool, err error)
	CreateAnswerOnTextTask(task domain.TextTask) error
	UpdateAnswerOnTextTask(task domain.TextTask) error
}

func (s *storage) GetNewTextTask(teamId int) (domain.TextTask, error) {
	return s.TextTask.GetNewTextTask(teamId)
}

func (s *storage) GetLastTextTask(teamId int) (domain.TextTask, error) {
	return s.TextTask.GetLastTextTask(teamId)
}

func (s *storage) GetStatusLastTextTask(teamId int) (status bool, err error) {
	return s.TextTask.GetStatusLastTextTask(teamId)
}

func (s *storage) CreateAnswerOnTextTask(task domain.TextTask) error {
	return s.TextTask.CreateAnswerOnTextTask(task)
}

func (s *storage) UpdateAnswerOnTextTask(task domain.TextTask) error {
	return s.TextTask.UpdateAnswerOnTextTask(task)
}
