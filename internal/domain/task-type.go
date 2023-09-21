package domain

// TaskType
//
// Таблица типов заданий
type TaskType struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}
