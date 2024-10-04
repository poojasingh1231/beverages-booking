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

func (ar *AdminRepository) AdminUserExists(userId int, userName string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM admins WHERE id = ? AND username = ?)"
	err := ar.db.QueryRow(query, userId, userName).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return exists, nil
}
