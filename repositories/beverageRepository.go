package repositories

import (
	"database/sql"
	"fmt"
	"beverages-booking/models"
)

type BeverageRepository struct {
	db         *sql.DB
	transaction *sql.Tx
}

func NewBeverageRepository(dbHandler *sql.DB) *BeverageRepository {
	return &BeverageRepository{
		db: dbHandler,
	}
}

func (br BeverageRepository) GetAllBeverages() ([]*models.Beverage, error) {
	rows, err := br.db.Query("SELECT id, name, type, description, price FROM beverages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beverages []*models.Beverage
	for rows.Next() {
		beverage := new(models.Beverage) // Initialize as a pointer
		if err := rows.Scan(&beverage.ID, &beverage.Name, &beverage.Type, &beverage.Description, &beverage.Price); err != nil {
			return nil, err
		}
		beverages = append(beverages, beverage)
	}
	return beverages, nil
}

func (br BeverageRepository) CreateBeverage(beverage *models.Beverage) (int64, error) {
	result, err := br.db.Exec("INSERT INTO beverages (name, type, description, price) VALUES (?, ?, ?, ?)", beverage.Name, beverage.Type, beverage.Description, beverage.Price)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (br BeverageRepository) DeleteBeverage(id int) error {
	_, err := br.db.Exec("DELETE FROM beverages WHERE id = ?", id)
	return err
}

func (br BeverageRepository) GetBeveragesByFilters(beverageType string) ([]*models.Beverage, error) {
	var query string
	if beverageType == "" {
		query = "SELECT id, name, type, description, price FROM beverages"
	} else {
		query = "SELECT id, name, type, description, price FROM beverages WHERE type = ?"
	}

	rows, err := br.db.Query(query, beverageType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beverages []*models.Beverage
	for rows.Next() {
		beverage := new(models.Beverage) // Initialize as a pointer
		if err := rows.Scan(&beverage.ID, &beverage.Name, &beverage.Type, &beverage.Description, &beverage.Price); err != nil {
			return nil, err
		}
		beverages = append(beverages, beverage)
	}
	return beverages, nil
}

func (br BeverageRepository) GetBeverageByID(id string) (*models.Beverage, error) {
	beverage := new(models.Beverage) // Initialize as a pointer
	query := "SELECT id, name, type, description, price FROM beverages WHERE id = ?"
	err := br.db.QueryRow(query, id).Scan(&beverage.ID, &beverage.Name, &beverage.Type, &beverage.Description, &beverage.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return beverage, fmt.Errorf("beverage not found")
		}
		return beverage, err
	}
	return beverage, nil
}
