// controllers/cartController.go
package controllers

import (
	"beverages-booking/services"
	"beverages-booking/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CartController struct {
	cartService *services.CartService
}

func NewCartController(cartService *services.CartService) *CartController {
	return &CartController{
		cartService: cartService,
	}
}

func (cc CartController) GetCartItems(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	cartItems, err := cc.cartService.GetCartItemsService(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve cart items"})
		return
	}
	c.JSON(http.StatusOK, cartItems)
}


func (cc CartController) AddItem(c *gin.Context) {
	var cartItem models.Cart

	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := cc.cartService.AddItemService(cartItem.UserID, &cartItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add item to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart successfully"})
}

// RemoveItem removes a beverage item from the cart.
func (cc CartController) RemoveItem(c *gin.Context) {
	userIdStr := c.Query("user_id")
	beverageIdStr := c.Query("beverage_id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	beverageId, err := strconv.Atoi(beverageIdStr)
	if (err != nil) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid beverage Id"})
		return
	}
	
	if err := cc.cartService.RemoveItemService(userId, beverageId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not remove item from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart successfully"})
}
