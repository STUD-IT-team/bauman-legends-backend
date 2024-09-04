package request

type UpdatePointsOnAnswerOnMediaTask struct {
	Id      int    `json:"id"`
	Answer  bool   `json:"answer"`
	Comment string `json:"comment"`
}
