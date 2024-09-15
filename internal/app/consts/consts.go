package consts

import (
	"errors"
	"time"
)

const (
	LogDateFormat  = "02.01.2006 15:04:05"
	GrpcTimeFormat = "02.01.2006 15:04:05.000"
)

var (
	BadRequestError     = errors.New("bad request")
	UnAuthorizedError   = errors.New("unauthorized")
	ForbiddenError      = errors.New("forbidden")
	NotFoundError       = errors.New("not found")
	ConflictError       = errors.New("conflict")
	LockedError         = errors.New("locked")
	InternalServerError = errors.New("internal server error")
	NoRowsAffectedError = errors.New("no rows affected")
)

const (
	UserRole    = 1
	SellerRole  = 2
	CaptainRole = 3
	AdminRole   = 4
)

const (
	ReviewStatus  = "review"
	EmptyStatus   = "empty"
	CorrectStatus = "correct"
	WrongStatus   = "wrong"
)

const (
	StartTextTask      = "16.09.2024 10:00:00"
	EndTextTask        = "22.09.2024 23.59.59"
	FirstDayMediaTask  = "23.09.2024 00:00:00"
	SecondDayMediaTask = "24.09.2024 00:00:00"
	ThirdDayMediaTask  = "25.09.2024 00:00:00"
)

const (
	FirstDayCoefficient  = 1.5
	SecondDayCoefficient = 1.25
	ThirdDayCoefficient  = 1.0
)

var FirstDayMediaTaskTime = time.Date(2024, time.September, 23, 0, 0, 0, 0, time.UTC)
var SecondDayMediaTaskTime = time.Date(2024, time.September, 24, 0, 0, 0, 0, time.UTC)
var ThirdDayMediaTaskTime = time.Date(2024, time.September, 25, 0, 0, 0, 0, time.UTC)

const (
	MinioUrl          = "http://localhost:9000/"
	VideoTaskBucket   = "video-task"
	PhotoAnswerBucket = "photo-answer"
)

const SecDay = "2024-09-24 "
