package app

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
)

type Api struct {
	AuthClient grpc2.AuthClient
}

func NewApi(conn grpc.ClientConnInterface) *Api {
	return &Api{
		AuthClient: grpc2.NewAuthClient(conn),
	}
}

func (a *Api) Register(req *request.Register) (*grpc2.RegisterResponse, error) {
	reqGrpc := mapper.MakeGrpcRequestRegister(req)

	if reqGrpc == nil {
		log.WithField(
			"origin.function", "Register",
		).Error(
			"Пустой grpc-запрос",
		)
		return nil, errors.New("пустой grpc-запрос")
	}

	res, err := a.AuthClient.Register(context.Background(), reqGrpc)
	if err != nil {
		log.WithField(
			"origin.function", "Register",
		).Errorf(
			"Не удалось зарегистрировать пользователя %s: %s",
			reqGrpc.Name,
			err.Error(),
		)
		return nil, err
	}

	return res, nil
}

func (a *Api) Login(req *request.Login) (*grpc2.LoginResponse, error) {
	reqGrpc := mapper.MakeGrpcRequestLogin(req)

	if reqGrpc == nil {
		log.WithField(
			"origin.function", "Login",
		).Error(
			"Пустой grpc-запрос",
		)
		return nil, errors.New("пустой grpc-запрос")
	}

	res, err := a.AuthClient.Login(context.Background(), reqGrpc)
	if err != nil {
		log.WithField(
			"origin.function", "Login",
		).Errorf(
			"Не удалось войти в аккаунт %s: %s",
			req.Email,
			err.Error(),
		)
		return nil, err
	}

	return res, nil
}

func (a *Api) Logout(req *grpc2.LogoutRequest) error {
	if _, err := a.AuthClient.Logout(context.Background(), req); err != nil {
		log.WithField(
			"origin.function", "Logout",
		).Errorf(
			"Ошибка при выходе из аккаунта: %s",
			err.Error(),
		)
		return err
	}

	return nil
}

func (a *Api) GetProfile(req *grpc2.GetProfileRequest) (*response.UserProfile, error) {
	profile, err := a.AuthClient.GetProfile(context.Background(), req)

	if err != nil {
		log.WithField(
			"origin.function", "GetProfile",
		).Errorf(
			"Ошибка при получении профиля пользователя: %s",
			err.Error(),
		)
		return nil, err
	}

	return mapper.MakeProfileResponse(profile), nil
}

func (a *Api) ChangeProfile(req *grpc2.ChangeProfileRequest) error {
	_, err := a.AuthClient.ChangeProfile(context.Background(), req)

	if err != nil {
		log.WithField(
			"origin.function", "ChangeProfile",
		).Errorf(
			"Ошибка при изменении профиля пользователя: %s",
			err.Error(),
		)
		return err
	}

	return nil
}

func (a *Api) ChangePassword(req *grpc2.ChangePasswordRequest) error {
	_, err := a.AuthClient.ChangePassword(context.Background(), req)

	if err != nil {
		log.WithField(
			"origin.function", "ChangePassword",
		).Errorf(
			"Ошибка при изменении пароля пользователя: %s",
			err.Error(),
		)
		return err
	}

	return nil
}
