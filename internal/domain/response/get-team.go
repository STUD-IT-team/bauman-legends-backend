package response

type GetTeam struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Members []Member `json:"members"`
	Captain Member   `json:"captain"`
	Points  int      `json:"points"`
}

type Member struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Grope string `json:"group"`
	Email string `json:"email"` //когда возвращаем команду в главном меню -
	// на самом деле нахуй не нужен email
}
