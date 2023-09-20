package domain

// AnswerType
//
// Таблица типов ответов
type AnswerType struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}
