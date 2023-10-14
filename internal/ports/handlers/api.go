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
	Api   *app.Api
	Teams *app.TeamService
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

func (h *HTTPHandler) RegisterTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "RegisterTeam",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.RegisterTeam

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "RegisterTeam",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	req.Session = cookie.Value
	res, err := h.Teams.RegisterTeam(&req)

	if err != nil {
		log.WithField(
			"origin.function", "RegisterTeam",
		).Errorf(
			"Ошибка регистрации команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(res)
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
	w.WriteHeader(http.StatusCreated)
}

func (h *HTTPHandler) ChangeTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "ChangeTeam",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.ChangeTeam

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "ChangeTeam",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	req.Session = cookie.Value

	res, err := h.Teams.UpdateTeam(&req)

	if err != nil {
		log.WithField(
			"origin.function", "RegisterTeam",
		).Errorf(
			"Ошибка регистрации команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	//todo

	err = json.NewEncoder(w).Encode(res)
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

func (h *HTTPHandler) GetTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetTeam",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.GetTeam

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "GetTeam",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	req.Session = cookie.Value

	res, err := h.Teams.GetTeam(&req)
	if err != nil {
		log.WithField(
			"origin.function", "GetTeam",
		).Errorf(
			"Ошибка при получении профиля команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(res)
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

func (h *HTTPHandler) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "DeleteTeam",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.DeleteTeam

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "DeleteTeam",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	req.Session = cookie.Value

	err = h.Teams.DeleteTeam(&req)
	if err != nil {
		log.WithField(
			"origin.function", "DeleteTeam",
		).Errorf(
			"Ошибка при получении профиля команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *HTTPHandler) Invite(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "Invite",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.InviteToTeam

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "InviteToTeam",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	req.Session = cookie.Value
	err = h.Teams.InviteToTeam(&req)

	if err != nil {
		log.WithField(
			"origin.function", "Invite",
		).Errorf(
			"Ошибка регистрации команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (h *HTTPHandler) DeleteMember(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "DeleteMember",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.DeleteFromTeam

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "InviteToTeam",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	req.Session = cookie.Value

	err = h.Teams.DeleteFromTeam(&req)

	if err != nil {
		log.WithField(
			"origin.function", "DeleteMember",
		).Errorf(
			"Ошибка регистрации команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (h *HTTPHandler) UpdateMember(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "UpdateMember",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.UpdateMember

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "UpdateMember",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	req.Session = cookie.Value

	err = h.Teams.UpdateMember(&req)

	if err != nil {
		log.WithField(
			"origin.function", "UpdateMember",
		).Errorf(
			"Ошибка регистрации команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
