package repository

import (
	log "github.com/sirupsen/logrus"
)

func (r *Repository) GetUserID(email string) (userID string, err error) {
	query := `select id from "user" where email = $1;`

	err = r.db.Select(&userID, query, email)
	if err != nil {
		log.WithField(
			"origin.function", "GetUserID",
		).Errorf("Ошибка при получении идентификатора пользователя: %s", err.Error())
		return "", err
	}

	return userID, nil
}
