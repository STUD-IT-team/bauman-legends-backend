package mapper

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
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

func MakeGrpcResponseProfile(res *response.UserProfile) *grpc2.GetProfileResponse {
	return &grpc2.GetProfileResponse{
		Name:        res.Name,
		Group:       res.Group,
		Email:       res.Email,
		Telegram:    res.Telegram,
		Vk:          res.VK,
		PhoneNumber: res.PhoneNumber,
	}
}

func MakeProfileResponse(res *grpc2.GetProfileResponse) *response.UserProfile {
	return &response.UserProfile{
		Name:        res.Name,
		Group:       res.Group,
		Email:       res.Email,
		Telegram:    res.Telegram,
		VK:          res.Vk,
		PhoneNumber: res.PhoneNumber,
	}
}

func MakeChangeProfileRequest(req *grpc2.ChangeProfileRequest) *request.ChangeProfile {
	return &request.ChangeProfile{
		Name:        req.Name,
		Group:       req.Group,
		Password:    req.Password,
		Telegram:    req.Telegram,
		VK:          req.Vk,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
}

func MakeRegisterTeamRequest(req *grpc2.RegisterTeamRequest) *request.RegisterTeam {
	return &request.RegisterTeam{
		TeamName: req.TeamName,
	}
}

func MakeGrpcRequestRegisterTeam(req *request.RegisterTeam) *grpc2.RegisterTeamRequest {
	return &grpc2.RegisterTeamRequest{
		AccessToken: "",
		TeamName:    req.TeamName,
	}
}
