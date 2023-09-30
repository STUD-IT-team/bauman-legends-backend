package app

import (
	"context"
	"encoding/json"
	consts "github.com/STUD-IT-team/bauman-legends-backend/internal/app/const"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

// TODO:
// 	 GET /user - получить свой профиль
//	 PUT /user - изменить свои данные

type Api struct {
	AuthClient grpc2.AuthClient
}

func NewApi(conn grpc.ClientConnInterface) *Api {
	return &Api{
		AuthClient: grpc2.NewAuthClient(conn),
	}
}

func (a *Api) Register(w http.ResponseWriter, r *http.Request) {
	var req request.Register
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "Register",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	reqGrpc := mapper.MakeGrpcRequestRegister(&req)

	if reqGrpc == nil {
		log.WithField(
			"origin.function", "Register",
		).Error(
			"Пустой grpc-запрос",
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
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
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	expires, err := time.Parse(consts.GrpcTimeFormat, res.GetExpires())
	if err != nil {
		log.WithField(
			"origin.function", "Register",
		).Errorf(
			"Ошибка чтения времени жизни сессии: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "access-token",
		Value:   res.GetAccessToken(),
		Expires: expires,
	})
}

func (a *Api) Login(w http.ResponseWriter, r *http.Request) {
	var req request.Login
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "Login",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	reqGrpc := mapper.MakeGrpcRequestLogin(&req)

	if reqGrpc == nil {
		log.WithField(
			"origin.function", "Login",
		).Error(
			"Пустой grpc-запрос",
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
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
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	expires, err := time.Parse(consts.GrpcTimeFormat, res.GetExpires())
	if err != nil {
		log.WithField(
			"origin.function", "Login",
		).Errorf(
			"Ошибка чтения времени жизни сессии: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "access-token",
		Value:   res.GetAccessToken(),
		Expires: expires,
	})
}

func (a *Api) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "Logout",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusForbidden)
		return
	}

	req := &grpc2.LogoutRequest{AccessToken: cookie.Value}

	if _, err = a.AuthClient.Logout(context.Background(), req); err != nil {
		log.WithField(
			"origin.function", "Logout",
		).Errorf(
			"Ошибка при выходе из аккаунта: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	cookie.MaxAge = -1
	cookie.Expires = time.Now()

	http.SetCookie(w, cookie)
}

func (a *Api) GetProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetProfile",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusForbidden)
		return
	}

	req := &grpc2.GetProfileRequest{AccessToken: cookie.Value}

	profile, err := a.AuthClient.GetProfile(context.Background(), req)

	if err != nil {
		log.WithField(
			"origin.function", "GetProfile",
		).Errorf(
			"Ошибка при получении профиля пользователя: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(mapper.MakeProfileResponse(profile))
	if err != nil {
		log.WithField(
			"origin.function", "GetProfile",
		).Errorf(
			"Ошибка при отправке профиля пользователя: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (a *Api) ChangeProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "ChangeProfile",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusForbidden)
		return
	}

	var reqBody request.ChangeProfile
	if err = json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.WithField(
			"origin.function", "ChangeProfile",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	req := &grpc2.ChangeProfileRequest{
		AccessToken: cookie.Value,
		Name:        reqBody.Name,
		Group:       reqBody.Group,
		Password:    reqBody.Password,
		Email:       reqBody.Email,
		Telegram:    reqBody.Telegram,
		Vk:          reqBody.VK,
		PhoneNumber: reqBody.PhoneNumber,
	}

	_, err = a.AuthClient.ChangeProfile(context.Background(), req)

	if err != nil {
		log.WithField(
			"origin.function", "ChangeProfile",
		).Errorf(
			"Ошибка при изменении профиля пользователя: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
