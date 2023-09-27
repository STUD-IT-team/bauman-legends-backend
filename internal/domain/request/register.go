package request

// Register
// Структура запроса регистрации
type Register struct {
	Name          string
	Group         string
	Email         string
	Password      string
	Telegram      string
	VK            string
	PhoneNumber   string
	ClientBrowser string
	ClientOS      string
}
