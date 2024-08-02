package request

type UpdateGivePoints struct {
	Session     string `json:"session"`
	TeamId      int    `json:"team_id"`
	DeltaPoints int    `json:"spend_points"`
}
