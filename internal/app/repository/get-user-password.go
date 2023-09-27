package repository

import (
	log "github.com/sirupsen/logrus"
)

func (r *Repository) GetUserPassword(email string) (password string, err error) {
	query := `select password from "user" where email = $1;`

	err = r.db.Select(&password, query, email)
	if err != nil {
		log.WithField(
			"origin.function", "GetUserPassword",
		).Errorf("Ошибка при получении пароля пользователя: %s", err.Error())
		return "", err
	}

	return password, nil
}
