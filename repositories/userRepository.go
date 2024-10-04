package repositories

import (
	"beverages-booking/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
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

func (ur UserRepository) CreateUser(user *models.User) (*models.User, *models.ResponseError) {
	query := `
		INSERT INTO runners(username, password, email)
		VALUES (?, ?, ?, ?)`

	res, err := ur.db.Exec(query, user.Username, user.Password, user.Email)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	userId, err := res.LastInsertId()
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &models.User{
		ID:        strconv.FormatInt(userId, 10),
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
	}, nil
}
