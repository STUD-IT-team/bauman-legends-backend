package domain

import "time"

// Task
//
// Таблица заданий
type Task struct {
	ID           ID        `db:"id"`
	Title        string    `db:"title"`
	Description  string    `db:"description"`
	TimeLimit    time.Time `db:"time_limit"`
	DifficultyID int       `db:"difficulty_id"`
	TypeID       int       `db:"type_id"`
}
