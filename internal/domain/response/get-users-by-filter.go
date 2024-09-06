package response

type GetUsersByFilter struct {
	Users []UserByFilter `json:"users"`
}

type UserByFilter struct {
	Id          int    `json:"id"`
	Role        string `json:"role"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Group       string `json:"group"`
	Telegram    string `json:"telegram"`
	VK          string `json:"VK"`
	PhoneNumber string `json:"phone_number"`
	TeamName    string `json:"team"`
}
