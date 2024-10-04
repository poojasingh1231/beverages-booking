// services/cartService.go
package services

import (
	"beverages-booking/repositories"
	"beverages-booking/models"
)

type CartService struct {
	cartRepository *repositories.CartRepository
}

func NewCartService(cartRepository *repositories.CartRepository) *CartService {
	return &CartService{
		cartRepository: cartRepository,
	}
}

// GetCartItemsService retrieves all cart items for a specific user.
func (cs CartService) GetCartItemsService(userID int) ([]*models.Cart, error) {
	cartItems, err := cs.cartRepository.GetCartItems(userID)
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}

// AddItemService adds a beverage item to the cart for a specific user.
func (cs CartService) AddItemService(userID int, cart *models.Cart) error {
	err := cs.cartRepository.AddItem(userID, cart)
	if err != nil {
		return err
	}
	return nil
}

// RemoveItemService removes a beverage item from the cart for a specific user.
func (cs CartService) RemoveItemService(userID int, beverageID int) error {
	err := cs.cartRepository.RemoveItem(userID, beverageID)
	if err != nil {
		return err
	}
	return nil
}
