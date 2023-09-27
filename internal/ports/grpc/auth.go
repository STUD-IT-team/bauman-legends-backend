package grpc

import (
	"context"
	"errors"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/adapters/cache"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/repository"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/session"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"os"
	"strconv"
	"time"
)

func NewServer(opt ...grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(opt...)
}

type Auth struct {
	SessionCache cache.ICache[string, session.Session]
	Repository   repository.IRepository
}

func NewAuth(sc cache.ICache[string, session.Session], r repository.IRepository) *Auth {
	return &Auth{
		SessionCache: sc,
		Repository:   r,
	}
}

func (s *Auth) mustEmbedUnimplementedAuthServer() {}

func (s *Auth) Register(_ context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	mappedReq := req.MakeRequestRegister()
	exists, err := s.Repository.CheckUser(req.Email)
	if err != nil {
		log.WithField(
			"origin.function", "Register",
		).Errorf("Не удалось проверить наличие пользователя в базе: %s", err.Error())
		return nil, err
	}

	if exists {
		log.WithField(
			"origin.function", "Register",
		).Warnf("Пользователь с логином %s уже существует", req.Email)
		return nil, errors.New("user exists")
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
	value := session.Session{
		UserID:        userID,
		ExpireAt:      time.Now().Add(time.Hour * time.Duration(sessionDuration)),
		EnteredAt:     time.Now(),
		ClientIP:      req.ClientIP,
		ClientBrowser: req.ClientBrowser,
		ClientOS:      req.ClientOS,
	}

	s.SessionCache.Put(sessionID, value)

	log.WithField(
		"origin.function", "Register",
	).Infof("Пользователь %s зарегистрирован", req.Email)

	return &RegisterResponse{AccessToken: sessionID}, nil
}

func (s *Auth) Login(_ context.Context, req *LoginRequest) (*LoginResponse, error) {
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
		return nil, errors.New("user does not exist")
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
		return nil, errors.New("incorrect password")
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
	value := session.Session{
		UserID:        userID,
		ExpireAt:      time.Now().Add(time.Hour * time.Duration(sessionDuration)),
		EnteredAt:     time.Now(),
		ClientIP:      req.ClientIP,
		ClientBrowser: req.ClientBrowser,
		ClientOS:      req.ClientOS,
	}

	s.SessionCache.Put(sessionID, value)

	log.WithField(
		"origin.function", "Login",
	).Infof("Пользователь %s осуществил вход", req.Email)

	return &LoginResponse{AccessToken: sessionID}, nil
}

func (s *Auth) Logout(_ context.Context, req *LogoutRequest) (*LogoutResponse, error) {
	accessToken := req.GetAccessToken()
	s.SessionCache.Delete(accessToken)

	return &LogoutResponse{
		Message: "success",
	}, nil
}

func (s *Auth) Check(_ context.Context, req *CheckRequest) (*CheckResponse, error) {
	accessToken := req.GetAccessToken()
	record := s.SessionCache.Find(accessToken)

	if record == nil ||
		record.ExpireAt.Before(time.Now()) {
		return &CheckResponse{
			Valid: false,
		}, nil
	}

	return &CheckResponse{
		Valid: true,
	}, nil
}

// MakeRequestRegister
//
// Преобразование grpc-запроса регистрации в структуру
func (r *RegisterRequest) MakeRequestRegister() *request.Register {
	return &request.Register{
		Name:          r.Name,
		Group:         r.Group,
		Email:         r.Email,
		Password:      r.Password,
		Telegram:      r.Telegram,
		VK:            r.Vk,
		PhoneNumber:   r.PhoneNumber,
		ClientBrowser: r.ClientBrowser,
		ClientOS:      r.ClientOS,
	}
}
