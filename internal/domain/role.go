package domain

// Role
//
// Таблица ролей в команде
type Role struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}
