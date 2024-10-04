package controllers

import (
	"beverages-booking/services"
	"beverages-booking/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BeverageController struct {
	beverageService *services.BeverageService
}

func NewBeverageController(beverageService *services.BeverageService) *BeverageController {
	return &BeverageController{
		beverageService: beverageService,
	}
}

func (bc BeverageController) ListBeverages(c *gin.Context) {
	beverages, err := bc.beverageService.GetAllBeveragesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve beverages"})
		return
	}
	c.JSON(http.StatusOK, beverages)
}


func (bc BeverageController) CreateBeverage(c *gin.Context) {
	var beverage =  new(models.Beverage)
	if err := c.ShouldBindJSON(&beverage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := bc.beverageService.CreateBeverageService(beverage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create beverage"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}


func (bc BeverageController) DeleteBeverage(c *gin.Context) {
	id := c.Param("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	if err := bc.beverageService.DeleteBeverageService(intID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete beverage"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Beverage deleted successfully"})
}
