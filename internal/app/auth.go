package app

import (
	"context"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
)

type GRPCServer struct{}

func (s *GRPCServer) Register(ctx context.Context, req *grpc.RegisterRequest) (*grpc.RegisterResponse, error) {

}

func (s *GRPCServer) Login(ctx context.Context, req *grpc.LoginRequest) (*grpc.LoginResponse, error) {

}

func (s *GRPCServer) Logout(ctx context.Context, req *grpc.LogoutRequest) (*grpc.LogoutResponse, error) {

}
