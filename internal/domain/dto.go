package domain

import "time"

// Team
//
// Таблица команд
type Team struct {
	ID    string `db:"id"`
	Title string `db:"title"`
}

// Role
//
// Таблица ролей в команде
type Role struct {
	ID    string `db:"id"`
	Title string `db:"title"`
}

// User
//
// Таблица пользователей
type User struct {
	ID             string `db:"id"`
	PhoneNumber    string `db:"phone_number"`
	Email          string `db:"email"`
	EmailConfirmed bool   `db:"email_confirmed"`
	Telegram       string `db:"telegram"`
	VK             string `db:"vk"`
	StudyGroup     string `db:"study_group"`
	FIO            string `db:"fio"`
	TeamID         string `db:"team_db"`
	RoleID         string `db:"role_id"`
	IsAdmin        bool   `db:"is_admin"`
}

// Session
//
// Таблица пользовательских сессий
type Session struct {
	ID            string    `db:"id"`
	UserID        string    `db:"user_id"`
	Token         string    `db:"token"`
	ExpireAt      time.Time `db:"expire_at"`
	EnteredAt     time.Time `db:"entered_at"`
	ClientIP      string    `db:"client_ip"`
	ClientBrowser string    `db:"client_browser"`
	ClientOS      string    `db:"client_os"`
	ClientGeo     string    `db:"client_geo"`
}

// Difficulty
//
// Таблица сложностей заданий
type Difficulty struct {
	ID    string `db:"id"`
	Title string `db:"title"`
}

// TaskType
//
// Таблица типов заданий
type TaskType struct {
	ID    string `db:"id"`
	Title string `db:"title"`
}

// Task
//
// Таблица заданий
type Task struct {
	ID           string    `db:"id"`
	Title        string    `db:"title"`
	Description  string    `db:"description"`
	TimeLimit    time.Time `db:"time_limit"`
	DifficultyID string    `db:"difficulty_id"`
	TypeID       string    `db:"type_id"`
}

// Answer
//
// Таблица ответов на задания
type Answer struct {
	ID           string `db:"id"`
	AnswerTypeID string `db:"answer_type_id"`
	Contents     string `db:"contents"`
}

// AnswerType
//
// Таблица типов ответов (текст, QR-код)
type AnswerType struct {
	ID    string `db:"id"`
	Title string `db:"title"`
}

// TaskAnswer
//
// Таблица связи заданий и ответов на них
type TaskAnswer struct {
	ID       string `db:"id"`
	TaskID   string `db:"task_id"`
	AnswerID string `db:"answer_id"`
}

// TeamTask
//
// Таблица связи заданий и команд
type TeamTask struct {
	ID               string    `db:"id"`
	TaskID           string    `db:"task_id"`
	TeamID           string    `db:"team_id"`
	StartTime        time.Time `db:"start_time"`
	EndTime          time.Time `db:"end_time"`
	AdditionalPoints int       `db:"additional_points"`
}

// Secret
//
// Таблица секретных заданий
type Secret struct {
	ID          string `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

// TeamSecret
//
// Таблица связи секретных заданий и команд
type TeamSecret struct {
	ID        string    `db:"id"`
	SecretID  string    `db:"secret_id"`
	TeamID    string    `db:"team_id"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}
