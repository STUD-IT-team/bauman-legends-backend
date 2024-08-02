package domain

type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Members []Member
	Captain Member
	Points  int `json:"points"`
}

type Member struct {
	ID     int    `json:"id"`
	RoleId int    `json:"roleId"`
	Name   string `json:"name"`
	Group  string `json:"group"`
	Email  string `json:"email"`
}
