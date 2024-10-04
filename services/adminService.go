package services

import (
	"beverages-booking/repositories"
	"beverages-booking/context"
	"beverages-booking/models"
	"errors"
)

type AdminService struct {
	adminRepository *repositories.AdminRepository
}

func NewAdminService(adminRepository *repositories.AdminRepository) *AdminService {
	return &AdminService{
		adminRepository: adminRepository,
	}
}

func (as AdminService) AdminLogin(username, password string) (*models.Admin, error) {

	admin, err := as.adminRepository.AdminLogin(username, password)
	if err != nil {
		return admin, errors.New("invalid credentials")
	}
	context.IsLoggedIn = true
	context.IsAdmin = true
	return admin, nil
}

func (as AdminService) AdminLogout() {
	context.IsLoggedIn = false
	context.IsAdmin = false
}
