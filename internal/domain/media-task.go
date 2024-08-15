package domain

type MediaTask struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Points      int    `json:"points"`
	VideoId     int    `json:"video_id"`
	VideoKey    string `json:"video_key"`
	Video       []byte `json:"video"`
	PhotoId     int    `json:"photo_id"`
	PhotoKey    string `json:"photo_key"`
	Answer      []byte `json:"answer"`
	TeamId      int    `json:"team_id"`
	AnswerId    int    `json:"answer_id"`
	Status      string `json:"status"`
	Comment     string `json:"comment"`
}
