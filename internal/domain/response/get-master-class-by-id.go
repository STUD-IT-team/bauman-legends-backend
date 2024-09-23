package response

type GetMasterClassByID struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	FIO           string `json:"fio"`
	Phone         string `json:"phone"`
	Telegram      string `json:"telegram"`
	MasterClassId int    `json:"master_class_id"`
	StartedAt     string `json:"started_at"`
	EndedAt       string `json:"ended_at"`
	Capacity      int    `json:"capacity"`
	FreePlace     int    `json:"free_place"`
	PhotoUrl      string `json:"photo_url"`
}
