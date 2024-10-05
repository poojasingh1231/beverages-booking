package models

type Rating struct {
    ID         int     `json:"id"`
    UserID     int     `json:"user_id"`
    BeverageID int     `json:"beverage_id"`
    Score      float64 `json:"score"`
    Comment    string  `json:"comment"`
}
