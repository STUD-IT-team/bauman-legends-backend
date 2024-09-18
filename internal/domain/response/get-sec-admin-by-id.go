package response

type GetSecAdminById struct {
	Id            int                    `json:"id"`
	Name          string                 `json:"name"`
	Description   string                 `json:"description"`
	FIO           string                 `json:"fio"`
	Phone         string                 `json:"phone"`
	Telegram      string                 `json:"telegram"`
	MasterClasses []MasterClassAdminById `json:"master_classes"`
	PhotoUrl      string                 `json:"photo_url"`
}

type MasterClassAdminById struct {
	Id        int    `json:"id"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
	Capacity  int    `json:"capacity"`
	FreePlace int    `json:"free_place"`
}
