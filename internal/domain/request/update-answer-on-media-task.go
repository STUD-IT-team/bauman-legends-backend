package request

type UpdateAnswerOnMediaTask struct {
	ID     int    `json:"id"`
	Answer []byte `json:"answer"`
}
