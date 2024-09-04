package main

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/adapters/minio"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
)

func main() {
	storage, err := minio.NewMinioStorage(
		"localHost:9000",
		"user",
		"password",
		false)

	if err != nil {
		panic(err)
	}

	pathDir := "/home/verendaya/Documents/Киношки/Три блатных аккорда/"

	files, err := os.ReadDir(pathDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileName := file.Name()
		fmt.Println(fileName)

		data, err := ioutil.ReadFile(pathDir + fileName)
		if err != nil {
			log.Println("Error reading file:", err)
		}

		objName := uuid.New().String()

		obj := domain.Object{
			BucketName: "video-task",
			ObjectName: objName,
			TypeData:   "video",
			Data:       data,
			Size:       int64(len(data)),
		}

		key, err := storage.PutObject(obj)
		if err != nil {
			log.Println("Error:", err)
		}

		fmt.Println(key) // 4be6c96c-c424-4f14-8dad-257d352a62b7*/
	}

}
