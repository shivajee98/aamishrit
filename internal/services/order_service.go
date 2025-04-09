package services

import (
	"errors"
	"fmt"

	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/repository"
)

type OrderService interface {
	PlaceOrder(clerkID string, order *model.Order) error
	GetOrder(orderID uint) (*model.Order, error)
	GetUserOrders(userID uint) ([]model.Order, error)
	UpdateOrderStatus(orderID uint, status string) error
	CancelOrder(orderID uint) error
}

type orderService struct {
	orderRepo repository.OrderRepository
	userRepo  repository.UserRepository
}

func NewOrderService(orderRepo repository.OrderRepository, userRepo repository.UserRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (s *orderService) PlaceOrder(clerkID string, order *model.Order) error {
	user, err := s.userRepo.GetUserByClerkID(clerkID)
	if err != nil || user == nil {
		return fmt.Errorf("user not found for clerkID: %s", clerkID)
	}

	order.UserID = user.ID

	return s.orderRepo.CreateOrder(order)
}

func (s *orderService) GetOrder(orderID uint) (*model.Order, error) {
	return s.orderRepo.GetOrder(orderID)
}

func (s *orderService) GetUserOrders(userID uint) ([]model.Order, error) {
	return s.orderRepo.GetOrdersByUser(userID)
}

func (s *orderService) UpdateOrderStatus(orderID uint, status string) error {
	validStatuses := map[string]bool{
		"pending":   true,
		"shipped":   true,
		"delivered": true,
		"cancelled": true,
	}

	if !validStatuses[status] {
		return errors.New("invalid order status")
	}

	return s.orderRepo.UpdateOrderStatus(orderID, status)
}

func (s *orderService) CancelOrder(orderID uint) error {
	order, err := s.orderRepo.GetOrder(orderID)
	if err != nil {
		return err
	}

	if order.Status == "delivered" {
		return errors.New("cannot cancel a delivered order")
	}

	return s.orderRepo.UpdateOrderStatus(orderID, "cancelled")
}
