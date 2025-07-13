package users

import "github.com/fmo/tasks/models"

type Repository interface {
	Save(user *models.User) (*models.User, error)
}
