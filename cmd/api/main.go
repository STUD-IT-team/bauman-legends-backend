package main

import (
	"fmt"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/settings"
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

	r := chi.NewRouter()
	r.Post("/user", api.Register)
	r.Post("/user/auth", api.Login)
	r.Delete("/user/session", api.Logout)
	r.Get("/user", api.GetProfile)
	r.Put("/user", api.ChangeProfile)

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
