package repositories

import (
	"beverages-booking/models"
	"database/sql"
	"fmt"
)

type AdminRepository struct {
	db   *sql.DB
	transaction *sql.Tx
}

func NewAdminRepository(dbHAndler *sql.DB) *AdminRepository {
	return &AdminRepository{
		db: dbHAndler,
	}
}

func (ar AdminRepository) AdminLogin(username, password string) (*models.Admin, error) {
	admin := models.Admin{}
	query := "SELECT id, username, password FROM admins WHERE username=? AND password=?"
	err := ar.db.QueryRow(query, username, password).Scan(&admin.ID, &admin.Username, &admin.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	return &admin, nil
}
