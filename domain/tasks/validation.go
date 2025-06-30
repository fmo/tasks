package tasks

import (
	"fmt"
	"github.com/fmo/tasks/models"
	"github.com/fmo/tasks/project_errors"
)

func Validate(task *models.Task) error {
	if task.Priority < 1 || task.Priority > 5 {
		return fmt.Errorf("%w", project_errors.ErrPriorityRange)
	}

	return nil
}
