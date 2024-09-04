package app

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"strconv"
	"time"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/storage"
)

type MediaTaskService struct {
	storage storage.Storage
	auth    grpc2.AuthClient
}

func NewMediaTaskService(conn *grpc.ClientConn, s storage.Storage) *MediaTaskService {
	return &MediaTaskService{
		storage: s,
		auth:    grpc2.NewAuthClient(conn),
	}
}

func (s *MediaTaskService) GetMediaTask(session request.Session) (response.GetMediaTask, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: session.Value})
	if err != nil {
		return response.GetMediaTask{}, err
	}

	if !res.Valid {
		return response.GetMediaTask{}, errors.New("invalid token")
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: session.Value})
	if err != nil {
		return response.GetMediaTask{}, err
	}

	teamId, err := strconv.Atoi(profile.TeamID)
	if err != nil {
		return response.GetMediaTask{}, err
	}

	status, err := s.storage.GetStatusLastMediaTask(teamId)
	if err != nil {
		return response.GetMediaTask{}, err
	}

	var task domain.MediaTask

	if status == consts.CorrectStatus {
		task, err = s.storage.GetNewMediaTask(teamId)
		if err != nil {
			return response.GetMediaTask{}, err
		}

		_, err = s.storage.CreateAnswerOnMediaTask(teamId, task.ID)
		if err != nil {
			return response.GetMediaTask{}, err
		}
	} else {
		task, err = s.storage.GetLastMediaTask(teamId)
		if err != nil {
			return response.GetMediaTask{}, err
		}
	}

	task.VideoUrl = consts.MinioUrl + consts.VideoTaskBucket + "/" + task.VideoKey

	return *mapper.MakeGetMediaTaskResponse(task), nil
}

func (s *MediaTaskService) UpdateAnswerOnMediaTask(req request.UpdateAnswerOnMediaTask, session request.Session) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: session.Value})
	if err != nil {
		return err
	}

	if !res.Valid {
		return errors.New("invalid token")
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: session.Value})
	if err != nil {
		return err
	}

	teamId, err := strconv.Atoi(profile.TeamID)
	if err != nil {
		return err
	}

	exist, err := s.storage.CheckAnswerOnMediaTaskById(req.ID, teamId)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("invalid id")
	}

	answer := *mapper.ParseUpdateAnswerOnMediaTask(req)
	answer.TeamId = teamId

	photo := domain.Object{
		BucketName: consts.PhotoAnswerBucket,
		ObjectName: uuid.New().String(),
		TypeData:   req.TypeData,
		Size:       int64(len(req.Answer)),
		Data:       req.Answer,
	}

	answer.PhotoKey, err = s.storage.PutObject(photo)
	if err != nil {
		return err
	}

	err = s.storage.UpdateAnswerOnMediaTask(req.ID, answer)
	if err != nil {
		return err
	}

	return nil
}

func (s *MediaTaskService) GetAnswersOnMediaTaskByFilter(req request.GetAnswerOnMediaTaskFilter, session request.Session) (response.GetAnswersOnMediaTaskByFilter, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: session.Value})
	if err != nil {
		return response.GetAnswersOnMediaTaskByFilter{}, err
	}

	if !res.Valid {
		return response.GetAnswersOnMediaTaskByFilter{}, errors.New("invalid token")
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return response.GetAnswersOnMediaTaskByFilter{}, err
	}

	isAdmin, err := s.storage.CheckUserRoleById(userId, consts.AdminRole)
	if err != nil {
		return response.GetAnswersOnMediaTaskByFilter{}, err
	}

	if !isAdmin {
		return response.GetAnswersOnMediaTaskByFilter{}, errors.New("user is not admin")
	}

	var answers []domain.MediaTask

	if req.Status == "" {
		answers, err = s.storage.GetAllAnswerOnMediaTasks()
		if err != nil {
			return response.GetAnswersOnMediaTaskByFilter{}, err
		}
	} else {
		answers, err = s.storage.GetAnswersOnMediaTasksByFilter(req.Status)
		if err != nil {
			return response.GetAnswersOnMediaTaskByFilter{}, err
		}
	}

	for i, _ := range answers {
		answers[i].Answer.PhotoUrl = consts.MinioUrl + consts.PhotoAnswerBucket + "/" + answers[i].Answer.PhotoKey
	}

	return *mapper.MakeGetAnswerOnMediaTaskByFilter(answers), nil
}

func (s *MediaTaskService) GetAnswersOnMediaTaskById(req request.GetAnswerOnMediaTaskById, session request.Session) (response.GetAnswerOnTextTaskByID, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: session.Value})
	if err != nil {
		return response.GetAnswerOnTextTaskByID{}, err
	}

	if !res.Valid {
		return response.GetAnswerOnTextTaskByID{}, errors.New("invalid token")
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return response.GetAnswerOnTextTaskByID{}, err
	}

	isAdmin, err := s.storage.CheckUserRoleById(userId, consts.AdminRole)
	if err != nil {
		return response.GetAnswerOnTextTaskByID{}, err
	}

	if !isAdmin {
		return response.GetAnswerOnTextTaskByID{}, errors.New("user is not admin")
	}

	answer, err := s.storage.GetAnswerOnMediaTaskById(req.ID)
	if err != nil {
		return response.GetAnswerOnTextTaskByID{}, err
	}

	answer.Answer.PhotoUrl = consts.MinioUrl + consts.PhotoAnswerBucket + "/" + answer.Answer.PhotoKey

	return *mapper.MakeGetAnswerOnMediaTask(answer), nil
}

func (s *MediaTaskService) UpdatePointsOnAnswerOnMediaTask(session request.Session, req request.UpdatePointsOnAnswerOnMediaTask) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: session.Value})
	if err != nil {
		return err
	}

	if !res.Valid {
		return errors.New("invalid token")
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return err
	}

	isAdmin, err := s.storage.CheckUserRoleById(userId, consts.AdminRole)
	if err != nil {
		return err
	}

	if !isAdmin {
		return errors.New("user is not an admin")
	}

	var status string
	taskId := req.Id

	points := 0
	var point float32

	if req.Answer {
		status = consts.CorrectStatus
		points, err = s.storage.GetPointsOnMediaTask(taskId)
		if err != nil {
			return err
		}

		if consts.FirstDayMediaTaskTime.After(time.Now()) && consts.SecondDayMediaTaskTime.Before(time.Now()) {
			point *= consts.FirstDayCoefficient
			points = int(point)
		}
		if consts.SecondDayMediaTaskTime.After(time.Now()) && consts.ThirdDayMediaTaskTime.Before(time.Now()) {
			point *= consts.SecondDayCoefficient
			points = int(point)
		}
		if consts.ThirdDayMediaTaskTime.After(time.Now()) {
			point *= consts.ThirdDayCoefficient
			points = int(point)
		}

	} else {
		status = consts.WrongStatus
	}

	err = s.storage.UpdatePointsOnMediaTask(status, taskId, points, req.Comment)
	if err != nil {
		return err
	}

	return nil
}

func (s *MediaTaskService) GetAllAnswersByTeam(session request.Session) (response.GetAllAnswerByTeam, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: session.Value})
	if err != nil {
		return response.GetAllAnswerByTeam{}, err
	}

	if !res.Valid {
		return response.GetAllAnswerByTeam{}, errors.New("invalid token")
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return response.GetAllAnswerByTeam{}, err
	}

	team, err := s.storage.GetTeamByUserId(userId)
	if err != nil {
		return response.GetAllAnswerByTeam{}, err
	}

	answers, err := s.storage.GetAllMediaTaskByTeam(team.ID)
	if err != nil {
		return response.GetAllAnswerByTeam{}, err
	}

	return *mapper.MakeGetAllAnswerByTeam(answers), nil
}

func (s *MediaTaskService) GetAnswersByTeamById(session request.Session, req request.GetAnswerByTeamByID) (response.GetAnswerByTeamByID, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: session.Value})
	if err != nil {
		return response.GetAnswerByTeamByID{}, err
	}

	if !res.Valid {
		return response.GetAnswerByTeamByID{}, errors.New("invalid token")
	}

	userId, err := strconv.Atoi(res.UserID)
	if err != nil {
		return response.GetAnswerByTeamByID{}, err
	}

	team, err := s.storage.GetTeamByUserId(userId)
	if err != nil {
		return response.GetAnswerByTeamByID{}, err
	}

	answer, err := s.storage.GetMediaTaskByTeamById(team.ID, req.Id)
	if err != nil {
		return response.GetAnswerByTeamByID{}, err
	}

	answer.VideoUrl = consts.MinioUrl + consts.VideoTaskBucket + "/" + answer.VideoKey
	answer.Answer.PhotoUrl = consts.MinioUrl + consts.PhotoAnswerBucket + "/" + answer.Answer.PhotoKey

	return *mapper.MakeGetAnswersByTeamById(answer), nil
}
