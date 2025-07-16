package users

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

func (r *RepositoryImpl) Save(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	if err := r.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}
