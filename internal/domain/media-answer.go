package domain

type MediaAnswer struct {
	Id          int    `json:"id"`
	MediaTaskId int    `json:"media_task_id"`
	TeamId      int    `json:"team_id"`
	Points      int    `json:"points"`
	Status      string `json:"status"`
	Comment     string `json:"comment"`
	PhotoId     int    `json:"photo_id"`
	PhotoKey    string `json:"photo_key"`
	PhotoUrl    string `json:"photo_url"`
	PhotoAnswer []byte `json:"photo_answer"`
	PhotoType   string `json:"photo_type"`
}
