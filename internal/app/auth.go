package app

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	cache2 "github.com/STUD-IT-team/bauman-legends-backend/internal/adapters/cache"
	consts "github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
	"github.com/STUD-IT-team/bauman-legends-backend/pkg/cache"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExists = errors.New("user exists")
	ErrUserNotFound      = errors.New("user does not exist")
	ErrInvalidPassword   = errors.New("incorrect password")
	ErrSessionNotFound   = errors.New("session not found")
)

type Auth struct {
	SessionCache cache.ICache[string, cache2.Session]
	Repository   repository.IUserAuthStorage
	grpc2.UnimplementedAuthServer
}

func NewAuth(sc cache.ICache[string, cache2.Session], r repository.IUserAuthStorage) *Auth {
	return &Auth{
		SessionCache: sc,
		Repository:   r,
	}
}

func (s *Auth) Register(_ context.Context, req *grpc2.RegisterRequest) (*grpc2.RegisterResponse, error) {
	mappedReq := mapper.MakeRequestRegister(req)
	exists, err := s.Repository.CheckUser(mappedReq.Email)
	if err != nil {
		log.WithField(
			"origin.function", "Register",
		).Errorf("Не удалось проверить наличие пользователя в базе: %s", err.Error())
		return nil, err
	}

	if exists {
		log.WithField(
			"origin.function", "Register",
		).Warnf("Пользователь с логином %s уже существует", mappedReq.Email)
		return nil, ErrUserAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		log.WithField(
			"origin.function", "Register",
		).Errorf("Не удалось создать хэш пароля пользователя: %s", err.Error())
		return nil, err
	}
	mappedReq.Password = string(hash)

	userID, err := s.Repository.CreateUser(*mappedReq)
	if err != nil {
		log.WithField(
			"origin.function", "Register",
		).Errorf("Не удалось зарегистрировать пользователя: %s", err.Error())
		return nil, err
	}

	sessionID := uuid.NewString()
	var sessionDuration int
	if sessionDuration, err = strconv.Atoi(os.Getenv("SESSION_DURATION_HOURS")); err != nil {
		sessionDuration = 12
	}
	value := cache2.Session{
		UserID:        userID,
		ExpireAt:      time.Now().Add(time.Hour * time.Duration(sessionDuration)),
		EnteredAt:     time.Now(),
		ClientBrowser: req.ClientBrowser,
		ClientOS:      req.ClientOS,
	}

	s.SessionCache.Put(sessionID, value)

	log.WithField(
		"origin.function", "Register",
	).Infof("Пользователь %s зарегистрирован", mappedReq.Email)

	return &grpc2.RegisterResponse{
		AccessToken: sessionID,
		Expires:     value.ExpireAt.Format(consts.GrpcTimeFormat),
	}, nil
}

func (s *Auth) Login(_ context.Context, req *grpc2.LoginRequest) (*grpc2.LoginResponse, error) {
	exists, err := s.Repository.CheckUser(req.Email)
	if err != nil {
		log.WithField(
			"origin.function", "Login",
		).Errorf("Не удалось проверить наличие пользователя в базе: %s", err.Error())
		return nil, err
	}

	if !exists {
		log.WithField(
			"origin.function", "Login",
		).Warnf("Пользователь с логином %s не найден", req.Email)
		return nil, ErrUserNotFound
	}

	hashedPassword, err := s.Repository.GetUserPassword(req.Email)
	if err != nil {
		log.WithField(
			"origin.function", "Login",
		).Errorf("Не удалось получить пароль пользователя: %s", err.Error())
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		log.WithField(
			"origin.function", "Login",
		).Warn("Неверный пароль")
		return nil, ErrInvalidPassword
	}

	userID, err := s.Repository.GetUserID(req.Email)
	if err != nil {
		log.WithField(
			"origin.function", "Login",
		).Errorf("Не удалось получить идентификатор пользователя: %s", err.Error())
		return nil, err
	}

	sessionID := uuid.NewString()
	var sessionDuration int
	if sessionDuration, err = strconv.Atoi(os.Getenv("SESSION_DURATION_HOURS")); err != nil {
		sessionDuration = 12
	}
	value := cache2.Session{
		UserID:        userID,
		ExpireAt:      time.Now().Add(time.Hour * time.Duration(sessionDuration)),
		EnteredAt:     time.Now(),
		ClientBrowser: req.ClientBrowser,
		ClientOS:      req.ClientOS,
	}

	s.SessionCache.Put(sessionID, value)

	log.WithField(
		"origin.function", "Login",
	).Infof("Пользователь %s осуществил вход", req.Email)

	return &grpc2.LoginResponse{
		AccessToken: sessionID,
		Expires:     value.ExpireAt.Format(consts.GrpcTimeFormat),
	}, nil
}

func (s *Auth) Logout(_ context.Context, req *grpc2.LogoutRequest) (*grpc2.EmptyResponse, error) {
	accessToken := req.GetAccessToken()
	s.SessionCache.Delete(accessToken)

	return &grpc2.EmptyResponse{}, nil
}

func (s *Auth) Check(_ context.Context, req *grpc2.CheckRequest) (*grpc2.CheckResponse, error) {
	accessToken := req.GetAccessToken()
	record := s.SessionCache.Find(accessToken)

	if record == nil ||
		record.ExpireAt.Before(time.Now()) {
		return &grpc2.CheckResponse{
			Valid: false,
		}, nil
	}

	return &grpc2.CheckResponse{
		Valid: true,
	}, nil
}

func (s *Auth) GetProfile(_ context.Context, req *grpc2.GetProfileRequest) (*grpc2.GetProfileResponse, error) {
	accessToken := req.GetAccessToken()
	session := s.SessionCache.Find(accessToken)

	if session == nil {
		log.WithField(
			"origin.function", "GetProfile",
		).Errorf("Сессия %s не найдена", req.AccessToken)
		return nil, ErrSessionNotFound
	}

	profile, err := s.Repository.GetUserProfile(session.UserID)

	if err != nil {
		log.WithField(
			"origin.function", "GetProfile",
		).Errorf("Ошибка при получении профиля пользователя: %s", err.Error())
		return nil, err
	}

	return mapper.MakeGrpcResponseProfile(profile), nil
}

func (s *Auth) ChangeProfile(_ context.Context, req *grpc2.ChangeProfileRequest) (*grpc2.EmptyResponse, error) {
	accessToken := req.GetAccessToken()
	session := s.SessionCache.Find(accessToken)

	if session == nil {
		log.WithField(
			"origin.function", "ChangeProfile",
		).Errorf("Сессия %s не найдена", req.AccessToken)
		return nil, ErrSessionNotFound
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		log.WithField(
			"origin.function", "ChangeProfile",
		).Errorf("Не удалось создать хэш пароля пользователя: %s", err.Error())
		return nil, err
	}
	req.Password = string(hash)

	err = s.Repository.ChangeUserProfile(session.UserID, mapper.MakeChangeProfileRequest(req))

	if err != nil {
		log.WithField(
			"origin.function", "ChangeProfile",
		).Errorf("Ошибка при изменении профиля пользователя: %s", err.Error())
		return nil, err
	}

	return &grpc2.EmptyResponse{}, nil
}

func (s *Auth) ChangePassword(_ context.Context, req *grpc2.ChangePasswordRequest) (*grpc2.EmptyResponse, error) {
	accessToken := req.GetAccessToken()
	session := s.SessionCache.Find(accessToken)

	if session == nil {
		log.WithField(
			"origin.function", "ChangePassword",
		).Errorf("Сессия %s не найдена", req.AccessToken)
		return nil, ErrSessionNotFound
	}

	curHashPassword, err := s.Repository.GetUserPasswordById(session.UserID)
	if err != nil {
		log.WithField(
			"origin.function", "ChangePassword",
		).Errorf("Не удалось получить текущий пароль пользователя: %s", err.Error())
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(curHashPassword), []byte(req.OldPassword))
	if err != nil {
		log.WithField(
			"origin.function", "ChangePassword",
		).Errorf("Пароли не совпадают: %s", err.Error())
		return nil, ErrInvalidPassword
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 8)
	if err != nil {
		log.WithField(
			"origin.function", "ChangePassword",
		).Errorf("Не удалось создать хэш нового пароля пользователя: %s", err.Error())
		return nil, err
	}
	req.NewPassword = string(hash)

	err = s.Repository.ChangeUserPassword(session.UserID, req.NewPassword)

	if err != nil {
		log.WithField(
			"origin.function", "ChangePassword",
		).Errorf("Ошибка при изменении пароля пользователя: %s", err.Error())
		return nil, err
	}

	return &grpc2.EmptyResponse{}, nil
}
