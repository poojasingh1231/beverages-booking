package services

import (
	"beverages-booking/repositories"
	"beverages-booking/models"
	"errors"
	"net/http"
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

func (us UserService) UserLogout() {
}

func validateUser(user *models.User) *models.ResponseError {
	if user.Username == "" {
		return &models.ResponseError{
			Message: "Invalid username",
			Status:  http.StatusBadRequest,
		}
	}

	if user.Password == "" {
		return &models.ResponseError{
			Message: "Invalid password",
			Status:  http.StatusBadRequest}
	}

	return nil
}


func (us UserService) CreateUser(user *models.User) (*models.User, *models.ResponseError) {
	responseErr := validateUser(user)
	if responseErr != nil {
		return nil, responseErr
	}

	return us.userRepository.CreateUser(user)
}

