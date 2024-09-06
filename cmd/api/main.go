package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/STUD-IT-team/bauman-legends-backend/cmd/api/docs"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/adapters/minio"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/adapters/postgres"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/settings"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/ports/handlers"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/storage"
)

// @title           Backend Bauman Legends
// @version         1.0
// @description     This is backend server for bauman legends 2024.

// @host      localhost:3000
// @BasePath  /api/

// @securityDefinitions.basic  BasicAuth

func main() {
	settings.LogSetup()
	grpcURl := fmt.Sprintf("%s:%s", os.Getenv("AUTH_DN"), os.Getenv("AUTH_PORT"))
	conn, err := grpc.NewClient(
		grpcURl,
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
	pgString := os.Getenv("DB_SOURCE")

	teamStorage, err := postgres.NewTeamStorage(pgString)
	if err != nil {
		log.WithField(
			"origin.function", "main",
		).Fatalf(
			"Невозможно установить соединение с базой данных: %s",
			err.Error(),
		)
	}
	log.Info("NewTeamStorage connected to db")

	textTaskStorage, err := postgres.NewTextTaskStorage(pgString)
	if err != nil {
		log.WithField(
			"origin.function", "main",
		).Fatalf(
			"Невозможно установить соединение с базой данных: %s",
			err.Error(),
		)
	}
	log.Info("NewTextTaskStorage connected to db")

	mediaTaskStorage, err := postgres.NewMediaTaskStorage(pgString)
	if err != nil {
		log.WithField(
			"origin.function", "main",
		).Fatalf("Невозможно установить соединение с базой данных: %s",
			err.Error(),
		)
	}

	userStorage, err := postgres.NewUserStorage(pgString)
	if err != nil {
		log.WithField(
			"origin.function", "main",
		).Fatalf("Невозможно установить соединение с базой данных: %s",
			err.Error(),
		)
	}

	objectStorage, err := minio.NewMinioStorage("localhost:9000", "user", "password", false)
	if err != nil {
		log.WithField(
			"origin.function", "main",
		).Fatalf("Невозможно установить соединение с обьектным хранилищем: %s",
			err.Error(),
		)
	}

	storage := storage.NewStorage(teamStorage, textTaskStorage, mediaTaskStorage, objectStorage, userStorage)

	teams := app.NewTeamService(conn, storage)
	textTask := app.NewTextTaskService(conn, storage)
	mediaTask := app.NewMediaTaskService(conn, storage)
	users := app.NewUserService(conn, storage)

	api := app.NewApi(conn)
	handler := handlers.NewHTTPHandler(api, teams, textTask, mediaTask, users)

	r := chi.NewRouter()

	r.Post("/api/user/auth/register", handler.CreateUser)
	r.Post("/api/user/auth/login", handler.Login)
	r.Delete("/api/user/auth/logout", handler.Logout)
	r.Get("/api/user", handler.GetProfile)
	r.Put("/api/user", handler.UpdateProfile)

	r.Get("/api/admin/user", handler.GetUsersByFilter)
	r.Get("/api/admin/user/{id}", handler.GetUserById)

	r.Post("/api/team", handler.CreateTeam)
	r.Delete("/api/team", handler.DeleteTeam)
	r.Put("/api/team", handler.UpdateTeam)
	r.Get("/api/team", handler.GetTeam)

	r.Post("/api/team/member", handler.AddUserInTeam)
	r.Delete("/api/team/member", handler.DeleteUserFromTeam)

	r.Put("/api/admin/team/{id}/point/spend", handler.SpendPointsTeam)
	r.Put("/api/admin/team/{id}/point/give", handler.GivesPointsTeam)
	r.Get("/api/admin/team", handler.GetTeamsByFilter)
	r.Get("/api/admin/team/{id}", handler.GetTeamById)

	r.Get("/api/task/text", handler.GetTextTask)
	r.Put("/api/task/text/answer", handler.UpdateAnswerOnTextTaskById)

	r.Get("/api/task/media", handler.GetMediaTask)
	r.Put("/api/task/media/answer/{id}", handler.UpdateAnswerOnMediaTaskById)
	r.Get("/api/task/media/answer", handler.GetAllMediaTaskByTeam)
	r.Get("/api/task/media/answer/{id}", handler.GetMediaTaskByTeamById)

	r.Get("/api/admin/task/media/answer", handler.GetAnswerOnMediaByFilter)
	r.Get("/api/admin/task/media/answer/{id}", handler.GetAnswerOnMediaTaskById)
	r.Put("/api/admin/task/media/answer/{id}", handler.UpdateStatusAnswerOnMediaTask)

	r.Get("/api/docs/*", httpSwagger.WrapHandler)
	log.WithField(
		"origin.function", "main",
	).Info(
		"Сервер запущен",
	)

	s := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

	/*if err = http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("API_DN"), os.Getenv("API_PORT")), r); err != nil {
		log.WithField(
			"origin.function", "main",
		).Infof(
			"Сервер завершил работу: %s",
			err.Error(),
		)
	}*/
}
