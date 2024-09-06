package response

type GetMediaTask struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Points      int    `json:"points"`
	VideoUrl    string `json:"video_url"`
}
