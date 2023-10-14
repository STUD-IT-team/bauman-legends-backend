package response

type UserProfile struct {
	Name        string `json:"name" db:"name"`
	Group       string `json:"group" db:"group"`
	Telegram    string `json:"telegram" db:"telegram"`
	VK          string `json:"vk" db:"vk"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Email       string `json:"email" db:"email"`
	TeamID      string `json:"team_id" db:"team_id"`
}
