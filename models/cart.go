// models/cart.go
package models

type Cart struct {
	ID          int     `json:"id"`
	ItemName    string  `json:"item_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	UserID      int     `json:"user_id"`
	BeverageID  int     `json:"beverage_id"`
}
