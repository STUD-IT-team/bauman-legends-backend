package response

import "time"

type GetTask struct {
	Title        string
	Text         string
	TypeId       int
	TypeName     string
	MaxPoints    int
	MinPoints    int
	Duration     time.Duration
	TimeStarted  time.Time
	HintsTaken   []Hints
	answerTypeId int
}

type Hints struct {
	Title         string
	Text          string
	PointsPenalty int
}
