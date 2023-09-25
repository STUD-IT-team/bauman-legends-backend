package request

// Login
// Структура запроса входа в аккаунт
type Login struct {
	Email         string
	Password      string
	ClientBrowser string
	ClientOS      string
}
