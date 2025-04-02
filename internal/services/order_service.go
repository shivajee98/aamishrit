package services

import (
	"errors"

	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/repository"
)

type OrderService interface {
	PlaceOrder(order *model.Order) error
	GetOrder(orderID uint) (*model.Order, error)
	GetUserOrders(userID uint) ([]model.Order, error)
	UpdateOrderStatus(orderID uint, status string) error
	CancelOrder(orderID uint) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) PlaceOrder(order *model.Order) error {
	if order.TotalAmount <= 0 {
		return errors.New("total amount must be greater than zero")
	}

	if len(order.Items) == 0 {
		return errors.New("order must contain at least one item")
	}

	return s.repo.CreateOrder(order)
}

func (s *orderService) GetOrder(orderID uint) (*model.Order, error) {
	return s.repo.GetOrder(orderID)
}

func (s *orderService) GetUserOrders(userID uint) ([]model.Order, error) {
	return s.repo.GetOrdersByUser(userID)
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

	return s.repo.UpdateOrderStatus(orderID, status)
}

func (s *orderService) CancelOrder(orderID uint) error {
	order, err := s.repo.GetOrder(orderID)
	if err != nil {
		return err
	}

	if order.Status == "delivered" {
		return errors.New("cannot cancel a delivered order")
	}

	return s.repo.UpdateOrderStatus(orderID, "cancelled")
}
