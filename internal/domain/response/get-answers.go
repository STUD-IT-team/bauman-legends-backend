package response

import "time"

type GetAnswers struct {
	Answers []Answer `json:"answers"`
}

type Answer struct {
	AnswerId          string    `json:"answerId"`
	TeamId            string    `json:"teamId"`
	TeamTitle         string    `json:"teamTitle"`
	TaskId            string    `json:"taskId"`
	TaskTitle         string    `json:"taskTitle"`
	TaskDescription   string    `json:"taskDescription"`
	TaskTypeId        int       `json:"taskTypeId"`
	TaskTypeName      string    `json:"taskTypeName"`
	AnswerText        string    `json:"answerText"`
	AnswerImageBase64 string    `json:"answerImageBase64"`
	TimeStart         time.Time `json:"timeGotten"`
	AdditionalPoints  int       `json:"additionalPoints"`
	TaskAnswerTypeId  int       `json:"taskAnswerTypeId"`
	IsComfirmed       bool      `json:"isComfirmed"`
}

//answerId, (То что у тебя в team_task записано)
//teamId,
//teamTitle
//taskId,
//taskTitle,
//taskDescription,
//taskTypeId,
//taskTypeName,
//answerText,
//answerImageBase64,
//timeGotten,
//additionalPoints,
//taskAnswerTypeId,
//isComfirmed: Boolean | null
