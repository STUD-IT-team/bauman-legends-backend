package request

type UpdateSpendPoints struct {
	TeamId      int `json:"team_id"`
	DeltaPoints int `json:"spend_points"`
}
