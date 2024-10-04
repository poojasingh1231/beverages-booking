package controllers

import (
	"beverages-booking/services"
	"beverages-booking/context"
	"github.com/gin-gonic/gin"
	"net/http"
)
type AdminController struct {
	adminService *services.AdminService
}


func NewAdminController(adminService *services.AdminService) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

func (ac AdminController) AdminLogin(c *gin.Context) {
	if (context.IsLoggedIn) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Already logged in, logout first"})
		return
	}
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	admin, err := ac.adminService.AdminLogin(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "admin": admin})
}

func (ac AdminController) AdminLogout(ctx *gin.Context) {
	if (context.IsAdmin == false) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Logout attempt"})
		return
	}
	ac.adminService.AdminLogout()
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
