package repository

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/request"
	log "github.com/sirupsen/logrus"
)

func (r *Repository) CreateUser(user request.Register) (userID string, err error) {
	query := `insert into "user" (
                    password, 
                    phone_number, 
                    email, 
                    telegram, 
                    vk, 
                    "group", 
                    name
                    ) values (
                              $1, $2, $3, $4, $5, $6, $7
                    ) returning id;
`
	err = r.db.Select(&userID, query,
		user.Password,
		user.PhoneNumber,
		user.Email,
		user.Telegram,
		user.VK,
		user.Group,
		user.Name)
	if err != nil {
		log.WithField(
			"origin.function", "CreateUser",
		).Errorf("Ошибка при создании пользователя: %s", err.Error())
		return "", err
	}

	return userID, nil
}
