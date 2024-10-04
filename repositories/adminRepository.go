package repositories

import (
	"beverages_booking/models"
	"database/sql"
	"fmt"
)

func AdminLogin(db *sql.DB, username, password string) (*models.Admin, error) {
	admin := models.Admin{}
	query := "SELECT id, username, password FROM admins WHERE username=? AND password=?"
	err := db.QueryRow(query, username, password).Scan(&admin.ID, &admin.Username, &admin.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	return &admin, nil
}
