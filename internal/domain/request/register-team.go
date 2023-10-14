package request

// RegisterTeam
// Структура запроса регистрации команды
type RegisterTeam struct {
	TeamName string `json:"name"`
}
