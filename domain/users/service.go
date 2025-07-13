package users

import "github.com/fmo/tasks/models"

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) Create(user *models.User) (*models.User, error) {
	return s.r.Save(user)
}
