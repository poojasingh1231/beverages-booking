package controllers

import (
	"net/http"
	"strconv"

	"beverages-booking/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

// Order handles placing a new order.
func (oc *OrderController) Order(c *gin.Context) {
	var request struct {
		UserID int `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Place the order
	if err := oc.orderService.PlaceOrder(request.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not place order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully"})
}

// ShowHistory retrieves the order history for a user.
func (oc *OrderController) ShowHistory(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	orders, err := oc.orderService.GetOrderHistory(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve order history"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
