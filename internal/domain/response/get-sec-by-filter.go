package response

type GetSecByFilter struct {
	SECs []SECByFilter `json:"secs"`
}

type SECByFilter struct {
	Id            int                   `json:"id"`
	Name          string                `json:"name"`
	Description   string                `json:"description"`
	FIO           string                `json:"fio"`
	Phone         string                `json:"phone"`
	Telegram      string                `json:"telegram"`
	MasterClasses []MasterClassByFilter `json:"master_classes"`
}

type MasterClassByFilter struct {
	Id        int    `json:"id"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
	Capacity  int    `json:"capacity"`
	FreePlace int    `json:"free_place"`
}
