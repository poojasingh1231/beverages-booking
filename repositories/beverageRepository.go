package repositories

import (
	"database/sql"
	// "log"
)

type Beverage struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func GetAllBeverages(db *sql.DB) ([]Beverage, error) {
	rows, err := db.Query("SELECT id, name, type, description, price FROM beverages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beverages []Beverage
	for rows.Next() {
		var beverage Beverage
		if err := rows.Scan(&beverage.ID, &beverage.Name, &beverage.Type, &beverage.Description, &beverage.Price); err != nil {
			return nil, err
		}
		beverages = append(beverages, beverage)
	}
	return beverages, nil
}

func CreateBeverage(db *sql.DB, beverage Beverage) (int64, error) {
	result, err := db.Exec("INSERT INTO beverages (name, type, description, price) VALUES (?, ?, ?, ?)", beverage.Name, beverage.Type, beverage.Description, beverage.Price)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func DeleteBeverage(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM beverages WHERE id = ?", id)
	return err
}
