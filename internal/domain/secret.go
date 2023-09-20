package domain

// Secret
//
// Таблица секретных заданий
type Secret struct {
	ID          ID     `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}
