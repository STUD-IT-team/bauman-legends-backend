package domain

// User
//
// Таблица пользователей
type User struct {
	ID             ID     `db:"id"`
	Password       string `db:"password"`
	PhoneNumber    string `db:"phone_number"`
	Email          string `db:"email"`
	EmailConfirmed bool   `db:"email_confirmed"`
	Telegram       string `db:"telegram"`
	VK             string `db:"vk"`
	StudyGroup     string `db:"group"`
	FIO            string `db:"name"`
	TeamID         *ID    `db:"team_id"`
	RoleID         *int   `db:"role_id"`
	IsAdmin        bool   `db:"is_admin"`
}
