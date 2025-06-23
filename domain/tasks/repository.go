package tasks

import "github.com/fmo/tasks/models"

type Repository interface {
	Save(task models.Task) (*models.Task, error)
}
