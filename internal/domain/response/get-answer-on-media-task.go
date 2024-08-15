package response

type GetAnswerOnMediaTask struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Points      int    `json:"points"`
	Answer      []byte `json:"answer"`
}
