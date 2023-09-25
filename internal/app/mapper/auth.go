package mapper

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
)

// TODO:
// 	Зачем нужны мапперы?

// MakeRequestLogin
//
// Преобразование grpc-запроса входа в аккаунт в структуру
func MakeRequestLogin(grpcRequest *grpc.LoginRequest) *request.Login {
	return &request.Login{
		Email:         grpcRequest.Email,
		Password:      grpcRequest.Password,
		ClientBrowser: grpcRequest.ClientBrowser,
		ClientOS:      grpcRequest.ClientOS,
	}
}

// MakeRequestRegister
//
// Преобразование grpc-запроса регистрации в структуру
func MakeRequestRegister(grpcRequest *grpc.RegisterRequest) *request.Register {
	return &request.Register{
		Name:          grpcRequest.Name,
		Group:         grpcRequest.Group,
		Email:         grpcRequest.Email,
		Password:      grpcRequest.Password,
		Telegram:      grpcRequest.Telegram,
		VK:            grpcRequest.Vk,
		ClientBrowser: grpcRequest.ClientBrowser,
		ClientOS:      grpcRequest.ClientOS,
	}
}

// MakeRequestLogout
//
// Преобразование grpc-запроса выхода из аккаунта в структуру
func MakeRequestLogout(grpcRequest *grpc.LogoutRequest) *request.Logout {
	return &request.Logout{
		AccessToken: grpcRequest.AccessToken,
	}
}

// MakeResponseLogin
//
// Преобразование структуры ответа на запрос входа в grpc-ответ
func MakeResponseLogin(response *response.Login) *grpc.LoginResponse {
	return &grpc.LoginResponse{
		AccessToken: response.AccessToken,
	}
}

// MakeResponseRegister
//
// Преобразование структуры ответа на запрос регистрации в grpc-ответ
func MakeResponseRegister(response *response.Register) *grpc.RegisterResponse {
	return &grpc.RegisterResponse{Message: response.Message}
}

// MakeResponseLogout
//
// Преобразование структуры ответа на запрос выхода в grpc-ответ
func MakeResponseLogout(response *response.Logout) *grpc.LogoutResponse {
	return &grpc.LogoutResponse{Message: response.Message}
}
