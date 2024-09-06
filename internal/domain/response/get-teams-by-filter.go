package response

type GetTeamsByFilter struct {
	Teams []GetTeam `json:"teams"`
}
