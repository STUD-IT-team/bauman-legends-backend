package request

type GetTeamById struct {
	TeamId  string `json:"team_id"`
	Session string `json:"session"`
}
