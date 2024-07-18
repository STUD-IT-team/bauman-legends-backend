package request

type CreateTeam struct {
	Name    string `json:"name"`
	Session string `json:"session"`
}
