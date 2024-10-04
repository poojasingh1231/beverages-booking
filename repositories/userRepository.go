package repositories

import (
	"beverages-booking/models"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db   *sql.DB
	transaction *sql.Tx
}

func NewUserRepository(dbHandler *sql.DB) *UserRepository {
	return &UserRepository{
		db: dbHandler,
	}
}

func (ur UserRepository) UserLogin(username, password string) (*models.User, error) {
	user := models.User{}
	query := "SELECT id, username, password FROM users WHERE username=? AND password=?"
	err := ur.db.QueryRow(query, username, password).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	return &user, nil
}
