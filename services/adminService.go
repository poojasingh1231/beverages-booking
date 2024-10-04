package services

import (
	"beverages_booking/repositories"
	"database/sql"
	"errors"
)

func AdminLoginService(db *sql.DB, username, password string) (repositories.Admin, error) {

	admin, err := repositories.AdminLogin(db, username, password)
	if err != nil {
		return admin, errors.New("invalid credentials")
	}

	return admin, nil
}
