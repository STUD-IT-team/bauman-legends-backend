package mapper

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
)

// MakeRequestRegister
//
// Преобразование grpc-запроса регистрации в структуру
func MakeRequestRegister(req *grpc2.RegisterRequest) *request.Register {
	return &request.Register{
		Name:          req.Name,
		Group:         req.Group,
		Email:         req.Email,
		Password:      req.Password,
		Telegram:      req.Telegram,
		VK:            req.Vk,
		PhoneNumber:   req.PhoneNumber,
		ClientBrowser: req.ClientBrowser,
		ClientOS:      req.ClientOS,
	}
}

// MakeGrpcRequestRegister
//
// Преобразование запроса регистрации в grpc-структуру
func MakeGrpcRequestRegister(req *request.Register) *grpc2.RegisterRequest {
	return &grpc2.RegisterRequest{
		Name:          req.Name,
		Group:         req.Group,
		Email:         req.Email,
		Password:      req.Password,
		Telegram:      req.Telegram,
		Vk:            req.VK,
		PhoneNumber:   req.PhoneNumber,
		ClientBrowser: req.ClientBrowser,
		ClientOS:      req.ClientOS,
	}
}

// MakeGrpcRequestLogin
//
// Преобразование запроса входа в grpc-структуру
func MakeGrpcRequestLogin(req *request.Login) *grpc2.LoginRequest {
	return &grpc2.LoginRequest{
		Email:         req.Email,
		Password:      req.Password,
		ClientBrowser: req.ClientBrowser,
		ClientOS:      req.ClientOS,
	}
}
