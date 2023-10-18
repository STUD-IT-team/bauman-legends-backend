package domain

import "time"

// Task
//
// Таблица заданий
type Task struct {
	ID           ID     `db:"id"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	TypeID       int    `db:"type_id"`
	MaxPoints    int    `db:"max_points"`
	MinPoints    int    `db:"min_points"`
	AnswerTypeID int    `db:"answer_type_id"`
	TypeName     string
	StartedTime  time.Time
}
