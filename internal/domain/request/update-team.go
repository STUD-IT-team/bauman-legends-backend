package request

type UpdateTeam struct {
	NewTeamName string `json:"new_team_name"`
	Session     string `json:"session"`
}
