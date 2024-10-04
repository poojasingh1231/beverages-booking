package services

import (
	"beverages-booking/repositories"
	"beverages-booking/models"
)

type BeverageService struct {
	beverageRepository *repositories.BeverageRepository
}

func NewBeverageService(beverageRepository *repositories.BeverageRepository) *BeverageService {
	return &BeverageService{
		beverageRepository: beverageRepository,
	}
}


// GetAllBeveragesService retrieves all beverages from the database.
func (bs BeverageService) GetAllBeveragesService() ([]*models.Beverage, error) {
	beverages, err := bs.beverageRepository.GetAllBeverages()
	if err != nil {
		return nil, err
	}
	return beverages, nil
}

// CreateBeverageService creates a new beverage in the database.
func (bs BeverageService) CreateBeverageService(beverage *models.Beverage) (int64, error) {
	id, err := bs.beverageRepository.CreateBeverage(beverage)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// DeleteBeverageService deletes a beverage from the database by ID.
func (bs BeverageService) DeleteBeverageService(id int) error {
	err := bs.beverageRepository.DeleteBeverage(id)
	if err != nil {
		return err
	}
	return nil
}

// GetBeveragesByFiltersService retrieves beverages from the database by type filter.
func (bs BeverageService) GetBeveragesByFiltersService(beverageType string) ([]*models.Beverage , error) {
	beverages, err := bs.beverageRepository.GetBeveragesByFilters(beverageType)
	if err != nil {
		return nil, err
	}
	return beverages, nil
}

func (bs BeverageService) GetBeverageByIDService(id string) (*models.Beverage , error) {
	beverage, err := bs.beverageRepository.GetBeverageByID(id)
	if err != nil {
		return beverage, err
	}
	return beverage, nil
}
