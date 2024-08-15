package response

type GetAnswerOnTextTaskByID struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Points      int    `json:"points"`
	Answer      []byte `json:"answer"`
	Comment     string `json:"comment"`
	Status      string `json:"status"`
	TeamId      int    `json:"team_id"`
}
