package domain

type Member struct {
	ID          int    `json:"id"`
	RoleId      int    `json:"roleId"`
	Name        string `json:"name"`
	Group       string `json:"group"`
	Email       string `json:"email"`
	Telegram    string `json:"telegram"`
	VK          string `json:"VK"`
	TeamName    string `json:"team_name"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
}
