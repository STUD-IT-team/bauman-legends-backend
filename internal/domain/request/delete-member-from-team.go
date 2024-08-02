package request

type DeleteMemberFromTeam struct {
	Session string `json:"session"`
	UserID  string `json:"user_id"`
}
