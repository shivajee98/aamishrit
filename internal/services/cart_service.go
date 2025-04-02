package services

import (
	"errors"

	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/repository"
)

type CartService interface {
	AddToCart(cart *model.Cart) error
	GetCartByUserID(userID uint) ([]model.Cart, error)
	UpdateCardItem(cart *model.Cart) error
	RemoveFromCart(cardID uint) error
	ClearCart(userID uint) error
}

type cartService struct {
	cartRepo repository.CartRepository
}

func InitCartService(cartRepo repository.CartRepository) CartService {
	return &cartService{cartRepo: cartRepo}
}

// Add to cart
func (s *cartService) AddToCart(cart *model.Cart) error {
	if cart.Quantity <= 0 {
		return errors.New("Invalid Quantity")
	}
	return s.cartRepo.AddToCart(cart)
}

// Get Cart By User ID
func (s *cartService) GetCartByUserID(userID uint) ([]model.Cart, error) {
	return s.cartRepo.GetCartByUserID(userID)
}

// Update Cart Item
func (s *cartService) UpdateCardItem(cart *model.Cart) error {
	if cart.ID == 0 {
		return errors.New("invalid cart id")
	}
	return s.cartRepo.UpdateCardItem(cart)
}

// Remove from Cart
func (s *cartService) RemoveFromCart(cardID uint) error {
	return s.cartRepo.RemoveFromCart(cardID)
}

// Clear Cart
func (s *cartService) ClearCart(userID uint) error {
	return s.cartRepo.ClearCart(userID)
}
