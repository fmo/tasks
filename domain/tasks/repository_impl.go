package tasks

import (
	"database/sql"
	"github.com/fmo/tasks/models"
)

type RepositoryImpl struct {
	*sql.DB
}

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

func (r RepositoryImpl) Save(task models.Task) (*models.Task, error) {
	query := `INSERT INTO tasks (name, owner, priority)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.QueryRow(query, task.Name, task.Owner, task.Priority).Scan(&task.ID)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
