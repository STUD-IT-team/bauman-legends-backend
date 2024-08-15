package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/storage"
)

type MediaTaskStorage struct {
	db *pgx.Conn
}

func NewMediaTaskStorage(dataSource string) (storage.MediaTaskStorage, error) {
	config, err := pgx.ParseConfig(dataSource)
	if err != nil {
		return nil, err
	}

	db, err := pgx.ConnectConfig(context.Background(), config)

	// db.DB.SetMaxOpenConns(1000) // The default is 0 (unlimited)
	// db.DB.SetMaxIdleConns(10)   // defaultMaxIdleConns = 2
	// db.DB.SetConnMaxLifetime(0) // 0, connections are reused forever.
	return &MediaTaskStorage{
		db: db,
	}, nil
}

func (s *MediaTaskStorage) GetNewMediaTask(teamId int) (task domain.MediaTask, err error) {
	err = s.db.QueryRow(context.TODO(), getNewMediaTask, teamId).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.VideoId,
		&task.VideoKey,
		&task.Points,
	)
	if err != nil {
		log.WithField(
			"origin.function", "GetNewMediaTask",
		).Errorf("Ошибка при выдаче новой таски: %s", err.Error())
		return domain.MediaTask{}, err
	}

	return task, nil
}

func (s *MediaTaskStorage) GetStatusLastMediaTask(teamId int) (status string, err error) {
	row := s.db.QueryRow(context.TODO(), getStatusLastMediaTask, teamId)
	if row.Scan(&status) == pgx.ErrNoRows {
		return consts.CorrectStatus, nil
	}
	if err != nil {
		log.WithField(
			"origin.function", "GetStatusLastMediaTask",
		).Errorf("Ошибка при получении статуса таски: %s", err.Error())
		return consts.EmptyStatus, err
	}
	return status, nil
}

func (s *MediaTaskStorage) GetLastMediaTask(teamId int) (task domain.MediaTask, err error) {
	err = s.db.QueryRow(context.TODO(), getLastMediaTask, teamId).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.VideoId,
		&task.VideoKey,
		&task.Points,
	)
	if err != nil {
		log.WithField(
			"origin.function", "GetLastMediaTask",
		).Errorf("Ошибка при получения последней таски: %s", err.Error())
		return domain.MediaTask{}, err
	}

	return task, nil
}

func (s *MediaTaskStorage) UpdateAnswerOnMediaTask(taskId int, task domain.MediaTask) (err error) {
	tx, err := s.db.Begin(context.TODO())
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnMediaTask",
		).Error("Ошибка в создании транзакции: %s", err)
		return err
	}

	defer tx.Rollback(context.TODO())

	err = tx.QueryRow(context.TODO(), createMediaObjectQuery, task.PhotoKey, "photo-answer").Scan(&task.PhotoId)
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnMediaTask",
		).Errorf("Ошибка создания обьекта: %s", err.Error())
		return err
	}

	_, err = tx.Exec(context.TODO(), updateAnswerOnMediaTask,
		task.PhotoId,
		consts.ReviewStatus,
		taskId,
	)
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnMediaTask",
		).Errorf("Ошибка при отправки ответа: %s", err.Error())
		return err
	}

	err = tx.Commit(context.TODO())
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnMediaTask",
		).Errorf("Ошибка закрыия транзакции: %s", err.Error())
		return err
	}

	return nil
}

func (s *MediaTaskStorage) GetAnswersOnMediaTasksByFilter(status string) (tasks []domain.MediaTask, err error) {
	rows, err := s.db.Query(context.TODO(), getAnswerOnMediaTaskByStatus, status)
	if err != nil {
		log.WithField("", "GetAnswersOnMediaTasksByFilter")
		return nil, err
	}

	for rows.Next() {
		var task domain.MediaTask
		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.PhotoId,
			&task.PhotoKey,
			&task.TeamId,
			&task.Points,
			&task.Comment,
			&task.Status,
		)
		if err != nil {
			log.WithField("", "GetAnswersOnMediaTasksByFilter")
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *MediaTaskStorage) GetAllAnswerOnMediaTasks() (tasks []domain.MediaTask, err error) {
	rows, err := s.db.Query(context.TODO(), getAllMediaTask)
	if err != nil {
		log.WithField("", "GetAllAnswerOnMediaTasks")
		return nil, err
	}

	for rows.Next() {
		var task domain.MediaTask
		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.PhotoId,
			&task.PhotoKey,
			&task.TeamId,
			&task.Points,
			&task.Comment,
			&task.Status,
		)
		if err != nil {
			log.WithField("", "GetAllAnswerOnMediaTasks")
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *MediaTaskStorage) UpdatePointsOnMediaTask(status string, taskId int, points int) (err error) {
	_, err = s.db.Exec(context.TODO(), setPointsOnMediaTask, points, status, taskId)
	if err != nil {
		log.WithField(
			"origin.function", "UpdatePointsOnMediaTask",
		).Errorf("Ошибка при проставлении баллов: %s", err.Error())
		return err
	}

	return nil
}

func (s *MediaTaskStorage) GetPointsOnMediaTask(mediaTaskId int) (points int, err error) {
	err = s.db.QueryRow(context.TODO(), getPointsOnMediaTask, mediaTaskId).Scan(&points)
	if err != nil {
		log.WithField(
			"origin.function", "GetPointsOnMediaTask",
		).Errorf("Ошибка при получении баллов за таску: %s", err.Error())
		return 0, err
	}

	return points, nil
}

func (s *MediaTaskStorage) CreateAnswerOnMediaTask(teamId int, pointTaskId int) (id int, err error) {
	_, err = s.db.Exec(context.TODO(), createAnswerOnMediaTask,
		teamId,
		pointTaskId,
		consts.EmptyStatus,
	)
	if err != nil {
		log.WithField(
			"origin.function", "CreateAnswerOnMediaTask",
		).Errorf("Ошибка в создании ответа: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (s *MediaTaskStorage) GetAnswerOnMediaTaskById(answerMediaTaskId int) (task domain.MediaTask, err error) {
	err = s.db.QueryRow(context.TODO(), getMediaTaskById, answerMediaTaskId).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.PhotoId,
		&task.PhotoKey,
		&task.TeamId,
		&task.Points,
		&task.Comment,
		&task.Status,
	)
	if err != nil {
		log.WithField(
			"origin.function", "GetAnswerOnMediaTaskById",
		).Errorf("Ошибка при получении ответа: %s", err.Error())
		return domain.MediaTask{}, err
	}

	return task, nil
}

func (s *MediaTaskStorage) CheckAnswerOnMediaTaskById(answerMediaTaskId int, teamId int) (exist bool, err error) {
	err = s.db.QueryRow(context.TODO(), checkAnswerIsExistQuery, teamId, answerMediaTaskId).Scan(&exist)
	if err != nil {
		log.WithField(
			"origin.function", "CheckAnswerOnMediaTaskById",
		).Errorf("Ошибка проверки существования ответа: %s", err.Error())
		return false, err
	}

	return exist, nil
}

func (s *MediaTaskStorage) GetUpdateTimeAnswerOnMediaTask(taskId int) (time time.Time, err error) {
	err = s.db.QueryRow(context.TODO(), getUpdateTimeMediaTask, taskId).Scan(&time)
	if err != nil {
		log.WithField(
			"origin.function", "GetUpdateTimeMediaTask",
		).Errorf("ошибка получения даты: %s", err.Error())
		return time, err
	}

	return time, nil
}
