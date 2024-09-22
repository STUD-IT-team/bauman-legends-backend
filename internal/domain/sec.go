package domain

import "time"

type Sec struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	FIO           string    `json:"fio"`
	Phone         string    `json:"phone"`
	Telegram      string    `json:"telegram"`
	MasterClassId int       `json:"master_class_id"`
	StartedAt     time.Time `json:"started_at"`
	EndedAt       time.Time `json:"ended_at"`
	Capacity      int       `json:"capacity"`
	Busy          int       `json:"busy"`
	PhotoUrl      string    `json:"photo_url"`
}
