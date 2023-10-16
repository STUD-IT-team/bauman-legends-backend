package response

// RegisterTeam
// Структура запроса регистрации команды
type RegisterTeam struct {
	TeamID string `json:"id"`
}

type UpdateTeam struct {
	TeamName string `json:"team_name"`
}

type GetTeam struct {
	TeamId  string   `json:"team_id"`
	Title   string   `json:"title"`
	Points  int      `json:"points"`
	Members []Member `json:"members"`
}

type Member struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Role int    `json:"role"`
}
