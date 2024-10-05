package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "beverages-booking/services"
    "beverages-booking/models"
	"strconv"
)

type RatingController struct {
    ratingService *services.RatingService
}

func NewRatingController(ratingService *services.RatingService) *RatingController {
    return &RatingController{ratingService: ratingService}
}

func (rc *RatingController) AddRating(c *gin.Context) {
    var rating models.Rating
    if err := c.ShouldBindJSON(&rating); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := rc.ratingService.AddRating(rating); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add rating"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Rating added successfully"})
}

func (rc *RatingController) GetRatings(c *gin.Context) {
    beverageIDStr := c.Param("beverage_id")

	beverageId, err := strconv.Atoi(beverageIDStr)
	if (err != nil) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid beverage Id"})
		return
	}

    ratings, err := rc.ratingService.GetRatings(beverageId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch ratings"})
        return
    }

    c.JSON(http.StatusOK, ratings)
}

func (rc *RatingController) GetAllReviews(c *gin.Context) {
    reviews, err := rc.ratingService.GetAllReviews()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch reviews"})
        return
    }

    c.JSON(http.StatusOK, reviews)
}
