package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/adapters/cache"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/adapters/postgres"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/settings"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
)

func main() {
	settings.LogSetup()

	sessionCache := cache.NewSessionCache()

	pgConnString := fmt.Sprintf(os.Getenv("DATA_SOURCE"), os.Getenv("DB_DN"))
	repo, err := postgres.NewUserAuthStorage(pgConnString)
	if err != nil {
		log.WithField(
			"origin.function", "main",
		).Fatalf(
			"Невозможно установить соединение с базой данных: %s",
			err.Error(),
		)
	}
	log.Info("Соединение с бд установленно")

	service := app.NewAuth(sessionCache, repo)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("AUTH_DN"), os.Getenv("AUTH_PORT")))
	if err != nil {
		log.WithField(
			"origin.function", "main",
		).Fatalf("Невозможно установить tcp-соединение: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	grpc2.RegisterAuthServer(grpcServer, service)

	log.WithField("origin.function", "main").Info("Сервер запущен на ", os.Getenv("AUTH_DN"), ":", os.Getenv("AUTH_PORT"))

	err = grpcServer.Serve(lis)
	if err != nil {
		log.WithField(
			"origin.function", "main",
		).Fatalf("Сервер остановлен: %s", err.Error())
	}

	log.WithField(
		"origin.function", "main",
	).Info("Сервер завершил работу")
}
