package request

type UpdateAnswerOnMediaTask struct {
	ID       int    `json:"id"`
	Answer   string `json:"answer"`
	TypeData string `json:"type_data"`
}
