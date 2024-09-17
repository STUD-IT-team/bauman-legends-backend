package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"

	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/storage"
)

type SecStorage struct {
	db *pgx.Conn
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

SELECT sec.id, sec.name, description, "user".name, "user".telegram, "user".phone_number, started_at, ended_at, capacity, COALESCE( SUM_USER.SUM, 0) as busy
FROM sec JOIN master_class on sec.id = master_class.sec_id
         JOIN "user" ON "user".id = sec.responsible_id
        FULL JOIN team_master_class ON master_class.id = team_master_class.master_class_id
        FULL JOIN SUM_USER ON team_master_class.master_class_id = SUM_USER.id
WHERE started_at > now()
ORDER BY  sec.id, started_at`

func (s *SecStorage) GetSECByFilter() ([]domain.Sec, error) {
	var secs []domain.Sec
	rows, err := s.db.Query(context.Background(), getSECByFilterQuery)
	if err != nil {
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
		)
		if err != nil {
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

SELECT sec.id, sec.name, description, "user".name, "user".telegram, "user".phone_number, started_at, ended_at, capacity, COALESCE( SUM_USER.SUM, 0), uuid_media as busy
FROM sec JOIN master_class on sec.id = master_class.sec_id
         JOIN "user" ON "user".id = sec.responsible_id
        FULL JOIN team_master_class ON master_class.id = team_master_class.master_class_id
        FULL JOIN SUM_USER ON team_master_class.master_class_id = SUM_USER.id
        JOIN media_obj ON master_class.media_id = media_obj.id
WHERE started_at > now() and sec.id = $1
ORDER BY  sec.id, started_at;`

func (s *SecStorage) GetSecByID(secId int) ([]domain.Sec, error) {
	var secs []domain.Sec
	rows, err := s.db.Query(context.Background(), getSecByIdQuery, secId)
	if err != nil {
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
		)
		if err != nil {
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

SELECT sec.id, sec.name, description, "user".name, "user".telegram, "user".phone_number, started_at, ended_at, capacity, COALESCE( SUM_USER.SUM, 0), uuid_media as busy
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
		)
		if err != nil {
			return secs, err
		}

		secs = append(secs, sec)
	}
	return secs, nil
}

const createRegisterOnSecQuery = `INSERT INTO team_master_class (team_id, master_class_id) 
	VALUES ($1, (SELECT master_class.id FROM master_class WHERE sec_id = $2 AND started_at = $3))`

func (s *SecStorage) CreateRegisterOnSEC(secId int, time string, teamId int) error {
	_, err := s.db.Exec(context.Background(), createRegisterOnSecQuery, teamId, secId, time)
	if err != nil {
		return err
	}

	return nil
}

const checkRegisterOnMasterClassQuery = `SELECT EXISTS (SELECT * FROM team_master_class WHERE team_id = $1 AND master_class_id = (
    SELECT master_class.id FROM master_class WHERE sec_id = $2 AND started_at = $3
))`

func (s *SecStorage) CheckRegisterOnMasterClass(secId int, time string, teamId int) (bool, error) {
	var exist bool
	err := s.db.QueryRow(context.Background(), checkRegisterOnMasterClassQuery, teamId, secId, time).Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil
}

const deleteRegisterOnSECQuery = `DELETE FROM team_master_class WHERE team_id = $1 AND master_class_id = (
    SELECT master_class.id FROM master_class WHERE sec_id = $2 AND started_at = $3
)`

func (s *SecStorage) DeleteRegisterOnSEC(secId int, time string, teamId int) error {
	_, err := s.db.Exec(context.Background(), deleteRegisterOnSECQuery, teamId, secId, time)
	if err != nil {
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

SELECT sec.id, sec.name, description, "user".name, "user".telegram, "user".phone_number, started_at, duration, capacity, COALESCE( SUM_USER.SUM, 0) as busy
FROM sec JOIN master_class on sec.id = master_class.sec_id
         JOIN "user" ON "user".id = sec.responsible_id
        FULL JOIN team_master_class ON master_class.id = team_master_class.master_class_id
        FULL JOIN SUM_USER ON team_master_class.master_class_id = SUM_USER.id
ORDER BY  sec.id, started_at`

func (s *SecStorage) GetSECAdmin() ([]domain.Sec, error) {
	var secs []domain.Sec
	rows, err := s.db.Query(context.Background(), getSECAdminByFilterQuery)
	if err != nil {
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
		)
		if err != nil {
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

SELECT sec.id, sec.name, description, "user".name, "user".telegram, "user".phone_number, started_at, ended_at, capacity, COALESCE( SUM_USER.SUM, 0), uuid_media as busy
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
		)
		if err != nil {
			return secs, err
		}

		secs = append(secs, sec)
	}
	return secs, nil
}

const checkRegisterOnSecQuery = `SELECT EXISTS (SELECT * FROM team_master_class 
    JOIN master_class ON team_master_class.master_class_id = master_class.id
                        WHERE sec_id = $1 AND team_id = $2)`

func (s *SecStorage) CheckRegisterOnSec(secId, teamId int) (bool, error) {
	var exist bool
	err := s.db.QueryRow(context.Background(), checkRegisterOnSecQuery, secId, teamId).Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil
}

const checkIntersectionTimeIntervalQuery = `SELECT EXISTS( SELECT * FROM master_class JOIN team_master_class ON master_class.id = team_master_class.master_class_id
 WHERE   TSRANGE(started_at, ended_at) && TSRANGE($1, (
     SELECT ended_at FROM master_class WHERE started_at = $1 AND sec_id = $3
 )) IS TRUE AND team_id = $2)`

func (s *SecStorage) CheckIntersectionTimeInterval(secId int, time string, teamId int) (bool, error) {
	var exist bool
	err := s.db.QueryRow(context.Background(), checkIntersectionTimeIntervalQuery, time, teamId, secId).Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil
}

const checkMasterClassIsExist = `SELECT EXISTS(SELECT * FROM master_class WHERE sec_id = $1 AND started_at = $2)`

func (s *SecStorage) CheckMasterClassIsExist(secId int, time string) (bool, error) {
	var exist bool
	err := s.db.QueryRow(context.Background(), checkMasterClassIsExist, secId, time).Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil
}

const checkCountUserRegisterOnMasterClassQuery = `SELECT master_class.capacity - (
SELECT COUNT("user".team_id)
FROM master_class JOIN team_master_class ON master_class.id = team_master_class.master_class_id
JOIN "user" ON "user".team_id = team_master_class.team_id
WHERE sec_id = $1 AND started_at = $2) - (
    SELECT COUNT("user".team_id) FROM "user" WHERE team_id = $3
    ) 
FROM master_class WHERE sec_id = $1 AND started_at = $2`

func (s *SecStorage) CheckMasterClassBusyPlaceById(secId int, time string, teamId int) (int, error) {
	var count int
	err := s.db.QueryRow(context.Background(), checkCountUserRegisterOnMasterClassQuery, secId, time, teamId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func NewSecStorage(dataSource string) (storage.SECStorage, error) {
	config, err := pgx.ParseConfig(dataSource)
	if err != nil {
		return nil, err
	}

	db, err := pgx.ConnectConfig(context.Background(), config)
	return &SecStorage{
		db: db,
	}, nil
}
