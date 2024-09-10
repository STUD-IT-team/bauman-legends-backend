package response

type GetTeam struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Members []Member `json:"members"`
	Captain Member   `json:"captain"`
	Points  int      `json:"points"`
}

type Member struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Grope string `json:"grope"`
	Email string `json:"email"`
}
