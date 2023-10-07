package handlers

import (
	"encoding/json"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app"
	consts "github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type HTTPHandler struct {
	Api *app.Api
}

func NewHTTPHandler(api *app.Api) *HTTPHandler {
	return &HTTPHandler{Api: api}
}

func (h *HTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
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

	res, err := h.Api.Register(&req)

	if err != nil {
		log.WithField(
			"origin.function", "Register",
		).Errorf(
			"Ошибка регистрации пользователя: %s",
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

func (h *HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
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

	res, err := h.Api.Login(&req)

	if err != nil {
		log.WithField(
			"origin.function", "Login",
		).Errorf(
			"Ошибка входа: %s",
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

func (h *HTTPHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "Logout",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	req := &grpc2.LogoutRequest{AccessToken: cookie.Value}

	if err = h.Api.Logout(req); err != nil {
		log.WithField(
			"origin.function", "Logout",
		).Errorf(
			"Ошибка при выходе из аккаунта: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	cookie.Value = ""
	cookie.MaxAge = -1

	http.SetCookie(w, cookie)
}

func (h *HTTPHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetProfile",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	req := &grpc2.GetProfileRequest{AccessToken: cookie.Value}

	profile, err := h.Api.GetProfile(req)

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

	err = json.NewEncoder(w).Encode(profile)
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

func (h *HTTPHandler) ChangeProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "ChangeProfile",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
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

	err = h.Api.ChangeProfile(req)

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
