package repository

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"gorm.io/gorm"
)

type CartRepository interface {
	AddToCart(cart *model.Cart) error
	GetCartByUserID(userID uint) ([]model.Cart, error)
	UpdateCardItem(cart *model.Cart) error
	RemoveFromCart(cartID uint) error
	ClearCart(userID uint) error
}

type cartRepository struct {
	db *gorm.DB
}

func InitCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) AddToCart(cart *model.Cart) error {
	return r.db.Create(cart).Error
}

func (r *cartRepository) GetCartByUserID(userID uint) ([]model.Cart, error) {
	var cart []model.Cart
	err := r.db.Where("user_id = ?", userID).Find(&cart).Error
	return cart, err
}

func (r *cartRepository) UpdateCardItem(cart *model.Cart) error {
	return r.db.Save(cart).Error
}

func (r *cartRepository) RemoveFromCart(cartID uint) error {
	return r.db.Delete(&model.Cart{}, cartID).Error
}

func (r *cartRepository) ClearCart(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.Cart{}).Error
}
