package app

import (
	"context"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/adapters/cache"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/repository"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/session"
	"time"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
)

type Auth struct {
	SessionCache cache.ICache[string, session.Session]
	Repository   repository.IRepository
}

func (s *Auth) Register(ctx context.Context, req *grpc.RegisterRequest) (*grpc.RegisterResponse, error) {
	// TODO:
	// 	После редактирования базы и репозитория - дописать
	return nil, nil
}

func (s *Auth) Login(ctx context.Context, req *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	// TODO:
	// 	После редактирования базы и репозитория - дописать
	return nil, nil
}

func (s *Auth) Logout(ctx context.Context, req *grpc.LogoutRequest) (*grpc.LogoutResponse, error) {
	accessToken := req.GetAccessToken()
	s.SessionCache.Delete(accessToken)

	return &grpc.LogoutResponse{
		Message: "success",
	}, nil
}

func (s *Auth) Check(ctx context.Context, req *grpc.CheckRequest) (*grpc.CheckResponse, error) {
	//select {
	//case <-ctx.Done():
	//	return nil, nil
	//}

	// TODO:
	// 	Как использовать контекст?

	accessToken := req.GetAccessToken()
	record := s.SessionCache.Find(accessToken)

	if record == nil ||
		record.ExpireAt.Before(time.Now()) {
		return &grpc.CheckResponse{
			Valid: false,
		}, nil
	}

	return &grpc.CheckResponse{
		Valid: true,
	}, nil
}
