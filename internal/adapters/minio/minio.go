package minio

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
	"io/ioutil"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/storage"
)

type MinioStorage struct {
	client *minio.Client
}

func NewMinioStorage(minioURL string, accessKey string, secretKey string, ssl bool) (storage.ObjectStorage, error) {
	client, err := minio.New(minioURL, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: ssl,
	})
	if err != nil {
		return nil, err
	}

	return &MinioStorage{client: client}, nil
}

func (s *MinioStorage) GetObject(object domain.Object) (domain.Object, error) {
	data, err := s.client.GetObject(
		context.Background(),
		object.BucketName,
		object.ObjectName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		log.WithField(
			"", "GetObject",
		).Errorf("GetObject Error: %s", err.Error())
		return domain.Object{}, err
	}
	defer data.Close()

	object.Data, err = ioutil.ReadAll(data)
	// buf := new(bytes.Buffer)
	// _, err = buf.ReadFrom(data)
	if err != nil {
		log.WithField(
			"origin.function", "GetObject",
		).Errorf("ReadAll Error: %s", err.Error())
		return domain.Object{}, err
	}

	// object.Data = buf.Bytes()

	return object, nil
}

func (s *MinioStorage) PutObject(object domain.Object) (key string, err error) {
	reader := bytes.NewReader(object.Data)
	info, err := s.client.PutObject(
		context.Background(),
		object.BucketName,
		object.ObjectName,
		reader,
		object.Size,
		minio.PutObjectOptions{},
	)
	if err != nil {
		log.WithField(
			"origin.function", "PutObject",
		).Errorf("PutObject Error: %s", err.Error())
		return "", err
	}

	return info.Key, nil
}

func (s *MinioStorage) DeleteObject(object domain.Object) (err error) {
	err = s.client.RemoveObject(
		context.Background(),
		object.BucketName,
		object.ObjectName,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		log.WithField(
			"origin.function", "DeleteObject",
		).Errorf("DeleteObject Error: %s", err.Error())
		return err
	}

	return nil
}
