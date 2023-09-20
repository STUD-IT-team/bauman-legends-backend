package domain

import "time"

// TeamSecret
//
// Таблица связи секретных заданий и команд
type TeamSecret struct {
	ID        ID         `db:"id"`
	SecretID  ID         `db:"secret_id"`
	TeamID    ID         `db:"team_id"`
	StartTime time.Time  `db:"start_time"`
	EndTime   *time.Time `db:"end_time"`
}
