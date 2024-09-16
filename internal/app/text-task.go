package app

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"strconv"
	"strings"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/storage"
)

type TextTaskService struct {
	storage storage.Storage
	auth    grpc2.AuthClient
}

func NewTextTaskService(conn grpc.ClientConnInterface, s storage.Storage) *TextTaskService {
	return &TextTaskService{
		storage: s,
		auth:    grpc2.NewAuthClient(conn),
	}
}

func (s *TextTaskService) GetTextTask(session request.Session) (response.GetTextTask, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: session.Value})
	if err != nil {
		return response.GetTextTask{}, err
	}

	if !res.Valid {
		return response.GetTextTask{}, errors.New("invalid token")
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: session.Value})
	if err != nil {
		return response.GetTextTask{}, err
	}

	teamId, err := strconv.Atoi(profile.TeamID)
	if err != nil {
		return response.GetTextTask{}, err
	}

	status, err := s.storage.GetStatusLastTextTask(teamId)
	if err != nil {
		return response.GetTextTask{}, err
	}

	exist, err := s.storage.CheckDayNewTask(teamId) //true если ответили сегодня
	if err != nil {
		return response.GetTextTask{}, err
	}

	var task domain.TextTask

	if !status {
		task, err = s.storage.GetLastTextTask(teamId)
		if err != nil {
			return response.GetTextTask{}, err
		}
	} else {
		if exist {
			return response.GetTextTask{}, consts.LockedError
		}

		task, err = s.storage.GetNewTextTask(teamId)
		if err != nil {
			return response.GetTextTask{}, err
		}

		task.TeamId = teamId

		err = s.storage.CreateAnswerOnTextTask(task)
		if err != nil {
			return response.GetTextTask{}, err
		}
	}

	return *mapper.MakeGetTextTaskResponse(task), nil
}

func (s *TextTaskService) UpdateAnswerOnTextTaskById(req request.UpdateAnswerOnTextTaskByID, session request.Session) (string, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: session.Value})
	if err != nil {
		return "", err
	}

	if !res.Valid {
		return "", errors.New("invalid token")
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: session.Value})
	if err != nil {
		return "", err
	}

	teamId, err := strconv.Atoi(profile.TeamID)
	if err != nil {
		return "", err
	}

	answer := *mapper.ParseUpdateAnswerOnTextTask(req)
	answer.TeamId = teamId

	userId, err := strconv.Atoi(profile.Id)
	if err != nil {
		return "", err
	}

	isCaptain, err := s.storage.CheckUserRoleById(userId, consts.CaptainRole)
	if err != nil {
		return "", err
	}

	if !isCaptain {
		return "", consts.ForbiddenError
	}

	task, err := s.storage.GetLastTextTask(teamId)
	if err != nil {
		return "", err
	}

	if strings.ToLower(task.Answer) == strings.ToLower(answer.Answer) {
		answer.Status = true
		answer.Points = task.Points
	} else {
		answer.Status = false
	}

	err = s.storage.UpdateAnswerOnTextTask(answer)
	if err != nil {
		return "", err
	}

	if answer.Status {
		return "true", nil
	} else {
		return "false", consts.TeaPodCode
	}
}
