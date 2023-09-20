package domain

// Team
//
// Таблица команд
type Team struct {
	ID    ID     `db:"id"`
	Title string `db:"title"`
}
