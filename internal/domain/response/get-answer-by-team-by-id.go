package response

type GetAnswerByTeamByID struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Points      int    `json:"points"`
	Status      string `json:"status"`
	Comment     string `json:"comment"`
	VideoUrl    string `json:"video_url"`
	PhotoUrl    string `json:"photo_url"`
}
