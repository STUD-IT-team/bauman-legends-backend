package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/storage"
)

type TextTaskStorage struct {
	db *pgx.Conn
}

func NewTextTaskStorage(dataSource string) (storage.TextTaskStorage, error) {
	config, err := pgx.ParseConfig(dataSource)
	if err != nil {
		return nil, err
	}

	db, err := pgx.ConnectConfig(context.Background(), config)

	// db.DB.SetMaxOpenConns(1000) // The default is 0 (unlimited)
	// db.DB.SetMaxIdleConns(10)   // defaultMaxIdleConns = 2
	// db.DB.SetConnMaxLifetime(0) // 0, connections are reused forever.
	return &TextTaskStorage{
		db: db,
	}, nil
}

func (s *TextTaskStorage) GetNewTextTask(teamId int) (domain.TextTask, error) {
	var task domain.TextTask
	err := s.db.QueryRow(context.TODO(), getTextTaskQuery, teamId).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Answer,
		&task.Points,
	)
	if err != nil {
		log.WithField(
			"origin.function", "GetNewTextTask",
		).Errorf("Ошибка при выдаче новой таски: %s", err.Error())
		return domain.TextTask{}, errors.Join(consts.InternalServerError, err)
	}

	return task, nil
}

func (s *TextTaskStorage) CreateAnswerOnTextTask(task domain.TextTask) error {
	_, err := s.db.Exec(context.TODO(), createAnswerOnTextTask,
		task.TeamId,
		task.ID,
		task.Status,
	)
	if err != nil {
		log.WithField(
			"origin.function", "CreateAnswerOnTextTask",
		).Errorf("Ошибка при создания сущности ответа: %s", err.Error())
		return err
	}

	return nil
}

func (s *TextTaskStorage) UpdateAnswerOnTextTask(task domain.TextTask) error {
	_, err := s.db.Exec(context.TODO(), setAnswerOnTextTask,
		task.Answer,
		task.Points,
		task.Status,
		task.ID)
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnTextTask",
		).Errorf("Ошибка при отправке ответа на таску: %s", err.Error())
		return err
	}

	return nil
}

func (s *TextTaskStorage) GetStatusLastTextTask(teamId int) (status string, err error) {
	err = s.db.QueryRow(context.TODO(), getStatusLastTextTask, teamId).Scan(&status)
	if err != nil {
		return consts.CorrectStatus, nil
	}

	return status, nil
}

func (s *TextTaskStorage) GetLastTextTask(teamId int) (domain.TextTask, error) {
	var task domain.TextTask
	err := s.db.QueryRow(context.TODO(), getLastTextTask, teamId).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Answer,
		&task.Points,
	)
	if err != nil {
		log.WithField(
			"origin.function", "GetLastTextTask",
		).Errorf("Ошибка в получении последней таски: %s", err.Error())
		return domain.TextTask{}, err
	}

	return task, nil
}
