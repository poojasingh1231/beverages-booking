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

func NewAdminRepository(dbHandler *sql.DB) *AdminRepository {
	var repo = &AdminRepository{
		db: dbHandler,
	}
	repo.CreateAdminTable();
	return repo;
}

func (ar AdminRepository) CreateAdminTable() error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS admins (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL
	)`
	_, err := ar.db.Exec(createTableQuery)
	return err
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
