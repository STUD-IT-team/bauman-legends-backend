package request

// Login
// Структура запроса входа в аккаунт
type Login struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	ClientBrowser string `json:"client_browser"`
	ClientOS      string `json:"client_os"`
}
