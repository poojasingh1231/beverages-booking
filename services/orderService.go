package services

import (
	"beverages-booking/models"
	"beverages-booking/repositories"
	"log"
)

type OrderService struct {
	orderRepository *repositories.OrderRepository
	cartRepository  *repositories.CartRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository, cartRepo *repositories.CartRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepo,
		cartRepository:  cartRepo,
	}
}

func (os *OrderService) PlaceOrder(userID int) error {
    items, err := os.cartRepository.GetCartItems(userID)
    if err != nil {
        return err
    }

    for _, item := range items {
        order := models.Order{
            UserID:      userID,
            ItemName:    item.ItemName,
            Description: item.Description,
            Price:       item.Price,
            Quantity:    item.Quantity,
            BeverageID:  item.BeverageID,
        }
        
        if err := os.orderRepository.PlaceOrder(order); err != nil {
            log.Printf("Error placing order for item %s: %v", item.ItemName, err)
            return err
        }
    }

    if err := os.cartRepository.ClearCart(userID); err != nil {
        log.Printf("Error clearing cart for user ID %d: %v", userID, err)
        return err
    }

    return nil
}


func (os *OrderService) GetOrderHistory(userID int) ([]models.Order, error) {
	orders, err := os.orderRepository.GetOrderHistory(userID)
	if err != nil {
		// Print the error to the log
		log.Printf("Error fetching order history for user ID %d: %v", userID, err)
		return nil, err
	}
	return orders, nil
}