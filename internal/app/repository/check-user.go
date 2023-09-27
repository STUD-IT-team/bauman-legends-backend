package repository

import log "github.com/sirupsen/logrus"

func (r *Repository) CheckUser(email string) (exists bool, err error) {
	query := `select exists (select 1 from "user" where email = $1)`

	err = r.db.Select(&exists, query, email)
	if err != nil {
		log.WithField(
			"origin.function", "CheckUser",
		).Errorf("Ошибка при проверке существования пользователя: %s", err.Error())
		return false, err
	}

	return exists, nil
}
