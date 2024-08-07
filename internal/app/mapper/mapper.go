package mapper

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
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
		Id:          res.ID,
		Name:        res.Name,
		Group:       res.Group,
		Email:       res.Email,
		Telegram:    res.Telegram,
		Vk:          res.VK,
		PhoneNumber: res.PhoneNumber,
		TeamID:      res.TeamID,
		IsAdmin:     res.IsAdmin,
	}
}

func MakeProfileResponse(res *grpc2.GetProfileResponse) *response.UserProfile {
	return &response.UserProfile{
		ID:          res.Id,
		Name:        res.Name,
		Group:       res.Group,
		Email:       res.Email,
		Telegram:    res.Telegram,
		VK:          res.Vk,
		PhoneNumber: res.PhoneNumber,
		TeamID:      res.TeamID,
		IsAdmin:     res.IsAdmin,
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

func MakeChangePasswordRequest(req *grpc2.ChangePasswordRequest) *request.ChangePassword {
	return &request.ChangePassword{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
}

func MakeGetTeamResponse(dom domain.Team) *response.GetTeam {
	return &response.GetTeam{
		ID:     dom.ID,
		Name:   dom.Name,
		Points: dom.Points,
		Captain: response.Member{
			Id:    dom.Captain.ID,
			Name:  dom.Captain.Name,
			Grope: dom.Captain.Group,
			Email: dom.Captain.Email,
		},
		Members: *MakeMembersResponse(dom.Members),
	}
}

func MakeMembersResponse(dom []domain.Member) *[]response.Member {
	var members []response.Member
	for _, member := range dom {
		mem := response.Member{
			Id:    member.ID,
			Name:  member.Name,
			Grope: member.Group,
			Email: member.Email,
		}
		members = append(members, mem)
	}
	return &members
}

func MakeGetTeamsResponse(dom []domain.Team) *response.GetTeamsByFilter {
	var teams []response.GetTeam
	for _, team := range dom {
		t := *MakeGetTeamResponse(team)
		teams = append(teams, t)
	}
	return &response.GetTeamsByFilter{Teams: teams}
}

func MakeGetTeamByIdResponse(dom domain.Team) *response.GetTeamByID {
	return &response.GetTeamByID{
		ID:     dom.ID,
		Name:   dom.Name,
		Points: dom.Points,
		Captain: response.Member{
			Id:    dom.Captain.ID,
			Name:  dom.Captain.Name,
			Grope: dom.Captain.Group,
			Email: dom.Captain.Email,
		},
		Members: *MakeMembersResponse(dom.Members),
	}
}
