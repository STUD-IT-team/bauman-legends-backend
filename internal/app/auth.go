package main

import (
	"fmt"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/repository"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/session"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
)

func main() {
	logSetup()

	sessionCache := session.NewSessionCache()
	repo := repository.NewRepository(os.Getenv("DATA_SOURCE"))
	service := grpc.NewAuth(sessionCache, repo)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
	if err != nil {
		log.WithField(
			"origin.function", "main",
		).Fatalf("Невозможно установить tcp-соединение: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	grpc.RegisterAuthServer(grpcServer, service)

	log.WithField("origin.function", "main").Info("Сервер запущен")

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

func logSetup() {
	formatter := new(log.TextFormatter)
	formatter.TimestampFormat = "02.01.2006 15:04:05"
	formatter.FullTimestamp = true
	formatter.DisableLevelTruncation = true
	log.SetFormatter(formatter)
}
