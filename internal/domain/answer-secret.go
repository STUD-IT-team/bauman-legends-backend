package domain

// AnswerSecret
//
// Таблица ответов на секретные задания
type AnswerSecret struct {
	ID           ID     `db:"id"`
	SecretID     ID     `db:"secret_id"`
	AnswerTypeID int    `db:"answer_type_id"`
	Data         string `db:"data"`
}
