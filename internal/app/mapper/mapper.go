package mapper

import (
	"fmt"
	"strconv"
	"time"

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
	member := *MakeMembersResponse(dom.Members)
	if member == nil {
		member = make([]response.Member, 0)
	}
	return &response.GetTeam{
		ID:     strconv.Itoa(dom.ID),
		Name:   dom.Name,
		Points: dom.Points,
		Captain: response.Member{
			Id:    strconv.Itoa(dom.Captain.ID),
			Name:  dom.Captain.Name,
			Grope: dom.Captain.Group,
			Email: dom.Captain.Email,
		},
		Members: member,
	}
}

func MakeMembersResponse(dom []domain.Member) *[]response.Member {
	var members []response.Member
	for _, member := range dom {
		mem := response.Member{
			Id:    strconv.Itoa(member.ID),
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
	member := *MakeMembersResponse(dom.Members)
	if member == nil {
		member = make([]response.Member, 0)
	}
	return &response.GetTeamByID{
		ID:     dom.ID,
		Name:   dom.Name,
		Points: dom.Points,
		Captain: response.Member{
			Id:    strconv.Itoa(dom.Captain.ID),
			Name:  dom.Captain.Name,
			Grope: dom.Captain.Group,
			Email: dom.Captain.Email,
		},
		Members: member,
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

func MakeGetMediaTaskResponse(t domain.MediaTask) *response.GetMediaTask {
	return &response.GetMediaTask{
		Id:          t.ID,
		Title:       t.Title,
		Description: fmt.Sprintf("<video src=\"%s\" width=\"500\" height=\"281\" controls></video>", t.VideoUrl),
		Points:      t.Points,
		VideoUrl:    t.VideoUrl,
	}
}

func ParseUpdateAnswerOnMediaTask(req request.UpdateAnswerOnMediaTask) *domain.MediaAnswer {
	return &domain.MediaAnswer{
		Id: req.ID,
	}
}

func MakeGetAnswerOnMediaTask(d domain.MediaTask) *response.GetAnswerOnTextTaskByID {
	return &response.GetAnswerOnTextTaskByID{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		Points:      d.Points,
		PhotoUrl:    d.Answer.PhotoUrl,
		Comment:     d.Answer.Comment,
		Status:      d.Answer.Status,
		TeamId:      d.Answer.TeamId,
	}
}

func MakeGetAnswerOnMediaTaskByFilter(dom []domain.MediaTask) *response.GetAnswersOnMediaTaskByFilter {
	var answers []response.AnswerMediaTask
	for _, d := range dom {
		answer := response.AnswerMediaTask{
			Id:          d.ID,
			Title:       d.Title,
			Description: d.Description,
			Status:      d.Answer.Status,
			PhotoUrl:    d.Answer.PhotoUrl,
		}
		answers = append(answers, answer)
	}
	return &response.GetAnswersOnMediaTaskByFilter{Answers: answers}
}

func MakeGetAllAnswerByTeam(dom []domain.MediaTask) *response.GetAllAnswerByTeam {
	var answers []response.AnswerByTeam
	for _, d := range dom {
		answer := response.AnswerByTeam{
			Id:          d.ID,
			Title:       d.Title,
			Description: d.Description,
			AnswerId:    d.Answer.Id,
			Points:      d.Points,
			Status:      d.Answer.Status,
			Comment:     d.Answer.Comment,
		}
		answers = append(answers, answer)
	}

	return &response.GetAllAnswerByTeam{Answers: answers}
}

func MakeGetAnswersByTeamById(dom domain.MediaTask) *response.GetAnswerByTeamByID {
	return &response.GetAnswerByTeamByID{
		Id:          dom.ID,
		Title:       dom.Title,
		Description: dom.Description,
		Points:      dom.Points,
		Status:      dom.Answer.Status,
		Comment:     dom.Answer.Comment,
		VideoUrl:    dom.VideoUrl,
		PhotoUrl:    dom.Answer.PhotoUrl,
	}
}

func MakeGetUsersByFilter(dom []domain.Member) *response.GetUsersByFilter {
	var users []response.UserByFilter
	for _, d := range dom {
		user := response.UserByFilter{
			Id:          d.ID,
			Name:        d.Name,
			Email:       d.Email,
			Group:       d.Group,
			Telegram:    d.Telegram,
			VK:          d.VK,
			PhoneNumber: d.PhoneNumber,
			TeamName:    d.TeamName,
			Role:        d.Role,
		}

		users = append(users, user)
	}

	return &response.GetUsersByFilter{
		Users: users,
	}
}

func MakeGetUserById(d domain.Member) *response.GetUserById {
	return &response.GetUserById{
		Id:          d.ID,
		Name:        d.Name,
		Email:       d.Email,
		Group:       d.Group,
		Telegram:    d.Telegram,
		VK:          d.VK,
		PhoneNumber: d.PhoneNumber,
		Team:        d.TeamName,
		Role:        d.Role,
	}
}

func MakeGetSECByFilter(dom []domain.Sec) *response.GetSecByFilter {
	secs := make([]response.SECByFilter, 0)
	var sec response.SECByFilter
	masterClasses := make([]response.MasterClassByFilter, 0)
	for _, d := range dom {
		if sec.Id == d.Id {
			masterClasses = append(masterClasses,
				response.MasterClassByFilter{
					Id:        d.MasterClassId,
					StartedAt: d.StartedAt.Add(time.Hour * -3).Format(time.RFC3339),
					EndedAt:   d.EndedAt.Add(time.Hour * -3).Format(time.RFC3339),
					Capacity:  d.Capacity,
					FreePlace: d.Capacity - d.Busy,
				})
		} else {
			sec.MasterClasses = masterClasses
			secs = append(secs, sec)

			masterClasses = make([]response.MasterClassByFilter, 0)

			masterClasses = append(masterClasses,
				response.MasterClassByFilter{
					Id:        d.MasterClassId,
					StartedAt: d.StartedAt.Add(time.Hour * -3).Format(time.RFC3339),
					EndedAt:   d.EndedAt.Add(time.Hour * -3).Format(time.RFC3339),
					Capacity:  d.Capacity,
					FreePlace: d.Capacity - d.Busy,
				})

			sec = response.SECByFilter{
				Id:          d.Id,
				Name:        d.Name,
				Description: d.Description,
				FIO:         d.FIO,
				Phone:       d.Phone,
				Telegram:    d.Telegram,
			}

		}

	}

	sec.MasterClasses = masterClasses
	secs = append(secs, sec)

	return &response.GetSecByFilter{
		SECs: secs[1:],
	}
}

func MakeGetSECById(dom []domain.Sec) *response.GetSecById {
	if len(dom) == 0 {
		return &response.GetSecById{}
	}
	var sec response.GetSecById
	masterClasses := make([]response.MasterClassById, 0)

	for _, d := range dom {
		masterClasses = append(masterClasses,
			response.MasterClassById{
				Id:        d.MasterClassId,
				StartedAt: d.StartedAt.Add(time.Hour * -3).Format(time.RFC3339),
				EndedAt:   d.EndedAt.Add(time.Hour * -3).Format(time.RFC3339),
				Capacity:  d.Capacity,
				FreePlace: d.Capacity - d.Busy,
			})
	}

	sec = response.GetSecById{
		Id:            dom[0].Id,
		Name:          dom[0].Name,
		Description:   dom[0].Description,
		FIO:           dom[0].FIO,
		Phone:         dom[0].Phone,
		Telegram:      dom[0].Telegram,
		PhotoUrl:      dom[0].PhotoUrl,
		MasterClasses: masterClasses,
	}

	return &sec
}

func MakeGetSECByTeamId(dom []domain.Sec) *response.GetSecByTeamId {
	secs := make([]response.SECByTeamId, 0)
	for _, d := range dom {
		sec := response.SECByTeamId{
			Id:            d.Id,
			Name:          d.Name,
			Description:   d.Description,
			FIO:           d.FIO,
			Phone:         d.Phone,
			Telegram:      d.Telegram,
			PhotoUrl:      d.PhotoUrl,
			MasterClassId: d.MasterClassId,
			StartedAt:     d.StartedAt.Add(time.Hour * -3).Format(time.RFC3339),
			EndedAt:       d.EndedAt.Add(time.Hour * -3).Format(time.RFC3339),
			Capacity:      d.Capacity,
			FreePlace:     d.Capacity - d.Busy,
		}
		secs = append(secs, sec)
	}

	return &response.GetSecByTeamId{
		Secs: secs,
	}
}

func MakeGetSECAdminById(dom []domain.Sec) *response.GetSecAdminById {
	if len(dom) == 0 {
		return &response.GetSecAdminById{}
	}
	var sec response.GetSecAdminById
	masterClasses := make([]response.MasterClassAdminById, 0)

	for _, d := range dom {
		masterClasses = append(masterClasses,
			response.MasterClassAdminById{
				Id:        d.MasterClassId,
				StartedAt: d.StartedAt.Add(time.Hour * -3).Format(time.RFC3339),
				EndedAt:   d.EndedAt.Add(time.Hour * -3).Format(time.RFC3339),
				Capacity:  d.Capacity,
				FreePlace: d.Capacity - d.Busy,
			})
	}

	sec = response.GetSecAdminById{
		Id:            dom[0].Id,
		Name:          dom[0].Name,
		Description:   dom[0].Description,
		FIO:           dom[0].FIO,
		Phone:         dom[0].Phone,
		Telegram:      dom[0].Telegram,
		PhotoUrl:      dom[0].PhotoUrl,
		MasterClasses: masterClasses,
	}

	return &sec
}

func MakeGetSECAdminByFilter(dom []domain.Sec) *response.GetSecAdminByFilter {
	secs := make([]response.SECAdminByFilter, 0)
	var sec response.SECAdminByFilter
	masterClasses := make([]response.MasterClassAdminByFilter, 0)
	for _, d := range dom {
		if sec.Id == d.Id {
			masterClasses = append(masterClasses,
				response.MasterClassAdminByFilter{
					Id:        d.MasterClassId,
					StartedAt: d.StartedAt.Add(time.Hour * -3).Format(time.RFC3339),
					EndedAt:   d.EndedAt.Add(time.Hour * -3).Format(time.RFC3339),
					Capacity:  d.Capacity,
					FreePlace: d.Capacity - d.Busy,
				})
		} else {
			sec.MasterClasses = masterClasses
			secs = append(secs, sec)

			masterClasses = make([]response.MasterClassAdminByFilter, 0)

			masterClasses = append(masterClasses,
				response.MasterClassAdminByFilter{
					Id:        d.MasterClassId,
					StartedAt: d.StartedAt.Add(time.Hour * -3).Format(time.RFC3339),
					EndedAt:   d.EndedAt.Add(time.Hour * -3).Format(time.RFC3339),
					Capacity:  d.Capacity,
					FreePlace: d.Capacity - d.Busy,
				})

			sec = response.SECAdminByFilter{
				Id:          d.Id,
				Name:        d.Name,
				Description: d.Description,
				FIO:         d.FIO,
				Phone:       d.Phone,
				Telegram:    d.Telegram,
			}

		}

	}

	sec.MasterClasses = masterClasses
	secs = append(secs, sec)

	return &response.GetSecAdminByFilter{
		SECs: secs[1:],
	}
}

func MakeGetMasterClassById(d domain.Sec) *response.GetMasterClassByID {
	return &response.GetMasterClassByID{
		Id:            d.Id,
		Name:          d.Name,
		Description:   d.Description,
		FIO:           d.FIO,
		Phone:         d.Phone,
		Telegram:      d.Telegram,
		PhotoUrl:      d.PhotoUrl,
		MasterClassId: d.MasterClassId,
		StartedAt:     d.StartedAt.Add(time.Hour * -3).Format(time.RFC3339),
		EndedAt:       d.EndedAt.Add(time.Hour * -3).Format(time.RFC3339),
		Capacity:      d.Capacity,
		FreePlace:     d.Capacity - d.Busy,
	}
}
