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

func MakeHttpResponseGetTeam(team *domain.Team) *response.GetTeam {
	var memb []response.Member
	for i := range team.Members {
		role := 0
		if team.Members[i].Role.Valid {
			role = int(team.Members[i].Role.Int64)
		}
		memb = append(memb, response.Member{
			Id:   team.Members[i].Id,
			Name: team.Members[i].Name,
			Role: role,
		})
	}
	return &response.GetTeam{
		TeamId:  team.TeamId,
		Title:   team.Title,
		Points:  team.Points,
		Members: memb,
	}
}

func MakeTaskTypesResponse(in domain.TaskTypes) *response.GetTaskTypes {
	out := make([]response.TaskType, 0, len(in))

	for _, taskType := range in {
		out = append(out, response.TaskType{
			Name:     taskType.Title,
			ID:       taskType.ID,
			IsActive: taskType.IsActive,
		})
	}

	return &response.GetTaskTypes{
		TaskTypes: out,
	}
}

func MakeGetTaskResponse(in domain.Task) *response.GetTask {
	return &response.GetTask{
		Title:        in.Title,
		Text:         in.Description,
		TypeId:       in.TypeID,
		TypeName:     in.TypeName,
		MaxPoints:    in.MaxPoints,
		MinPoints:    in.MinPoints,
		TimeStarted:  in.StartedTime,
		AnswerTypeId: in.AnswerTypeID,
	}
}
