package postgres

import (
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/repository"
	"github.com/STUD-IT-team/bauman-legends-backend/internal/domain/response"
	"github.com/jmoiron/sqlx"
)

type TaskStorage struct {
	db *sqlx.DB
}

func NewTaskStorage(dataSource string) (repository.IUserAuthStorage, error) {
	db, err := sqlx.Open("pgx", dataSource)
	if err != nil {
		return nil, err
	}
	return &UserAuthStorage{
		db: db,
	}, err
}

func (r *TaskStorage) GetTaskTypes() (*response.TaskTypes, error) {
	var types []response.TaskType

	query := `select team_task.team_id, task.type_id, count(task.type_id) 
			  from team_task left join task on team_task.task_id = task.id 
			  where team_id=$1 group by team_id, task.type_id;`

	r.db.QueryRow(query, teamID)

	r.db.Select(&types, query)
	return nil, nil

}
