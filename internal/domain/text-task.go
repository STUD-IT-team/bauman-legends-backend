package domain

type TextTask struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Answer      string `json:"answer"`
	Points      int    `json:"points"`
	Status      string `json:"status"`
	TeamId      int    `json:"team_id"`
}
