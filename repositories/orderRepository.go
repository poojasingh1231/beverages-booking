package repositories

import (
	"database/sql"
	"beverages-booking/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(dbHandler *sql.DB) *OrderRepository {
	var repo = &OrderRepository{
		db: dbHandler,
	}

	repo.CreateOrderTable()
	return repo
}

func (or *OrderRepository) CreateOrderTable() error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS orders (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		item_name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		price DECIMAL(10,2) NOT NULL,
		quantity INT NOT NULL,
		beverage_id INT NOT NULL
	)`
	_, err := or.db.Exec(createTableQuery)
	if err != nil {
		return err
	}
	return nil
}

func (or *OrderRepository) PlaceOrder(order models.Order) error {
    query := `INSERT INTO orders (user_id, item_name, description, price, quantity, beverage_id) 
              VALUES (?, ?, ?, ?, ?, ?)`

    _, err := or.db.Exec(query, order.UserID, order.ItemName, order.Description, order.Price, order.Quantity, order.BeverageID)
    if err != nil {
        return err
    }
    return nil
}

func (ar *OrderRepository) GetOrderHistory(userID int) ([]models.Order, error) {
    rows, err := ar.db.Query("SELECT id, user_id, item_name, description, price, quantity, beverage_id FROM orders WHERE user_id = ?", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var orders []models.Order
    for rows.Next() {
        var order models.Order
        err := rows.Scan(&order.ID, &order.UserID, &order.ItemName, &order.Description, &order.Price, &order.Quantity, &order.BeverageID)
        if err != nil {
            return nil, err
        }
        orders = append(orders, order)
    }
    return orders, nil
}