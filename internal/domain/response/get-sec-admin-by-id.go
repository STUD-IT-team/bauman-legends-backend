package response

type GetSecAdminById struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	FIO         string   `json:"fio"`
	Phone       string   `json:"phone"`
	Telegram    string   `json:"telegram"`
	Times       []string `json:"times"`
	Capacity    []int    `json:"capacity"`
	BusyPlace   []int    `json:"busy_place"`
	PhotoUrl    string   `json:"photo_url"`
}
