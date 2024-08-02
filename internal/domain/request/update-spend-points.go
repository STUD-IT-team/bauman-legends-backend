package request

type UpdateSpendPoints struct {
	Session     string `json:"session"`
	TeamId      int    `json:"team_id"`
	DeltaPoints int    `json:"spend_points"`
}
