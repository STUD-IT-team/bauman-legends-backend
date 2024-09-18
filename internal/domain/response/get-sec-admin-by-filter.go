package response

type GetSecAdminByFilter struct {
	SECs []SECAdminByFilter `json:"secs"`
}

type SECAdminByFilter struct {
	Id            int                        `json:"id"`
	Name          string                     `json:"name"`
	Description   string                     `json:"description"`
	FIO           string                     `json:"fio"`
	Phone         string                     `json:"phone"`
	Telegram      string                     `json:"telegram"`
	MasterClasses []MasterClassAdminByFilter `json:"master_classes"`
}

type MasterClassAdminByFilter struct {
	Id        int    `json:"id"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
	Capacity  int    `json:"capacity"`
	FreePlace int    `json:"free_place"`
}
