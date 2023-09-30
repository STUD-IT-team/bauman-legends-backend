package request

type ChangeProfile struct {
	Name        string `json:"name"`
	Group       string `json:"group"`
	Password    string `json:"password"`
	Telegram    string `json:"telegram"`
	VK          string `json:"vk"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
