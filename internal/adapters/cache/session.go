package cache

import (
	"github.com/STUD-IT-team/bauman-legends-backend/pkg/cache"
	"time"
)

// Session
//
// Запись о сессии для хранения в кэше
type Session struct {
	UserID        string
	ExpireAt      time.Time
	EnteredAt     time.Time
	ClientIP      string
	ClientBrowser string
	ClientOS      string
}

// NewSessionCache
//
// Создание нового объекта кэша сессий
func NewSessionCache() cache.ICache[string, Session] {
	return &cache.Cache[string, Session]{}
}
