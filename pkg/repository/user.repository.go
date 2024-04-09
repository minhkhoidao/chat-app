package repository

import (
	"chat-app/pkg/models"

	"github.com/jmoiron/sqlx"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
}

// userRepository implements UserRepository with a PostgreSQL backend
type userRepository struct {
	db *sqlx.DB
}

// NewUserRepository returns a new instance of a PostgreSQL UserRepository
func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// Create inserts a new user into the database
func (repo *userRepository) Create(user *models.User) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	row := repo.db.QueryRow(query, user.Username, user.Password)
	return row.Scan(&user.ID)
}

// FindByUsername looks up a user by username and returns it
func (repo *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password FROM users WHERE username = $1`
	err := repo.db.Get(&user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
