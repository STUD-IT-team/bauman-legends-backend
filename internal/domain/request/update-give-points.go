package request

type UpdateGivePoints struct {
	TeamId      int `json:"team_id"`
	DeltaPoints int `json:"spend_points"`
}
