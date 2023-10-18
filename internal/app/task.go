package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	grpc2 "github.com/STUD-IT-team/bauman-legends-backend/internal/ports/grpc"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"math/rand"
	"strconv"
	"time"
)

type TaskService struct {
	storage repository.TaskStorage
	auth    grpc2.AuthClient
}

func NewTaskService(conn grpc.ClientConnInterface, r repository.TaskStorage) *TaskService {
	return &TaskService{
		storage: r,
		auth:    grpc2.NewAuthClient(conn),
	}
}

func (s *TaskService) GetTaskTypes(req *request.GetTaskTypes) (response.GetTaskTypes, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.AccessToken})
	if err != nil {
		return response.GetTaskTypes{}, fmt.Errorf("can't auth.Check on GetTaskTypes: %w", err)
	}

	if !res.Valid {
		return response.GetTaskTypes{}, errors.New("valid check error")
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.AccessToken})
	if err != nil {
		return response.GetTaskTypes{}, fmt.Errorf("can't auth.GetProfile on GetTaskTypes: %w", err)
	}
	//log.Infof("**************************************************%v", profile)
	taskTypes, err := s.storage.GetTaskTypes(profile.TeamID)
	if err != nil {
		return response.GetTaskTypes{}, fmt.Errorf("can't storage.GetTaskTypes on GetTaskTypes: %w", err)
	}

	//for _, taskType := range taskTypes {
	//	//if taskType.ID == 2 {
	//	//	busyNoc, err := s.storage.GetBusyNocPlaceses()
	//	//	if err != nil {
	//	//		return response.GetTaskTypes{}, fmt.Errorf("can't storage.GetBusyNocPlaceses on GetTaskTypes: %w", err)
	//	//	}
	//	//	if busyNoc < 6 {
	//	//		taskType.IsActive = true
	//	//	}
	//	//}
	//}
	return mapper.MakeTaskTypesResponse(taskTypes), nil
}

func (s *TaskService) TakeTask(req *request.TakeTask) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.AccessToken})
	if err != nil {
		return fmt.Errorf("can't auth.Check on TakeTask: %w", err)
	}
	taskTypeId, err := strconv.Atoi(req.TaskTypeId)
	if err != nil {
		return err
	}
	if !res.Valid {
		return errors.New("valid check error")
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.AccessToken})
	if err != nil {
		return fmt.Errorf("can't auth.GetProfile on TakeTask: %w", err)
	}

	taskAmount, err := s.storage.GetTaskAmount(taskTypeId)
	if err != nil {
		return fmt.Errorf("can't storage.GetTaskAmount on TakeTask: %w", err)
	}

	teamTaskAmount, err := s.storage.GetTeamTaskAmount(taskTypeId)
	if err != nil {
		return fmt.Errorf("can't take teamTaskAmount : %w", err)
	}
	if teamTaskAmount >= taskAmount { //проверка что есть еще задачи такого типа
		return errors.New("team already completed all tasks this type")
	}

	//if req.TaskTypeId == 2 {
	//	busyNoc, err := s.storage.GetBusyNocPlaceses()
	//	if err != nil {
	//		return fmt.Errorf("can't storage.GetBusyNocPlaceses on GetTaskTypes: %w", err)
	//	}
	//	if busyNoc >= 6 { //проверка что ноц свободны если taskTypeID = 2
	//		return errors.New("all Noc placec are busy")
	//	}
	//
	//}
	availableTasks, err := s.storage.GetAvailableTaskID(profile.TeamID, taskTypeId)
	if err != nil {
		return fmt.Errorf("can't take availableTasks : %w", err)
	}
	log.Infof("++++++++++++%v", availableTasks)
	//rand.Seed()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(availableTasks), func(i, j int) { availableTasks[i], availableTasks[j] = availableTasks[j], availableTasks[i] })
	taskID := availableTasks[0] //выбрать одну случайную задача такого типа
	task, err := s.storage.GetTask(taskID)
	if err != nil {
		return fmt.Errorf("can't storage.GetTask : %w", err)
	}
	err = s.storage.SetTaskToTeam(taskID, task.TypeID, profile.TeamID) // insert into team_task
	if err != nil {
		return fmt.Errorf("can't storage.SetTaskToTeam : %w", err)
	}
	return nil
}

func (s *TaskService) GetTask(req *request.GetTask) (*response.GetTask, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.AccessToken})
	if err != nil {
		return nil, fmt.Errorf("can't auth.Check on GetTask: %w", err)
	}
	if !res.Valid {
		return nil, errors.New("valid check error")
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.AccessToken})
	if err != nil {
		return nil, fmt.Errorf("can't auth.GetProfile on GetTask: %w", err)
	}

	checkTask, err := s.storage.CheckActiveTaskExist(profile.TeamID)
	if err != nil {
		return nil, fmt.Errorf("can't storage.CheckActiveTaskExist on GetTask: %w", err)
	}
	if !checkTask {
		return nil, errors.New("this team does not have any active task")
	}
	taskID, err := s.storage.GetActiveTaskID(profile.TeamID)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetActiveTaskID on GetTask: %w", err)
	}
	//answerText, answerImageUrl, result
	task, err := s.storage.GetTask(taskID)

	task.StartedTime, err = s.storage.GetTaskStartedTime(taskID, profile.TeamID)
	if err != nil {
		return nil, err
	}
	task.TypeName, err = s.storage.GetTaskTypeName(taskID)

	return mapper.MakeGetTaskResponse(task), nil
}

func (s *TaskService) Answer(req *request.Answer) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.AccessToken})
	if err != nil {
		return fmt.Errorf("can't auth.Check on Answer: %w", err)
	}
	if !res.Valid {
		return errors.New("valid check error")
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.AccessToken})
	if err != nil {
		return fmt.Errorf("can't auth.GetProfile on Answer: %w", err)
	}
	taskID, err := s.storage.GetActiveTaskID(profile.TeamID)
	if err != nil {
		return err
	}
	//task, err := s.storage.GetTask(taskID)

	if req.Text == nil {
		err := s.storage.SetAnswerText(*req.Text, profile.TeamID, taskID)
		if err != nil {
			return err
		}
	} else {
		err := s.storage.SetAnswerImageBase64(*req.ImageUrl, profile.TeamID, taskID)
		if err != nil {
			return err
		}
	}
	return nil

}

func (s *TaskService) UploadPhoto(req request.UploadPhoto) error {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.AccessToken})
	if err != nil {
		return fmt.Errorf("can't auth.Check on UploadPhoto: %w", err)
	}
	if !res.Valid {
		return errors.New("valid check error")
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.AccessToken})
	if err != nil {
		return fmt.Errorf("can't auth.GetProfile on Answer: %w", err)
	}
	taskID, err := s.storage.GetActiveTaskID(profile.TeamID)
	if err != nil {
		return fmt.Errorf("can't s.storage.GetActiveTaskID on UploadPhoto :%w", err)
	}
	err = s.storage.SetAnswerImageBase64(req.Base64Photo, profile.TeamID, taskID)
	if err != nil {
		return fmt.Errorf("can't s.storage.SetAnswerImageBase64 on UploadPhoto :%w", err)
	}
	return nil
}

func (s *TaskService) GetAnswers(req *request.GetAnswers) (*response.GetAnswers, error) {
	res, err := s.auth.Check(context.Background(), &grpc2.CheckRequest{AccessToken: req.AccessToken})
	if err != nil {
		return nil, fmt.Errorf("can't auth.Check on GetAnswers: %w", err)
	}
	if !res.Valid {
		return nil, errors.New("valid check error")
	}

	profile, err := s.auth.GetProfile(context.Background(), &grpc2.GetProfileRequest{AccessToken: req.AccessToken})
	if err != nil {
		return nil, fmt.Errorf("can't auth.GetProfile on GetAnswers: %w", err)
	}

	dbAnswers, err := s.storage.GetAnswers(profile.TeamID)

	var httpAnswers response.GetAnswers

	for _, answ := range dbAnswers {
		var hAnsw response.Answer

		hAnsw.AnswerId = answ.ID
		hAnsw.TeamId = answ.TeamID
		//hAnsw.TeamTitle
		hAnsw.TaskId = answ.TaskID
		//hAnsw.TaskTitle = answ.
		httpAnswers.Answers = append(httpAnswers.Answers, hAnsw)

	}

	return &httpAnswers, nil

}
