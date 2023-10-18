package domain

import "time"

// Answer
//
// Таблица ответов на задания
type Answer struct {
	ID                string    `db:"id"`
	TaskID            string    `db:"task_id"`
	TeamID            string    `db:"team_id"`
	StartTime         time.Time `db:"start_time"`
	AdditionalPoints  int       `db:"additional_points"`
	AnswerText        string    `db:"answer_text"`
	AnswerImageBase64 string    `db:"answerImageBase64"`
	Result            bool      `db:"result"`
}
