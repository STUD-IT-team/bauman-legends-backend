package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/storage"
)

type SecStorage struct {
	db *pgxpool.Pool
}

const getSECByFilterQuery = `WITH sum_user AS (
    WITH count_user AS (
    SELECT count(*) as cnt, team_id
    FROM "user" JOIN team ON "user".team_id = team.id
    GROUP BY team_id
    )

    SELECT sum(CNT) as sum, master_class.id
    FROM team_master_class JOIN COUNT_USER on count_user.team_id = team_master_class.team_id
     JOIN master_class ON team_master_class.master_class_id = master_class.id
    GROUP BY master_class.id
)

SELECT sec.id, sec.name, description, "user".name, "user".telegram, "user".phone_number, started_at, ended_at, capacity, COALESCE( SUM_USER.SUM, 0) as busy, master_class.id, ended_at
FROM sec JOIN master_class on sec.id = master_class.sec_id
         JOIN "user" ON "user".id = sec.responsible_id
        FULL JOIN team_master_class ON master_class.id = team_master_class.master_class_id
        FULL JOIN SUM_USER ON team_master_class.master_class_id = SUM_USER.id
ORDER BY  sec.id, started_at`

func (s *SecStorage) GetSECByFilter() ([]domain.Sec, error) {
	var secs []domain.Sec
	rows, err := s.db.Query(context.Background(), getSECByFilterQuery)
	if err != nil {
		log.WithField(
			"origin.function", "GetSECByFilter",
		).Errorf("ошибка запроса: %s", err.Error())
		return secs, err
	}
	defer rows.Close()
	for rows.Next() {
		var sec domain.Sec
		err = rows.Scan(
			&sec.Id,
			&sec.Name,
			&sec.Description,
			&sec.FIO,
			&sec.Telegram,
			&sec.Phone,
			&sec.StartedAt,
			&sec.EndedAt,
			&sec.Capacity,
			&sec.Busy,
			&sec.MasterClassId,
			&sec.EndedAt,
		)
		if err != nil {
			log.WithField(
				"origin.function", "GetSECByFilter",
			).Errorf("ошибка получения данных: %s", err.Error())
			return secs, err
		}

		secs = append(secs, sec)
	}

	return secs, rows.Err()
}

const getSecByIdQuery = `WITH sum_user AS (
    WITH count_user AS (
    SELECT count(*) as cnt, team_id
    FROM "user" JOIN team ON "user".team_id = team.id
    GROUP BY team_id
    )

    SELECT sum(CNT) as sum, master_class.id
    FROM team_master_class JOIN COUNT_USER on count_user.team_id = team_master_class.team_id
     JOIN master_class ON team_master_class.master_class_id = master_class.id
    GROUP BY master_class.id
)

SELECT sec.id, sec.name, description, "user".name, "user".telegram, "user".phone_number, started_at, ended_at, capacity, COALESCE( SUM_USER.SUM, 0), uuid_media as busy, master_class.id, ended_at
FROM sec JOIN master_class on sec.id = master_class.sec_id
         JOIN "user" ON "user".id = sec.responsible_id
        FULL JOIN team_master_class ON master_class.id = team_master_class.master_class_id
        FULL JOIN SUM_USER ON team_master_class.master_class_id = SUM_USER.id
        JOIN media_obj ON master_class.media_id = media_obj.id
WHERE sec.id = $1
ORDER BY  sec.id, started_at;`

func (s *SecStorage) GetSecByID(secId int) ([]domain.Sec, error) {
	var secs []domain.Sec
	rows, err := s.db.Query(context.Background(), getSecByIdQuery, secId)
	if err != nil {
		log.WithField(
			"origin.function", "GetSecByID",
		).Errorf("ошибка запроса: %s", err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var sec domain.Sec
		err = rows.Scan(
			&sec.Id,
			&sec.Name,
			&sec.Description,
			&sec.FIO,
			&sec.Telegram,
			&sec.Phone,
			&sec.StartedAt,
			&sec.EndedAt,
			&sec.Capacity,
			&sec.Busy,
			&sec.PhotoUrl,
			&sec.MasterClassId,
			&sec.EndedAt,
		)
		sec.PhotoUrl = os.Getenv("MINIO_URL") + consts.SECBucket + "/" + sec.PhotoUrl
		if err != nil {
			log.WithField(
				"origin.function", "GetSecByID",
			).Errorf("ошибка получения данных: %s", err.Error())
			return secs, err
		}

		secs = append(secs, sec)
	}
	return secs, nil
}

const getSecByTeamIdQuery = `WITH sum_user AS (
    WITH count_user AS (
    SELECT count(*) as cnt, team_id
    FROM "user" JOIN team ON "user".team_id = team.id
    GROUP BY team_id
    )

    SELECT sum(CNT) as sum, master_class.id
    FROM team_master_class JOIN COUNT_USER on count_user.team_id = team_master_class.team_id
     JOIN master_class ON team_master_class.master_class_id = master_class.id
    GROUP BY master_class.id
)

SELECT sec.id, sec.name, description, "user".name, "user".telegram, "user".phone_number, started_at, ended_at, capacity, COALESCE( SUM_USER.SUM, 0), uuid_media as busy, master_class.id, ended_at
FROM sec JOIN master_class on sec.id = master_class.sec_id
         JOIN "user" ON "user".id = sec.responsible_id
        FULL JOIN team_master_class ON master_class.id = team_master_class.master_class_id
        FULL JOIN SUM_USER ON team_master_class.master_class_id = SUM_USER.id
        JOIN media_obj ON master_class.media_id = media_obj.id
WHERE team_master_class.team_id = $1
ORDER BY  sec.id, started_at`

func (s *SecStorage) GetSecByTeamId(teamId int) ([]domain.Sec, error) {
	var secs []domain.Sec
	rows, err := s.db.Query(context.Background(), getSecByTeamIdQuery, teamId)
	if err != nil {
		log.WithField(
			"origin.function", "GetSecByTeamId",
		).Errorf("ошибка запроса: %s", err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var sec domain.Sec
		err = rows.Scan(
			&sec.Id,
			&sec.Name,
			&sec.Description,
			&sec.FIO,
			&sec.Telegram,
			&sec.Phone,
			&sec.StartedAt,
			&sec.EndedAt,
			&sec.Capacity,
			&sec.Busy,
			&sec.PhotoUrl,
			&sec.MasterClassId,
			&sec.EndedAt,
		)
		sec.PhotoUrl = os.Getenv("MINIO_URL") + consts.SECBucket + "/" + sec.PhotoUrl
		if err != nil {
			log.WithField(
				"origin.function", "GetSecByTeamId",
			).Errorf("ошибка получения данных: %s", err.Error())
			return secs, err
		}

		secs = append(secs, sec)
	}
	return secs, nil
}

const createRegisterOnSecQuery = `INSERT INTO team_master_class (team_id, master_class_id) VALUES ($1, $2)`

func (s *SecStorage) CreateRegisterOnSEC(masterClassId, teamId int) error {
	_, err := s.db.Exec(context.Background(), createRegisterOnSecQuery, teamId, masterClassId)
	if err != nil {
		log.WithField(
			"origin.function", "CreateRegisterOnSEC",
		).Errorf("ошибка запроса: %s", err.Error())
		return err
	}

	return nil
}

const checkRegisterOnMasterClassQuery = `SELECT EXISTS (SELECT * FROM team_master_class WHERE team_id = $1 AND master_class_id = $2)`

func (s *SecStorage) CheckRegisterOnMasterClass(masterClassId, teamId int) (bool, error) {
	var exist bool
	err := s.db.QueryRow(context.Background(), checkRegisterOnMasterClassQuery, teamId, masterClassId).Scan(&exist)
	if err != nil {
		log.WithField(
			"origin.function", "CheckRegisterOnMasterClass",
		).Errorf("ошибка запроса: %s", err.Error())
		return false, err
	}

	return exist, nil
}

const deleteRegisterOnSECQuery = `DELETE FROM team_master_class WHERE team_id = $1 AND master_class_id = $2`

func (s *SecStorage) DeleteRegisterOnSEC(masterClassId, teamId int) error {
	_, err := s.db.Exec(context.Background(), deleteRegisterOnSECQuery, teamId, masterClassId)
	if err != nil {
		log.WithField(
			"origin.function", "DeleteRegisterOnSEC",
		).Errorf("ошибка запроса: %s", err.Error())
		return err
	}

	return nil
}

const getSECAdminByFilterQuery = `WITH sum_user AS (
    WITH count_user AS (
    SELECT count(*) as cnt, team_id
    FROM "user" JOIN team ON "user".team_id = team.id
    GROUP BY team_id
    )

    SELECT sum(CNT) as sum, master_class.id
    FROM team_master_class JOIN COUNT_USER on count_user.team_id = team_master_class.team_id
     JOIN master_class ON team_master_class.master_class_id = master_class.id
    GROUP BY master_class.id
)

SELECT sec.id, sec.name, description, "user".name, "user".telegram, "user".phone_number, started_at, capacity, COALESCE( SUM_USER.SUM, 0) as busy, master_class.id, ended_at
FROM sec JOIN master_class on sec.id = master_class.sec_id
         JOIN "user" ON "user".id = sec.responsible_id
        FULL JOIN team_master_class ON master_class.id = team_master_class.master_class_id
        FULL JOIN SUM_USER ON team_master_class.master_class_id = SUM_USER.id
ORDER BY  sec.id, started_at`

func (s *SecStorage) GetSECAdmin() ([]domain.Sec, error) {
	var secs []domain.Sec
	rows, err := s.db.Query(context.Background(), getSECAdminByFilterQuery)
	if err != nil {
		log.WithField(
			"origin.function", "GetSECAdmin",
		).Errorf("ошибка запроса: %s", err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var sec domain.Sec
		err = rows.Scan(
			&sec.Id,
			&sec.Name,
			&sec.Description,
			&sec.FIO,
			&sec.Telegram,
			&sec.Phone,
			&sec.StartedAt,
			&sec.Capacity,
			&sec.Busy,
			&sec.MasterClassId,
			&sec.EndedAt,
		)
		if err != nil {
			log.WithField(
				"origin.function", "GetSECAdmin",
			).Errorf("ошибка получения данных: %s", err.Error())
			return secs, err
		}

		secs = append(secs, sec)
	}
	return secs, nil
}

const getSECAdminById = `WITH sum_user AS (
    WITH count_user AS (
    SELECT count(*) as cnt, team_id
    FROM "user" JOIN team ON "user".team_id = team.id
    GROUP BY team_id
    )

    SELECT sum(CNT) as sum, master_class.id
    FROM team_master_class JOIN COUNT_USER on count_user.team_id = team_master_class.team_id
     JOIN master_class ON team_master_class.master_class_id = master_class.id
    GROUP BY master_class.id
)

SELECT sec.id, sec.name, description, "user".name, "user".telegram, "user".phone_number, started_at, capacity, COALESCE( SUM_USER.SUM, 0), uuid_media as busy, master_class.id, ended_at
FROM sec JOIN master_class on sec.id = master_class.sec_id
         JOIN "user" ON "user".id = sec.responsible_id
        FULL JOIN team_master_class ON master_class.id = team_master_class.master_class_id
        FULL JOIN SUM_USER ON team_master_class.master_class_id = SUM_USER.id
        JOIN media_obj ON master_class.media_id = media_obj.id
WHERE sec.id = $1
ORDER BY  sec.id, started_at`

func (s *SecStorage) GetSECAdminById(secId int) ([]domain.Sec, error) {
	var secs []domain.Sec
	rows, err := s.db.Query(context.Background(), getSECAdminById, secId)
	if err != nil {
		log.WithField(
			"origin.function", "GetSECAdminById",
		).Errorf("ошибка запроса: %s", err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var sec domain.Sec
		err = rows.Scan(
			&sec.Id,
			&sec.Name,
			&sec.Description,
			&sec.FIO,
			&sec.Telegram,
			&sec.Phone,
			&sec.StartedAt,
			&sec.Capacity,
			&sec.Busy,
			&sec.PhotoUrl,
			&sec.MasterClassId,
			&sec.EndedAt,
		)
		sec.PhotoUrl = os.Getenv("MINIO_URL") + consts.SECBucket + "/" + sec.PhotoUrl
		if err != nil {
			log.WithField(
				"origin.function", "GetSECAdminById",
			).Errorf("ошибка получения данных: %s", err.Error())
			return secs, err
		}

		secs = append(secs, sec)
	}
	return secs, nil
}

const checkRegisterOnSecQuery = `
SELECT EXISTS (SELECT * FROM team_master_class
                                 JOIN master_class ON team_master_class.master_class_id = master_class.id
               WHERE sec_id = (SELECT sec_id FROM master_class WHERE master_class.id = $1) AND team_id = $2);`

func (s *SecStorage) CheckRegisterOnSec(secId, teamId int) (bool, error) {
	var exist bool
	err := s.db.QueryRow(context.Background(), checkRegisterOnSecQuery, secId, teamId).Scan(&exist)
	if err != nil {
		log.WithField(
			"origin.function", "CheckRegisterOnSec",
		).Errorf("ошибка получения данных: %s", err.Error())
		return false, err
	}

	return exist, nil
}

const checkIntersectionTimeIntervalQuery = `SELECT EXISTS( SELECT * FROM master_class JOIN team_master_class ON master_class.id = team_master_class.master_class_id
 WHERE   TSRANGE(started_at, ended_at) && TSRANGE((
     SELECT started_at FROM master_class WHERE id = $1
 ), (
     SELECT ended_at FROM master_class WHERE id = $1
 )) IS TRUE AND team_id = $2)`

func (s *SecStorage) CheckIntersectionTimeInterval(masterClassId, teamId int) (bool, error) {
	var exist bool
	err := s.db.QueryRow(context.Background(), checkIntersectionTimeIntervalQuery, masterClassId, teamId).Scan(&exist)
	if err != nil {
		log.WithField(
			"origin.function", "CheckIntersectionTimeInterval",
		).Errorf("ошибка получения данных: %s", err.Error())
		return false, err
	}

	return exist, nil
}

const checkMasterClassIsExist = `SELECT EXISTS(SELECT * FROM master_class WHERE id = $1)`

func (s *SecStorage) CheckMasterClassIsExist(masterClassId int) (bool, error) {
	var exist bool
	err := s.db.QueryRow(context.Background(), checkMasterClassIsExist, masterClassId).Scan(&exist)
	if err != nil {
		log.WithField(
			"origin.function", "CheckMasterClassIsExist",
		).Errorf("ошибка получения данных: %s", err.Error())
		return false, err
	}

	return exist, nil
}

const checkCountUserRegisterOnMasterClassQuery = `SELECT master_class.capacity - (
SELECT COUNT("user".team_id)
FROM master_class JOIN team_master_class ON master_class.id = team_master_class.master_class_id
JOIN "user" ON "user".team_id = team_master_class.team_id
WHERE master_class.id = $1) - (
    SELECT COUNT("user".team_id) FROM "user" WHERE team_id = $2
    ) 
FROM master_class WHERE id = $1`

func (s *SecStorage) CheckMasterClassBusyPlaceById(masterClassId, teamId int) (int, error) {
	var count int
	err := s.db.QueryRow(context.Background(), checkCountUserRegisterOnMasterClassQuery, masterClassId, teamId).Scan(&count)
	if err != nil {
		log.WithField(
			"origin.function", "CheckMasterClassBusyPlaceById",
		).Errorf("ошибка получения данных: %s", err.Error())
		return 0, err
	}

	return count, nil
}

const checkMasterClassTimeQuery = `SELECT started_at < NOW() FROM master_class WHERE id = $1`

func (s *SecStorage) CheckMasterClassTime(masterClass int) (bool, error) {
	var exist bool
	err := s.db.QueryRow(context.Background(), checkMasterClassTimeQuery, masterClass).Scan(&exist)
	if err != nil {
		log.WithField(
			"origin.function", "CheckMasterClassTime",
		).Errorf("ошибка получения данных: %s", err.Error())
		return false, err
	}

	return exist, nil
}

const getMasterClassByIdQuery = `WITH sum_user AS (WITH count_user AS (
    SELECT count(*) as cnt, team_id
    FROM "user" JOIN team ON "user".team_id = team.id
    GROUP BY team_id

)
                  SELECT sum(CNT) as sum, master_class_id
                  FROM team_master_class JOIN COUNT_USER on count_user.team_id = team_master_class.team_id
                  GROUP BY master_class_id
)

SELECT sec.id, sec.name, description, "user".name, "user".telegram, "user".phone_number, started_at, ended_at, capacity, COALESCE( SUM_USER.SUM, 0)
FROM sec JOIN master_class on sec.id = master_class.sec_id
         JOIN "user" ON "user".id = sec.responsible_id
         full JOIN team_master_class ON master_class.id = team_master_class.master_class_id
         full JOIN SUM_USER ON team_master_class.master_class_id = SUM_USER.master_class_id
WHERE master_class.id = $1;`

func (s *SecStorage) GetMasterClassByID(masterClassId int) (domain.Sec, error) {
	var sec domain.Sec
	err := s.db.QueryRow(context.Background(), getMasterClassByIdQuery, masterClassId).Scan(
		&sec.Id,
		&sec.Name,
		&sec.Description,
		&sec.FIO,
		&sec.Telegram,
		&sec.Phone,
		&sec.StartedAt,
		&sec.EndedAt,
		&sec.Capacity,
		&sec.Busy,
	)
	if err != nil {
		log.WithField(
			"origin.function", "GetMasterClassByID",
		).Errorf("ошибка чтения мк: %s", err.Error())
		return domain.Sec{}, err
	}

	return sec, nil
}

func NewSecStorage(dataSource string) (storage.SECStorage, error) {
	config, err := pgxpool.ParseConfig(dataSource)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	return &SecStorage{
		db: db,
	}, nil
}
