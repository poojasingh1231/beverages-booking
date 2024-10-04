package controllers

import (
	"beverages-booking/models"
	"beverages-booking/context"
	"beverages-booking/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"io"
	"log"
	"encoding/json"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc UserController) UserLogin(c *gin.Context) {
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

	user, err := uc.userService.UserLogin(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

func (uc UserController) UserLogout(ctx *gin.Context) {
	if (context.IsAdmin) {
		ctx.JSON(http.Unauthorized, gin.H{"message": "Invalid Logout attempt"})
		return
	}
	uc.userService.UserLogout()
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}



func (uc UserController) CreateUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create user request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling create user request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := uc.userService.CreateUser(&user)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}


