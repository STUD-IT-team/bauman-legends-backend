package domain

// TaskDifficulty
//
// Таблица уровней сложности заданий
type TaskDifficulty struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}
