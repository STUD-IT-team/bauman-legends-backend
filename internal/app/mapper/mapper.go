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

func MakeGetTextTaskResponse(dom domain.TextTask) *response.GetTextTask {
	return &response.GetTextTask{
		ID:          dom.ID,
		Title:       dom.Title,
		Description: dom.Description,
		Points:      dom.Points,
	}
}

func MakeGetAnswerOnTextTask(dom domain.TextTask) *response.GetAnswerOnTextTaskByID {
	return &response.GetAnswerOnTextTaskByID{
		ID:          dom.ID,
		Title:       dom.Title,
		Description: dom.Description,
		Points:      dom.Points,
	}
}

func ParseGetTextTaskRequest(req *request.GetTextTask) *domain.TextTask {
	return &domain.TextTask{}
}

func ParseUpdateAnswerOnTextTask(req request.UpdateAnswerOnTextTaskByID) *domain.TextTask {
	return &domain.TextTask{
		ID:     req.ID,
		Answer: req.Answer,
	}
}

func MakeGetMediaTaskResponse(t domain.MediaTask, o domain.Object) *response.GetMediaTask {
	return &response.GetMediaTask{
		Id:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Points:      t.Points,
		Video:       o.Data,
	}
}

func ParseUpdateAnswerOnMediaTask(req request.UpdateAnswerOnMediaTask) *domain.MediaTask {
	return &domain.MediaTask{
		ID: req.ID,
	}
}

func MakeGetAnswerOnMediaTask(d domain.MediaTask) *response.GetAnswerOnTextTaskByID {
	return &response.GetAnswerOnTextTaskByID{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		Points:      d.Points,
		Answer:      d.Answer,
		Comment:     d.Comment,
		Status:      d.Status,
		TeamId:      d.TeamId,
	}
}

func MakeGetAnswerOnMediaTaskByFilter(dom []domain.MediaTask) *response.GetAnswersOnMediaTaskByFilter {
	var answers []response.AnswerMediaTask
	for _, d := range dom {
		answer := response.AnswerMediaTask{
			Id:          d.ID,
			Title:       d.Title,
			Description: d.Description,
			Status:      d.Status,
			Answer:      d.Answer,
		}
		answers = append(answers, answer)
	}
	return &response.GetAnswersOnMediaTaskByFilter{Answers: answers}
}
