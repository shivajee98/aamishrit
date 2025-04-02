package repository

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *model.Order) error
	GetOrder(orderID uint) (*model.Order, error)
	GetOrdersByUser(userID uint) ([]model.Order, error)
	UpdateOrderStatus(orderID uint, status string) error
	DeleteOrder(orderID uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetOrder(orderID uint) (*model.Order, error) {
	var order model.Order
	err := r.db.Preload("Items").First(&order, orderID).Error
	return &order, err
}

func (r *orderRepository) GetOrdersByUser(userID uint) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Where("user_id = ?", userID).Preload("Items").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) UpdateOrderStatus(orderID uint, status string) error {
	return r.db.Model(&model.Order{}).Where("id = ?", orderID).Update("status", status).Error
}

func (r *orderRepository) DeleteOrder(orderID uint) error {
	return r.db.Delete(&model.Order{}, orderID).Error
}
