package request

// Register
// Структура запроса регистрации
type Register struct {
	Name          string `json:"name"`
	Group         string `json:"group"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Telegram      string `json:"telegram"`
	VK            string `json:"vk"`
	PhoneNumber   string `json:"phone_number"`
	ClientBrowser string `json:"client_browser"`
	ClientOS      string `json:"client_os"`
}
