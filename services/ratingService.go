package services

import (
    "log"
    "beverages-booking/repositories"
    "beverages-booking/models"
)

type RatingService struct {
    ratingRepository *repositories.RatingRepository
}

func NewRatingService(ratingRepo *repositories.RatingRepository) *RatingService {
    return &RatingService{ratingRepository: ratingRepo}
}

func (rs *RatingService) AddRating(rating models.Rating) error {
    if err := rs.ratingRepository.AddRating(rating); err != nil {
        log.Printf("Error adding rating for user ID %d and beverage ID %d: %v", rating.UserID, rating.BeverageID, err)
        return err
    }
    log.Printf("Rating added successfully for user ID %d and beverage ID %d", rating.UserID, rating.BeverageID)
    return nil
}

func (rs *RatingService) GetRatings(beverageID int) ([]models.Rating, error) {
    ratings, err := rs.ratingRepository.GetRatingsByBeverage(beverageID)
    if err != nil {
        log.Printf("Error fetching ratings for beverage ID %d: %v", beverageID, err)
        return nil, err
    }
    log.Printf("Fetched %d ratings for beverage ID %d", len(ratings), beverageID)
    return ratings, nil
}

func (rs *RatingService) GetAllReviews() ([]models.Rating, error) {
    reviews, err := rs.ratingRepository.GetAllReviews()
    if err != nil {
        log.Printf("Error fetching all reviews: %v", err)
        return nil, err
    }
    log.Printf("Fetched %d reviews", len(reviews))
    return reviews, nil
}
