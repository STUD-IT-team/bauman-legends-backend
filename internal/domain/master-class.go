package domain

type MasterClass struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	StartAt       string `json:"start_at"`
	EndAt         string `json:"end_at"`
	Responsible   Member `json:"responsible"`
	CountAll      int    `json:"count_all"`
	CountFree     int    `json:"count_free"`
	CountReserve  int    `json:"count_reserve"`
	PlacePhotoUrl string `json:"place_photo_url"`
}
