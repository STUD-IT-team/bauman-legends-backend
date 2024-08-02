package consts

import "errors"

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
)

const (
	UserRole    = 1
	SellerRole  = 2
	CaptainRole = 3
	AdminRole   = 4
)
