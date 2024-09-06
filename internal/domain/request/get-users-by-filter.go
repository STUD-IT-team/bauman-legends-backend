package request

type GetUsersByFilter struct {
	WithTeam    bool `json:"with_team"`
	CountInTeam int  `json:"count_in_team"`
}
