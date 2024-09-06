package domain

type MediaTask struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Points      int         `json:"points"`
	VideoId     int         `json:"video_id"`
	VideoKey    string      `json:"video_key"`
	VideoUrl    string      `json:"video_url"`
	Answer      MediaAnswer `json:"answer"`
}
