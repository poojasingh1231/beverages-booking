package services

import (
	"beverages-booking/repositories"
	"beverages-booking/models"
	"errors"
	"log"
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
	return admin, nil
}

func (as AdminService) AdminLogout() {
}

func (as AdminService) AdminUserExists(userId int, userName string) bool {
	exists, err := as.adminRepository.AdminUserExists(userId, userName)
	log.Printf("userid = %d, username= %s", userId, userName)
	if err != nil {
		log.Printf("Error checking if admin user exists: %v", err)
		return true
	}
	return exists
}



