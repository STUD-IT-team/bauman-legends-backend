package domain

import "time"

// TeamTask
//
// Таблица связи заданий и команд
type TeamTask struct {
	ID               ID         `db:"id"`
	TaskID           ID         `db:"task_id"`
	TeamID           ID         `db:"team_id"`
	StartTime        time.Time  `db:"start_time"`
	EndTime          *time.Time `db:"end_time"`
	AdditionalPoints int        `db:"additional_points"`
}
