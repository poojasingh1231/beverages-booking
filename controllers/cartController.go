// controllers/cartController.go
package controllers

import (
	"beverages-booking/services"
	"beverages-booking/models"
	"beverages-booking/context"
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

// GetCartItems retrieves all cart items for the logged-in user.
func (cc CartController) GetCartItems(c *gin.Context) {
	userID := context.UserID // Assuming context holds the logged-in user's ID
	cartItems, err := cc.cartService.GetCartItemsService(userID)
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
