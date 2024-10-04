package models

type Order struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	ItemName    string    `json:"item_name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	BeverageID  int       `json:"beverage_id"`
}