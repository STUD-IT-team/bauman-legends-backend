package app

import (
	"context"
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
	if err := r.ParseForm(); err != nil {
		log.WithField(
			"origin.function", "Register",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	req := mapper.MakeGrpcRequestRegister(
		&request.Register{
			Name:          r.PostFormValue("name"),
			Group:         r.PostFormValue("group"),
			Email:         r.PostFormValue("email"),
			Password:      r.PostFormValue("password"),
			Telegram:      r.PostFormValue("telegram"),
			VK:            r.PostFormValue("vk"),
			PhoneNumber:   r.PostFormValue("phone_number"),
			ClientBrowser: r.PostFormValue("client_browser"),
			ClientOS:      r.PostFormValue("client_os"),
		})

	res, err := a.AuthClient.Register(context.Background(), req)
	if err != nil {
		log.WithField(
			"origin.function", "Register",
		).Errorf(
			"Не удалось зарегистрировать пользователя %s: %s",
			req.Name,
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
	if err := r.ParseForm(); err != nil {
		log.WithField(
			"origin.function", "Login",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	req := mapper.MakeGrpcRequestLogin(
		&request.Login{
			Email:         r.FormValue("email"),
			Password:      r.FormValue("password"),
			ClientBrowser: r.FormValue("client_browser"),
			ClientOS:      r.FormValue("client_os"),
		})

	res, err := a.AuthClient.Login(context.Background(), req)
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
