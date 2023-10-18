package response

import "time"

type GetTask struct {
	Title        string    `json:"title"`
	Text         string    `json:"text"`
	TypeId       int       `json:"typeId"`
	TypeName     string    `json:"typeName"`
	MaxPoints    int       `json:"maxPoints"`
	MinPoints    int       `json:"minPoints"`
	TimeStarted  time.Time `json:"timeStarted"`
	AnswerTypeId int       `json:"answerTypeId"`
}
