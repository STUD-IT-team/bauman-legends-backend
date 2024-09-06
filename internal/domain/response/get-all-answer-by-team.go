package response

type GetAllAnswerByTeam struct {
	Answers []AnswerByTeam `json:"answers"`
}

type AnswerByTeam struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AnswerId    int    `json:"answer_id"`
	Points      int    `json:"points"`
	Status      string `json:"status"`
	Comment     string `json:"comment"`
}
