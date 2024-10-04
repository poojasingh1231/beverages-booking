package services

import (
	"beverages-booking/repositories"
	"beverages-booking/models"
	"errors"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us UserService) UserLogin(username, password string) (*models.User, error) {

	user, err := us.userRepository.UserLogin(username, password)
	if err != nil {
		return user, errors.New("invalid credentials")
	}

	return user, nil
}
