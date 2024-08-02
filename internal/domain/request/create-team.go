package request

type CreateTeam struct {
	TeamName string `json:"team_name"`
	Session  string `json:"session"`
}
