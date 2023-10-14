package request

// RegisterTeam
// Структура запроса регистрации команды
type RegisterTeam struct {
	Session  string `json:"session"`
	TeamName string `json:"name"`
}

type ChangeTeam struct {
	Session  string `json:"session"`
	TeamName string `json:"name"`
}

type GetTeam struct {
	Session string `json:"session"`
}

type DeleteTeam struct {
	Session string `json:"session"`
}
type InviteToTeam struct {
	Session string `json:"session"`
	UserID  string `json:"id"`
}
type DeleteFromTeam struct {
	Session string `json:"session"`
	UserID  string `json:"id"`
}
type UpdateMember struct {
	Session string `json:"session"`
	UserID  string `json:"id"`
	RoleID  int    `json:"roleid"`
}
