package repositories

import (
    "database/sql"
    "beverages-booking/models"
)

type RatingRepository struct {
    db *sql.DB
}

func NewRatingRepository(dbHandler *sql.DB) *RatingRepository {
    var repo = &RatingRepository{db: dbHandler}
	repo.CreateRatingTable()
	return repo
}

func (rr *RatingRepository) CreateRatingTable() error {
    createTableQuery := `CREATE TABLE IF NOT EXISTS ratings (
        id INT AUTO_INCREMENT PRIMARY KEY,
        user_id INT NOT NULL,
        beverage_id INT NOT NULL,
        score FLOAT NOT NULL CHECK (score >= 1.0 AND score <= 5.0),
        comment TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`
    _, err := rr.db.Exec(createTableQuery)
    return err
}

func (rr *RatingRepository) AddRating(rating models.Rating) error {
    query := `INSERT INTO ratings (user_id, beverage_id, score, comment) 
              VALUES (?, ?, ?, ?)`
    _, err := rr.db.Exec(query, rating.UserID, rating.BeverageID, rating.Score, rating.Comment)
    return err
}

func (rr *RatingRepository) GetRatingsByBeverage(beverageID int) ([]models.Rating, error) {
    rows, err := rr.db.Query("SELECT id, user_id, beverage_id, score, comment FROM ratings WHERE beverage_id = ?", beverageID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var ratings []models.Rating
    for rows.Next() {
        var rating models.Rating
        if err := rows.Scan(&rating.ID, &rating.UserID, &rating.BeverageID, &rating.Score, &rating.Comment); err != nil {
            return nil, err
        }
        ratings = append(ratings, rating)
    }

    return ratings, nil
}

func (rr *RatingRepository) GetAllReviews() ([]models.Rating, error) {
    rows, err := rr.db.Query("SELECT id, user_id, beverage_id, score, comment FROM ratings")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var reviews []models.Rating
    for rows.Next() {
        var review models.Rating
        if err := rows.Scan(&review.ID, &review.UserID, &review.BeverageID, &review.Score, &review.Comment); err != nil {
            return nil, err
        }
        reviews = append(reviews, review)
    }

    return reviews, nil
}
