// userService.go
package services

import (
	"chat-app/pkg/models"
	"chat-app/pkg/repository"
	"errors"
)

type UserService interface {
	Authenticate(username, password string) (*models.User, error)
	Register(registerReq models.RegisterRequest) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// Authenticate verifies the user's credentials without generating JWT
func (s *userService) Authenticate(username, password string) (*models.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("authentication failed")
	}
	return user, nil
}

// Register creates a new user account
func (s *userService) Register(registerReq models.RegisterRequest) error {
	_, err := s.userRepo.FindByUsername(registerReq.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	newUser := models.User{
		Username: registerReq.Username,
		Password: registerReq.Password, // Expect password to be hashed before calling Register
	}

	return s.userRepo.Create(&newUser)
}
