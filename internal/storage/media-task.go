package storage

import (
	"time"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
)

type MediaTaskStorage interface {
	GetNewMediaTask(teamId int) (task domain.MediaTask, err error)
	GetStatusLastMediaTask(teamId int) (status string, err error)
	GetLastMediaTask(teamId int) (task domain.MediaTask, err error)
	UpdateAnswerOnMediaTask(taskId int, task domain.MediaAnswer) (err error)
	GetAnswersOnMediaTasksByFilter(status string) (tasks []domain.MediaTask, err error)
	UpdatePointsOnMediaTask(status string, taskId int, points int, comment string) (err error)
	GetPointsOnMediaTask(mediaTaskId int) (points int, err error)
	CreateAnswerOnMediaTask(teamId int, pointTaskId int) (id int, err error)
	GetAnswerOnMediaTaskById(answerMediaTaskId int) (task domain.MediaTask, err error)
	CheckAnswerOnMediaTaskById(answerMediaTaskId int, teamId int) (exist bool, err error)
	GetAllAnswerOnMediaTasks() (tasks []domain.MediaTask, err error)
	GetUpdateTimeAnswerOnMediaTask(taskId int) (time time.Time, err error)
	GetAllMediaTaskByTeam(teamId int) (tasks []domain.MediaTask, err error)
	GetMediaTaskByTeamById(teamId int, answerId int) (task domain.MediaTask, err error)
}

func (s *storage) GetNewMediaTask(teamId int) (task domain.MediaTask, err error) {
	return s.MediaTask.GetNewMediaTask(teamId)
}

func (s *storage) GetStatusLastMediaTask(teamId int) (status string, err error) {
	return s.MediaTask.GetStatusLastMediaTask(teamId)
}

func (s *storage) GetLastMediaTask(teamId int) (task domain.MediaTask, err error) {
	return s.MediaTask.GetLastMediaTask(teamId)
}

func (s *storage) UpdateAnswerOnMediaTask(taskId int, task domain.MediaAnswer) (err error) {
	return s.MediaTask.UpdateAnswerOnMediaTask(taskId, task)
}

func (s *storage) GetAnswersOnMediaTasksByFilter(status string) (tasks []domain.MediaTask, err error) {
	return s.MediaTask.GetAnswersOnMediaTasksByFilter(status)
}

func (s *storage) UpdatePointsOnMediaTask(status string, taskId int, points int, comment string) (err error) {
	return s.MediaTask.UpdatePointsOnMediaTask(status, taskId, points, comment)
}

func (s *storage) GetPointsOnMediaTask(mediaTaskId int) (points int, err error) {
	return s.MediaTask.GetPointsOnMediaTask(mediaTaskId)
}

func (s *storage) CreateAnswerOnMediaTask(teamId int, pointTaskId int) (id int, err error) {
	return s.MediaTask.CreateAnswerOnMediaTask(teamId, pointTaskId)
}

func (s *storage) GetAnswerOnMediaTaskById(answerMediaTaskId int) (task domain.MediaTask, err error) {
	return s.MediaTask.GetAnswerOnMediaTaskById(answerMediaTaskId)
}

func (s *storage) CheckAnswerOnMediaTaskById(answerMediaTaskId int, teamId int) (exist bool, err error) {
	return s.MediaTask.CheckAnswerOnMediaTaskById(answerMediaTaskId, teamId)
}

func (s *storage) GetAllAnswerOnMediaTasks() (task []domain.MediaTask, err error) {
	return s.MediaTask.GetAllAnswerOnMediaTasks()
}

func (s *storage) GetUpdateTimeAnswerOnMediaTask(taskId int) (time time.Time, err error) {
	return s.MediaTask.GetUpdateTimeAnswerOnMediaTask(taskId)
}

func (s *storage) GetAllMediaTaskByTeam(teamId int) (tasks []domain.MediaTask, err error) {
	return s.MediaTask.GetAllMediaTaskByTeam(teamId)
}

func (s *storage) GetMediaTaskByTeamById(teamId int, answerId int) (task domain.MediaTask, err error) {
	return s.MediaTask.GetMediaTaskByTeamById(teamId, answerId)
}
