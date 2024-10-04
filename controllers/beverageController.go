package controllers

import (
	"beverages-booking/services"
	"beverages-booking/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"log"
)

type BeverageController struct {
	beverageService *services.BeverageService
	adminService *services.AdminService
}

func NewBeverageController(beverageService *services.BeverageService, adminService *services.AdminService) *BeverageController {
	return &BeverageController{
		beverageService: beverageService,
		adminService : adminService,
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
	if (!bc.validateAdmin(c)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user"})
		return
	}
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
	if (!bc.validateAdmin(c)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user"})
		return
	}
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


func (bc BeverageController) validateAdmin(c *gin.Context) bool {

	log.Printf("Hello")
	userIdStr := c.Query("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return false
	}
	userName := c.Query("user_name")

	log.Printf("userid = %d, username= %s", userId, userName)
	return bc.adminService.AdminUserExists(userId, userName)
}