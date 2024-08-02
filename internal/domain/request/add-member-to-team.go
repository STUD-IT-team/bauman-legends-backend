package request

type AddMemberToTeam struct {
	Session   string `json:"session"`
	UserEmail string `json:"user_email"`
}
