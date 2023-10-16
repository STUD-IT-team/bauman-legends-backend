package domain

import "database/sql"

// Team
//
// Таблица команд
type Team struct {
	TeamId  string `db:"id"`
	Title   string `db:"title"`
	Points  int    `db:"points"`
	Members []Member
}

type Member struct {
	Id   string        `db:"id"`
	Name string        `db:"name"`
	Role sql.NullInt64 `db:"role"`
}
