package controllers

import (
	"beverages_booking/config"
	"beverages_booking/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListBeverages(c *gin.Context) {
	beverages, err := repositories.GetAllBeverages(config.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve beverages"})
		return
	}
	c.JSON(http.StatusOK, beverages)
}

func CreateBeverage(c *gin.Context) {
	var beverage repositories.Beverage
	if err := c.ShouldBindJSON(&beverage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := repositories.CreateBeverage(config.DB, beverage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create beverage"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func DeleteBeverage(c *gin.Context) {
	id := c.Param("id")
	
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	if err := repositories.DeleteBeverage(config.DB, intID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete beverage"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Beverage deleted successfully"})
}
