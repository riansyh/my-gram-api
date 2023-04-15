package services

import (
	"errors"
	"my-gram/helpers"
	"my-gram/models"
	"my-gram/repositories"
)

type UserService interface {
	Register(user *models.User) (*models.User, error)
	Login(email, password string) (string, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(user *models.User) (*models.User, error) {
	err := us.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) Login(email, password string) (string, error) {
	user, err := us.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if !helpers.ComparePass([]byte(user.Password), []byte(password)) {
		return "", errors.New("invalid email/password")
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	return token, nil
}
