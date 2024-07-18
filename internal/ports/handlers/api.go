package handlers

import (
	"encoding/json"
	"net/http"
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
	Api *app.Api
	// Teams *app.TeamService
	// Tasks *app.TaskService
}

func NewHTTPHandler(api *app.Api) *HTTPHandler {
	return &HTTPHandler{
		Api: api,
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
// @Success      200  {string}  string    "ok"
// @Failure      400  {string}  string    "bad request"
// @Failure		 403  {string}  string	  "not rights"
// @Failure      401  {string}  string    "not authorized"
// @Failure      500  {string}  string    "internal server error"
// @Router       /admin/user [get]
func (h *HTTPHandler) GetUsersByFilter(w http.ResponseWriter, r *http.Request) {}

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
// @Router       /admin/user [get]
func (h *HTTPHandler) GetUserById(w http.ResponseWriter, r *http.Request) {}

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
// @Param        request.RegisterTeam	body		request.RegisterTeam	true	"Add team"
// @Success      201  {object}  response.RegisterTeam
// @Failure      400  {string}  string    "bad request"
// @Failure 	 401  {string}  string    "not authorized"
// @Failure		 409  {string}  string    "user exists"
// @Failure      500  {string}  string    "internal server error"
// @Router       /team [post]
func (h *HTTPHandler) CreateTeam(w http.ResponseWriter, r *http.Request) {
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

	var req request.CreateTeam
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithField(
			"origin.function", "ChangeProfile",
		).Errorf(
			"Ошибка чтения запроса: %s",
			err.Error(),
		)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	req.Session = cookie.Value
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
func (h *HTTPHandler) GetTeam(w http.ResponseWriter, r *http.Request) {}

// UpdateTeam
// @Summary      UpdateTeam
// @Description  Update info about team (only capitan of team)
// @Tags         team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        request.ChangeTeam	body		request.ChangeTeam	true	"Change team"
// @Success      200  {string}  string    "ok"
// @Failure      400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      500  {string}  string    "internal server error"
// @Router       /team [put]
func (h *HTTPHandler) UpdateTeam(w http.ResponseWriter, r *http.Request) {}

// DeleteTeam
// @Summary      DeleteTeam
// @Description  Delete team if team is empty (only capitan of team)
// @Tags         team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        request.DeleteTeam  body  request.DeleteTeam true  "Delete team"
// @Success      200  {string}  string    "ok"
// @Failure      401  {string}  string    "not authorized"
// @Failure      500  {string}  string    "internal server error"
// @Router       /team [delete]
func (h *HTTPHandler) DeleteTeam(w http.ResponseWriter, r *http.Request) {}

/* func (h *HTTPHandler) Invite(w http.ResponseWriter, r *http.Request) {
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

}*/

// AddUserInTeam
// @Summary		 AddUserInTeam
// @Description  Capitan of team can add user in his team
// @Tags		 team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param 		 request.InviteToTeam  body  request.InviteToTeam  true  "Invite to team"
// @Success      200  {string}  string    "ok"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /team/member [post]
func (h *HTTPHandler) AddUserInTeam(w http.ResponseWriter, r *http.Request) {}

// DeleteUserFromTeam
// @Summary		 DeleteUserFromTeam
// @Description  User can be removed from the team (not captain)
// @Tags		 team
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param 		 request.DeleteFromTeam  body  request.DeleteFromTeam  true  "Delete from team"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      500  {string}  string    "internal server error"
// @Router       /team/member [delete]
func (h *HTTPHandler) DeleteUserFromTeam(w http.ResponseWriter, r *http.Request) {}

/* func (h *HTTPHandler) DeleteMember(w http.ResponseWriter, r *http.Request) {
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

}*/

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
func (h *HTTPHandler) GetTeamsByFilter(w http.ResponseWriter, r *http.Request) {}

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
// @Router       /admin/team [get]
func (h *HTTPHandler) GetTeamById(w http.ResponseWriter, r *http.Request) {}

// SpendPointsTeam
// @Summary		 SpendPointsTeam
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
// @Router       /admin/team/point [post]
func (h *HTTPHandler) SpendPointsTeam(w http.ResponseWriter, r *http.Request) {}

// CreateAnswerOnTextTask
// @Summary		 CreateAnswerOnTextTask
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
// @Router       /task/text [post]
func (h *HTTPHandler) CreateAnswerOnTextTask(w http.ResponseWriter, r *http.Request) {}

// CreateAnswerOnMediaTask
// @Summary		 CreateAnswerOnMediaTask
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
// @Router       /task/media [post]
func (h *HTTPHandler) CreateAnswerOnMediaTask(w http.ResponseWriter, r *http.Request) {}

// GetAnswerOnMediaByFilter
// @Summary		 GetAnswerOnMediaByFilter
// @Description
// @Tags		 task
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Param        status         query     string        false  "string status"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /task/media/answer [get]
func (h *HTTPHandler) GetAnswerOnMediaByFilter(w http.ResponseWriter, r *http.Request) {}

// GetAnswerOnTextTaskById
// @Summary		 GetAnswerOnTextTaskById
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
// @Router       /task/text/answer [get]
func (h *HTTPHandler) GetAnswerOnTextTaskById(w http.ResponseWriter, r *http.Request) {}

// GetAnswerOnMediaTaskById
// @Summary		 GetAnswerOnMediaTaskById
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
// @Router       /task/media/answer [get]
func (h *HTTPHandler) GetAnswerOnMediaTaskById(w http.ResponseWriter, r *http.Request) {}

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
func (h *HTTPHandler) UpdateAnswerOnTextTaskById(w http.ResponseWriter, r *http.Request) {}

// UpdateAnswerOnMediaTaskById
// @Summary		 UpdateAnswerOnMediaTaskById
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
func (h *HTTPHandler) UpdateAnswerOnMediaTaskById(w http.ResponseWriter, r *http.Request) {}

// GivePointsOnTask
// @Summary		 GivePointsOnTask
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
// @Router       /task/text/answer [post]
func (h *HTTPHandler) GivePointsOnTask(w http.ResponseWriter, r *http.Request) {}

// GetAllMasterClass
// @Summary		 GetAllMasterClass
// @Description
// @Tags		 sec
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param 		 Authorization header string true "Authorization"
// @Success      200  {string}  string    "ok"
// @Failure		 400  {string}  string    "bad request"
// @Failure      401  {string}  string    "not authorized"
// @Failure      403  {string}  string    "not rights"
// @Failure      500  {string}  string    "internal server error"
// @Router       /sec [get]
func (h *HTTPHandler) GetAllMasterClass(w http.ResponseWriter, r *http.Request) {}

// GetMasterClassById
// @Summary		 GetMasterClassById
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
// @Router       /sec/ [get]
func (h *HTTPHandler) GetMasterClassById(w http.ResponseWriter, r *http.Request) {}

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
// @Router       /sec [post]
func (h *HTTPHandler) CreateRegisterOnMasterClass(w http.ResponseWriter, r *http.Request) {}

// UpdateRegisterOnMasterClass
// @Summary		 UpdateRegisterOnMasterClass
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
// @Router       /sec [put]
func (h *HTTPHandler) UpdateRegisterOnMasterClass(w http.ResponseWriter, r *http.Request) {}
