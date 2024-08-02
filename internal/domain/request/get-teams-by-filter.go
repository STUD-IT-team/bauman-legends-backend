package request

type GetTeamsByFilter struct {
	Session      string `json:"session"`
	MembersCount int
}
