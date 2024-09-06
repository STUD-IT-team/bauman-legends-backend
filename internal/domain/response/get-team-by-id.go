package response

type GetTeamByID struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Members []Member
	Captain Member
	Points  int `json:"points"`
}
