package domain

// Answer
//
// Таблица ответов на задания
type Answer struct {
	ID           ID     `db:"id"`
	TaskID       ID     `db:"task_id"`
	AnswerTypeID int    `db:"answer_type_id"`
	Data         string `db:"data"`
}
