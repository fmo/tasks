package tasks

import "github.com/fmo/tasks/models"

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) Create(task *models.Task) (*models.Task, error) {
	return s.r.Save(task)
}

func (s *Service) Update(task *models.Task) error {
	return s.r.Update(task)
}

// GO
// GO
// GO
// GO
// GO
// GO
// GO
// GO
// GO
// GO
// GO
// GO
// GO
// GO
// GO
// GO
