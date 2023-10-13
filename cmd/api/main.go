package main

import (
	"fmt"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/settings"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/ports/handlers"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os"
)

func main() {
	settings.LogSetup()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", os.Getenv("AUTH_DN"), os.Getenv("AUTH_PORT")),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.WithField(
			"origin.function", "main",
		).Fatalf(
			"Невозможно установить соединение с сервисом регистрации и авторизации: %s",
			err.Error(),
		)
	}

	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.WithField(
				"origin.function", "main",
			).Errorf(
				"Не удалось закрыть соединение с сервером регистрации и авторизации: %s",
				err.Error(),
			)
		}
	}(conn)

	api := app.NewApi(conn)
	handler := handlers.NewHTTPHandler(api)

	r := chi.NewRouter()
	r.Post("/api/user", handler.Register)
	r.Post("/api/user/auth", handler.Login)
	r.Delete("/api/user/session", handler.Logout)
	r.Get("/api/user", handler.GetProfile)
	r.Put("/api/user", handler.ChangeProfile)

	log.WithField(
		"origin.function", "main",
	).Info(
		"Сервер запущен",
	)

	if err = http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("API_DN"), os.Getenv("API_PORT")), r); err != nil {
		log.WithField(
			"origin.function", "main",
		).Infof(
			"Сервер завершил работу: %s",
			err.Error(),
		)
	}
}
