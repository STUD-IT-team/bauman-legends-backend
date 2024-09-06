package request

type UpdateAnswerOnTextTaskByID struct {
	ID     int    `json:"id"`
	Answer string `json:"answer"`
}
