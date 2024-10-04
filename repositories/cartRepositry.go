package repositories

import (
	"database/sql"
	"fmt"
	"beverages-booking/models"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(dbHandler *sql.DB) *CartRepository {
	var repo = &CartRepository{
		db: dbHandler,
	}

	repo.CreateCartTable()
	return repo
}

func (cr CartRepository) CreateCartTable() error {
	// Correct SQL creation inside a function
	createTableQuery := `CREATE TABLE IF NOT EXISTS carts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		item_name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		price DECIMAL(10,2) NOT NULL,
		quantity INT NOT NULL,
		user_id INT NOT NULL,
		beverage_id INT NOT NULL
	)`
	_, err := cr.db.Exec(createTableQuery)
	if err != nil {
		return err
	}
	return nil
}



func (cr CartRepository) GetCartItems(userID int) ([]*models.Cart, error) {
	rows, err := cr.db.Query("SELECT id, user_id, beverage_id, item_name, description, price, quantity FROM carts WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []*models.Cart
	for rows.Next() {
		var cart models.Cart
		if err := rows.Scan(&cart.ID, &cart.UserID, &cart.BeverageID, &cart.ItemName, &cart.Description, &cart.Price, &cart.Quantity); err != nil {
			return nil, err
		}
		carts = append(carts, &cart)
	}
	return carts, nil
}

func (cr CartRepository) AddItem(userID int, cart *models.Cart) error {
	// Check if the item already exists in the user's cart
	var currentQuantity int
	err := cr.db.QueryRow("SELECT quantity FROM carts WHERE beverage_id = ? AND user_id = ?", cart.BeverageID, userID).Scan(&currentQuantity)

	if err == sql.ErrNoRows {
		_, err = cr.db.Exec("INSERT INTO carts (user_id, beverage_id, item_name, description, price, quantity) VALUES (?, ?, ?, ?, ?, ?)",
			userID, cart.BeverageID, cart.ItemName, cart.Description, cart.Price, 1) // Initialize quantity as 1
		if err != nil {
			return err
		}
	} else if err == nil {
		_, err = cr.db.Exec("UPDATE carts SET quantity = quantity + 1 WHERE beverage_id = ? AND user_id = ?", cart.BeverageID, userID)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

// RemoveItem decrements the quantity of an item, or removes it if the quantity reaches 0
func (cr CartRepository) RemoveItem(userID int, beverageID int) error {
	var currentQuantity int
	err := cr.db.QueryRow("SELECT quantity FROM carts WHERE beverage_id = ? AND user_id = ?", beverageID, userID).Scan(&currentQuantity)
	if err == sql.ErrNoRows {
		return fmt.Errorf("item not found in cart for this user")
	} else if err != nil {
		return err
	}

	if currentQuantity > 1 {
		// Decrement the quantity
		_, err = cr.db.Exec("UPDATE carts SET quantity = quantity - 1 WHERE beverage_id = ? AND user_id = ?", beverageID, userID)
		if err != nil {
			return err
		}
	} else {
		// Quantity is 1, so remove the item
		_, err = cr.db.Exec("DELETE FROM carts WHERE beverage_id = ? AND user_id = ?", beverageID, userID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cr *CartRepository) ClearCart(userID int) error {
	// Clear the cart for the specified user ID
	deleteQuery := `DELETE FROM carts WHERE user_id = ?`
	_, err := cr.db.Exec(deleteQuery, userID)
	return err
}
