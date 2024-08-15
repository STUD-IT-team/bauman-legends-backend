package response

type GetAnswersOnMediaTaskByFilter struct {
	Answers []AnswerMediaTask `json:"answers"`
}

type AnswerMediaTask struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Answer      []byte `json:"answer"`
}
