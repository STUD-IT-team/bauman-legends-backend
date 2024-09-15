package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/app"
	consts "github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	_ "github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
)

type HTTPHandler struct {
	Api       *app.Api
	Teams     *app.TeamService
	TextTask  *app.TextTaskService
	MediaTask *app.MediaTaskService
	Users     *app.UserService
	Secs      *app.SECService
}

func NewHTTPHandler(api *app.Api, teams *app.TeamService, textTask *app.TextTaskService, mediaTask *app.MediaTaskService, user *app.UserService, sec *app.SECService) *HTTPHandler {
	return &HTTPHandler{
		Api:       api,
		Teams:     teams,
		TextTask:  textTask,
		MediaTask: mediaTask,
		Users:     user,
		Secs:      sec,
	}
}

// CreateUser
// @Summary      Register
// @Description  Register user in the system
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request.Register	body		request.Register	true	"Add account"
// @Success      200  {string}  string    "ok"
// @Failure      400  {string}  string    "bad request"
// @Failure		 409  {string}  string    "user exists"
// @Failure      500  {string}  string    "internal server error"
// @Router       /user/auth/register [post]
func (h *HTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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
		if status.Convert(err).Message() == app.ErrUserAlreadyExists.Error() {
			log.WithField(
				"origin.function", "Register",
			).Errorf(
				"Пользователь уже существует: %s",
				err.Error(),
			)
			http.Error(w, "user exists", http.StatusConflict)
		} else {
			log.WithField(
				"origin.function", "Register",
			).Errorf(
				"Ошибка регистрации пользователя: %s",
				err.Error(),
			)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
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
		Name:     "access-token",
		Value:    res.GetAccessToken(),
		Expires:  expires,
		HttpOnly: true,
		Path:     "/",
	})
}

// Login
// @Summary      Login
// @Description  Login in the system
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request.Login	body		request.Login	true	"Login account"
// @Success      200  {string}  string    "ok"
// @Failure      400  {string}  string    "bad request"
// @Failure      401  {string}  string    "user does not exist"
// @Failure      500  {string}  string    "internal server error"
// @Router       /user/auth/login [post]
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
		if status.Convert(err).Message() == app.ErrUserNotFound.Error() || status.Convert(err).Message() == app.ErrInvalidPassword.Error() {
			log.WithField(
				"origin.function", "Login",
			).Errorf(
				"Пользователь не найден: %s",
				err.Error(),
			)
			http.Error(w, "user does not exist", http.StatusUnauthorized)
		} else {
			log.WithField(
				"origin.function", "Login",
			).Errorf(
				"Ошибка входа: %s",
				err.Error(),
			)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
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
		Name:     "access-token",
		Value:    res.GetAccessToken(),
		Expires:  expires,
		HttpOnly: true,
		Path:     "/",
	})
}

// Logout
// @Summary      Logout
// @Description  Logout in the system
// @Tags         user
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Success      200  {string}  string    "ok"
// @Failure      401  {string}  string    "not authorized"
// @Failure      500  {string}  string    "internal server error"
// @Router       /user/logout [delete]
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

// GetProfile
// @Summary      GetProfile
// @Description  Get the user information about yourself
// @Tags         user
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Success      200  {object}      response.UserProfile
// @Failure      401  {string}  string    "not authorized"
// @Failure      500  {string}  string    "internal server error"
// @Router       /user [get]
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

// UpdateProfile
// @Summary      UpdateProfile
// @Description  Update the user information about yourself
// @Tags         user
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        request.ChangeProfile	body		request.ChangeProfile	true	"Change profile"
// @Success      200  {string}  string    "ok"
// @Failure      400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      500  {string}  string    "internal server error"
// @Router       /user [put]
func (h *HTTPHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
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

// GetUsersByFilter
// @Summary      GetUsersByFilter
// @Description  Get info about users by filter (only admin)
// @Tags         user
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        team         query     bool        false  "boolean team"
// @Param        count        query     int true "count members in team"
// @Success      200  {string}  string    "ok"
// @Failure      400  {string}  string    "bad request"
// @Failure		 403  {string}  string	  "not rights"
// @Failure      401  {string}  string    "not authorized"
// @Failure      500  {string}  string    "internal server error"
// @Router       /admin/user [get]
func (h *HTTPHandler) GetUsersByFilter(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetUsersByFilter",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	team := r.URL.Query().Get("team")
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	req := request.GetUsersByFilter{
		CountInTeam: count,
		WithTeam:    team != "false",
	}

	res, err := h.Users.GetUsersByFilter(request.Session{Value: cookie.Value}, req)
	if err != nil {
		log.WithField(
			"origin.function", "GetUsersByFilter",
		).Errorf(err.Error())
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetUsersByFilter",
		).Errorf(
			"Ошибка при отправке данных пользователей: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetUserById
// @Summary      GetUserById
// @Description  Get info about users by id (only admin)
// @Tags         user
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        id path string true "user ID"
// @Success      200  {string}  string    "ok"
// @Failure      400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure		 403  {string}  string	  "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /admin/user/{id} [get]
func (h *HTTPHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetUserById",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.WithField(
			"origin.function", "GetUserById",
		).Errorf(err.Error())
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	req := request.GetUserById{
		Id: id,
	}

	res, err := h.Users.GetUserById(request.Session{Value: cookie.Value}, req)
	if err != nil {
		log.WithField(
			"origin.function", "GetUserById",
		).Errorf(err.Error())
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetUserById",
		).Errorf(
			"Ошибка при отправке данных пользователя: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

/* func (h *HTTPHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "ChangePassword",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var reqBody request.ChangePassword
	if err = json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.WithField(
			"origin.function", "ChangePassword",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	req := &grpc2.ChangePasswordRequest{
		AccessToken: cookie.Value,
		OldPassword: reqBody.OldPassword,
		NewPassword: reqBody.NewPassword,
	}

	err = h.Api.ChangePassword(req)

	if err != nil {
		if status.Convert(err).Message() == app.ErrInvalidPassword.Error() {
			log.WithField(
				"origin.function", "ChangePassword",
			).Errorf(
				"Текущий пароль не совпадает с введенным: %s",
				err.Error(),
			)
			http.Error(w, "not authorized", http.StatusUnauthorized)
		} else {
			log.WithField(
				"origin.function", "ChangePassword",
			).Errorf(
				"Ошибка при изменении пароля пользователя: %s",
				err.Error(),
			)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}
}*/

// CreateTeam
// @Summary      CreateTeam
// @Description  User creates team and stay captain of team
// @Tags         team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        request.CreateTeam	body		request.CreateTeam	true	"Add team"
// @Success      201  {object}  response.CreateTeam
// @Failure      400  {string}  string    "bad request"
// @Failure 	 401  {string}  string    "not authorized"
// @Failure		 409  {string}  string    "user exists"
// @Failure      500  {string}  string    "internal server error"
// @Router       /team [post]
func (h *HTTPHandler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "CreateTeam",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.CreateTeam
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "CreateTeam",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.Teams.CreateTeam(req, request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "CreateTeam",
		).Errorf(
			"Ошибка регистрации команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "CreateTeam",
		).Errorf(
			"Ошибка при отправке профиля пользователя: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetTeam
// @Summary      GetTeam
// @Description  User gets info about team
// @Tags         team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Success      200  {object}  response.GetTeam
// @Failure      401  {string}  string    "not authorized"
// @Failure      500  {string}  string    "internal server error"
// @Router       /team [get]
func (h *HTTPHandler) GetTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "UpdateTeam",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	res, err := h.Teams.GetTeam(request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "GetTeam",
		).Errorf(
			"Ошибка при получении данных о команде: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetProfile",
		).Errorf(
			"Ошибка при отправке профиля пользователя: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UpdateTeam
// @Summary      UpdateTeam
// @Description  Update info about team (only capitan of team)
// @Tags         team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        request.UpdateTeam	body		request.UpdateTeam	true	"Change team"
// @Success      200  {string}  string    "ok"
// @Failure      400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      500  {string}  string    "internal server error"
// @Router       /team [put]
func (h *HTTPHandler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "UpdateTeam",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.UpdateTeam

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "UpdateTeam",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.Teams.UpdateTeam(&req, request.Session{Value: cookie.Value})

	if err != nil {
		log.WithField(
			"origin.function", "UpdateTeam",
		).Errorf(
			"Ошибка регистрации команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetProfile",
		).Errorf(
			"Ошибка при отправке профиля пользователя: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteTeam
// @Summary      DeleteTeam
// @Description  Delete team if team is empty (only capitan of team)
// @Tags         team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Success      200  {string}  string    "ok"
// @Failure      401  {string}  string    "not authorized"
// @Failure      500  {string}  string    "internal server error"
// @Router       /team [delete]
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

	err = h.Teams.DeleteTeam(request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "DeleteTeam",
		).Errorf(
			"Ошибка удалении команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// AddUserInTeam
// @Summary		 AddUserInTeam
// @Description  Capitan of team can add user in his team
// @Tags		 team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param 		 request.AddMemberToTeam  body  request.AddMemberToTeam  true  "Invite to team"
// @Success      200  {string}  string    "ok"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /team/member [post]
func (h *HTTPHandler) AddUserInTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "AddUserInTeam",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.AddMemberToTeam

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "AddUserInTeam",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = h.Teams.AddMemberToTeam(&req, request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "AddUserInTeam",
		).Errorf(
			"Ошибка добавлении участника в команду: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteUserFromTeam
// @Summary		 DeleteUserFromTeam
// @Description  User can be removed from the team (not captain)
// @Tags		 team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param 		 request.DeleteMemberFromTeam  body  request.DeleteMemberFromTeam  true  "Delete from team"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      500  {string}  string    "internal server error"
// @Router       /team/member [delete]
func (h *HTTPHandler) DeleteUserFromTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "DeleteUserFromTeam",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.DeleteMemberFromTeam

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "DeleteUserFromTeam",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = h.Teams.DeleteMemberFromTeam(&req, request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "DeleteUserFromTeam",
		).Errorf(
			"Ошибка удалении участника из команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetTeamsByFilter
// @Summary		 GetTeamsByFilter
// @Description  Get all info about team
// @Tags		 team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        count         query     int        false  "int count"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /admin/team [get]
func (h *HTTPHandler) GetTeamsByFilter(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetTeamsByFilter",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	membersCount := r.URL.Query().Get("members_count")
	count, err := strconv.Atoi(membersCount)
	if err != nil {
		log.WithField(
			"origin.function", "GetTeamsByFilter",
		).Errorf(
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.Teams.GetTeamsByFilter(&request.GetTeamsByFilter{MembersCount: count}, request.Session{Value: cookie.Value})

	if err != nil {
		log.WithField(
			"origin.function", "GetTeamsByFilter",
		).Errorf(
			err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetProfile",
		).Errorf(
			"Ошибка при отправке профиля пользователя: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

// GetTeamById
// @Summary		 GetTeamById
// @Description  Get info about team by id
// @Tags		 team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        id path string true "team ID"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /admin/team/{id} [get]
func (h *HTTPHandler) GetTeamById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetTeamById",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.GetTeamById

	if err = req.Bind(r); err != nil {
		log.WithField(
			"origin.function", "GetTeamById",
		).Errorf("Ошибка чтения запроса: %s",
			err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.Teams.GetTeamById(&req, request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "GetTeamById",
		).Errorf(
			"Ошибка удалении команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetTeamById",
		).Errorf(
			"Ошибка при отправке данных команды: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// SpendPointsTeam
// @Summary		 SpendPointsTeam
// @Description
// @Tags		 team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param  		 request.UpdateSpendPoints  body  request.UpdateSpendPoints  true  "Spend points team"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /admin/team/{id}/point/spend [put]
func (h *HTTPHandler) SpendPointsTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "SpendPointsTeam",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.UpdateSpendPoints

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "SpendPointsTeam",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.Teams.UpdateSpendPoints(&req, request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "SpendPointsTeam",
		).Errorf(
			"Ошибка списания баллов: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "SpendPointsTeam",
		).Errorf(
			"Ошибка при отправке баллов: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GivesPointsTeam
// @Summary		 GivesPointsTeam
// @Description
// @Tags		 team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /admin/team/{id}/point/give [put]
func (h *HTTPHandler) GivesPointsTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GivesPointsTeam",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.UpdateGivePoints

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "GivesPointsTeam",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.Teams.UpdateGiverPoints(&req, request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "GivesPointsTeam",
		).Errorf(
			"Ошибка списания баллов: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GivesPointsTeam",
		).Errorf(
			"Ошибка при отправке баллов: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetTextTask
// @Summary		 CreateAnswerOnTextTask
// @Description
// @Tags		 task
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Success      200    object     response.GetTextTask
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /task/text [get]
func (h *HTTPHandler) GetTextTask(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetTextTask",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	res, err := h.TextTask.GetTextTask(request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "GetTextTask",
		).Errorf(
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetTextTask",
		).Errorf(
			"Ошибка при отправке баллов: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateAnswerOnTextTaskById
// @Summary		 UpdateAnswerOnTextTaskById
// @Description
// @Tags		 task
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        id path string true "answer ID"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /task/text/answer [put]
func (h *HTTPHandler) UpdateAnswerOnTextTaskById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnTextTaskById",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.UpdateAnswerOnTextTaskByID
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnTextTaskById",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = h.TextTask.UpdateAnswerOnTextTaskById(req, request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnTextTaskById",
		).Errorf(
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetMediaTask
// @Summary		 Get media task
// @Description
// @Tags		 task
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Success      200  {object}     response.GetMediaTask
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /task/media [post]
func (h *HTTPHandler) GetMediaTask(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetMediaTask",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	res, err := h.MediaTask.GetMediaTask(request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "GetMediaTask",
		).Errorf(
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetMediaTask",
		).Errorf(
			"Ошибка при отправке баллов: %s",
			err.Error(),
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetAnswerOnMediaByFilter
// @Summary		 Get Answer On Media By Filter
// @Description
// @Tags		 task
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        status         query     string        false  "string status"
// @Success      200  {object}   response.GetAnswersOnMediaTaskByFilter
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /api/admin/task/media/answer [get]
func (h *HTTPHandler) GetAnswerOnMediaByFilter(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetAnswerOnMediaByFilter",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.GetAnswerOnMediaTaskFilter
	req.Status = r.URL.Query().Get("status")

	res, err := h.MediaTask.GetAnswersOnMediaTaskByFilter(req, request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "GetAnswerOnMediaByFilter",
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetAnswerOnMediaByFilter",
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetAnswerOnMediaTaskById
// @Summary		 Get Answer On Media Task By Id
// @Description
// @Tags		 task
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        id path string true "answer ID"
// @Success      200  {object}  response.GetAnswerOnTextTaskByID
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /api/admin/task/media/answer/{id} [get]
func (h *HTTPHandler) GetAnswerOnMediaTaskById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetAnswerOnMediaTaskById",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.GetAnswerOnMediaTaskById
	req.ID, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.WithField(
			"origin.function", "GetAnswerOnMediaTaskById",
		).Errorf("", err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.MediaTask.GetAnswersOnMediaTaskById(req, request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "GetAnswerOnMediaTaskById",
		).Errorf(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetAnswerOnMediaTaskById",
		)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateAnswerOnMediaTaskById
// @Summary		 Update Answer On Media Task By Id
// @Description
// @Tags		 task
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        id path string true "answer ID"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /task/media/answer/{id} [put]
func (h *HTTPHandler) UpdateAnswerOnMediaTaskById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnMediaTaskById",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.UpdateAnswerOnMediaTask
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnMediaTaskById",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	req.ID, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnMediaTaskById",
		).Errorf("Ошибка чтения id: %s", err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = h.MediaTask.UpdateAnswerOnMediaTask(req, request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnMediaTaskById",
		).Errorf("Ошибка отправки вопроса: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateStatusAnswerOnMediaTask
// @Summary		 Update Status Answer On Media Task
// @Description
// @Tags		 task
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        id path string true "answer ID"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /api/admin/task/media/answer/{id} [put]
func (h *HTTPHandler) UpdateStatusAnswerOnMediaTask(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnMediaTaskById",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.UpdatePointsOnAnswerOnMediaTask
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnMediaTaskById",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	req.Id, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnMediaTaskById",
		).Errorf("Ошибка чтения id: %s", err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = h.MediaTask.UpdatePointsOnAnswerOnMediaTask(request.Session{Value: cookie.Value}, req)
	if err != nil {
		log.WithField(
			"origin.function", "UpdateAnswerOnMediaTaskById",
		).Errorf("Ошибка проставления баллов: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetAllMediaTaskByTeam
// @Summary		 Get All Media Task By Team
// @Description
// @Tags		 task
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /api/task/media/answer [get]
func (h *HTTPHandler) GetAllMediaTaskByTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetAllMediaTaskByTeam",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	res, err := h.MediaTask.GetAllAnswersByTeam(request.Session{Value: cookie.Value})
	if err != nil {
		log.WithField(
			"origin.function", "GetAllMediaTaskByTeam",
		).Errorf("ошибка получения тасок: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetAllMediaTaskByTeam",
		).Errorf("ошибка : %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetMediaTaskByTeamById
// @Summary		 Get Media Task By Team By Id
// @Description
// @Tags		 task
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        id path string true "answer ID"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /api/task/media/answer/{id} [get]
func (h *HTTPHandler) GetMediaTaskByTeamById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetMediaTaskByTeamById",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	var req request.GetAnswerByTeamByID
	req.Id, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.WithField(
			"origin.function", "GetMediaTaskByTeamById",
		).Errorf("ошибка чтения url: %s", err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.MediaTask.GetAnswersByTeamById(request.Session{Value: cookie.Value}, req)
	if err != nil {
		log.WithField(
			"origin.function", "GetMediaTaskByTeamById",
		).Errorf("ошибка получения тасок: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetMediaTaskByTeamById",
		).Errorf("ошибка : %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetAllMasterClass
// @Summary		 GetAllMasterClass
// @Description  возвращает все мастеркласы, которые еще не начались
// @Tags		 sec
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Success      200  object     response.GetSecByFilter
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /sec [get]
func (h *HTTPHandler) GetAllMasterClass(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetAllMasterClass",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	filter := request.NewGetSecByFilter()
	if filter.Bind(r) != nil {
		log.WithField(
			"origin.function", "GetAllMasterClass",
		).Errorf(
			"ошибка запроса %s", err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.Secs.GetSECByFilter(*filter, request.Session{Value: cookie.Value})
	if err != nil {

		log.WithField(
			"origin.function", "GetAllMasterClass",
		).Errorf("%s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetAllMediaTaskByTeam",
		).Errorf("ошибка : %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetMasterClassById
// @Summary		 GetMasterClassById
// @Description
// @Tags		 sec
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        id path string true "master class ID"
// @Success      200   object    response.GetSecById
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /sec/{id} [get]
func (h *HTTPHandler) GetMasterClassById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetAllMasterClass",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	filter := request.NewGetSecByIdFilter()
	if err = filter.Bind(r); err != nil {
		log.WithField(
			"origin.function", "GetAllMasterClass",
		).Errorf(
			"ошибка запроса %s", err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.Secs.GetSECById(*filter, request.Session{Value: cookie.Value})
	if err != nil {

		log.WithField(
			"origin.function", "GetAllMasterClass",
		).Errorf("%s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetAllMediaTaskByTeam",
		).Errorf("ошибка : %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HTTPHandler) GetMasterClassByTeam(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetAllMasterClass",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	filter := request.NewGetSecByTeamId()
	if err = filter.Bind(r); err != nil {
		log.WithField(
			"origin.function", "GetAllMasterClass",
		).Errorf(
			"ошибка запроса %s", err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.Secs.GetSECByTeamId(*filter, request.Session{Value: cookie.Value})
	if err != nil {

		log.WithField(
			"origin.function", "GetAllMasterClass",
		).Errorf("%s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetAllMediaTaskByTeam",
		).Errorf("ошибка : %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HTTPHandler) GetMasterClassAdminById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetMasterClassAdminById",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	filter := request.NewGetSecAdminByIdFilter()
	if err = filter.Bind(r); err != nil {
		log.WithField(
			"origin.function", "GetMasterClassAdminById",
		).Errorf(
			"ошибка запроса %s", err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.Secs.GetSECAdminById(*filter, request.Session{Value: cookie.Value})
	if err != nil {

		log.WithField(
			"origin.function", "GetMasterClassAdminById",
		).Errorf("%s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetMasterClassAdminById",
		).Errorf("ошибка : %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HTTPHandler) GetAllAdminMasterClass(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "GetAllAdminMasterClass",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	filter := request.NewGetSecAdminByFilter()
	if filter.Bind(r) != nil {
		log.WithField(
			"origin.function", "GetAllAdminMasterClass",
		).Errorf(
			"ошибка запроса %s", err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res, err := h.Secs.GetSECAdminByFilter(*filter, request.Session{Value: cookie.Value})
	if err != nil {

		log.WithField(
			"origin.function", "GetAllAdminMasterClass",
		).Errorf("%s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.WithField(
			"origin.function", "GetAllAdminMasterClass",
		).Errorf("ошибка : %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// CreateRegisterOnMasterClass
// @Summary		 CreateRegisterOnMasterClass
// @Description
// @Tags		 sec
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        id path string true "master class ID"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /sec/{id} [put]
func (h *HTTPHandler) CreateRegisterOnMasterClass(w http.ResponseWriter, r *http.Request) {

}

// DeleteRegisterOnMasterClass
// @Summary		 DeleteRegisterOnMasterClass
// @Description
// @Tags		 sec
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        id path string true "master class ID"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /sec/{id} [delete]
func (h *HTTPHandler) DeleteRegisterOnMasterClass(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.WithField(
			"origin.function", "CreateRegisterOnMasterClass",
		).Errorf(
			"Cookie 'access-token' не найден: %s",
			err.Error(),
		)
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	filter := request.NewDeleteRegisterOnSecFilter()
	if err = filter.Bind(r); err != nil {
		log.WithField(
			"origin.function", "CreateRegisterOnMasterClass",
		).Errorf(
			"ошибка запроса %s", err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = h.Secs.DeleteRegisterOnSEC(*filter, request.Session{Value: cookie.Value})
	if err != nil {

		log.WithField(
			"origin.function", "CreateRegisterOnMasterClass",
		).Errorf("%s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
