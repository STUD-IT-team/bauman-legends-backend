package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/storage"
)

type UserStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(dataSource string) (storage.UserStorage, error) {
	config, err := pgxpool.ParseConfig(dataSource)
	if err != nil {
		return nil, err
	}
	db, err := pgxpool.NewWithConfig(context.Background(), config)
	// db.DB.SetMaxOpenConns(1000) // The default is 0 (unlimited)
	// db.DB.SetMaxIdleConns(10)   // defaultMaxIdleConns = 2
	// db.DB.SetConnMaxLifetime(0) // 0, connections are reused forever.
	return &UserStorage{
		db: db,
	}, nil
}

func (s *UserStorage) GetAllUsers() ([]domain.Member, error) {
	var users []domain.Member
	rows, err := s.db.Query(context.Background(), getAllUsersQuery)
	if err != nil {
		log.WithField(
			"origin.function", "GetAllUsers",
		).Errorf("ошибка получения пользователей: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		var user domain.Member
		err = rows.Scan(
			&user.ID,
			&user.Role,
			&user.Name,
			&user.Group,
			&user.Email,
			&user.Telegram,
			&user.VK,
			&user.PhoneNumber,
			&user.TeamName)
		if err != nil {
			log.WithField(
				"origin.function", "GetAllUsers",
			).Errorf("ошибка получения пользователей: %s", err.Error())
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *UserStorage) GetUserWithoutTeam() ([]domain.Member, error) {
	var users []domain.Member
	rows, err := s.db.Query(context.Background(), getUserWithoutTeam)
	if err != nil {
		log.WithField(
			"origin.function", "GetAllUsers",
		).Errorf("ошибка получения пользователей: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		var user domain.Member
		err = rows.Scan(
			&user.ID,
			&user.Role,
			&user.Name,
			&user.Group,
			&user.Email,
			&user.Telegram,
			&user.VK,
			&user.PhoneNumber,
			&user.TeamName)
		if err != nil {
			log.WithField(
				"origin.function", "GetAllUsers",
			).Errorf("ошибка получения пользователей: %s", err.Error())
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
func (s *UserStorage) GetUserWithCountTeam(count int) ([]domain.Member, error) {
	var users []domain.Member
	rows, err := s.db.Query(context.Background(), getUserWithCountTeamQuery, count)
	if err != nil {
		log.WithField(
			"origin.function", "GetAllUsers",
		).Errorf("ошибка получения пользователей: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		var user domain.Member
		err = rows.Scan(
			&user.ID,
			&user.Role,
			&user.Name,
			&user.Group,
			&user.Email,
			&user.Telegram,
			&user.VK,
			&user.PhoneNumber,
			&user.TeamName)
		if err != nil {
			log.WithField(
				"origin.function", "GetAllUsers",
			).Errorf("ошибка получения пользователей: %s", err.Error())
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *UserStorage) GetUserById(id int) (domain.Member, error) {
	var user domain.Member
	err := s.db.QueryRow(context.Background(), getUserByIdQuery, id).Scan(
		&user.ID,
		&user.Role,
		&user.Name,
		&user.Group,
		&user.Email,
		&user.Telegram,
		&user.VK,
		&user.PhoneNumber,
		&user.TeamName)
	if err != nil {
		log.WithField(
			"origin.function", "GetUserById",
		).Errorf("ошибка получения пользователя: %s", err.Error())
		return domain.Member{}, err
	}

	return user, nil
}
