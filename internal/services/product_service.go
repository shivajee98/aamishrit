package services

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/repository"
)

type ProductService interface {
	GetProductsByID(id uint) (*model.Product, error)
	ListProducts(offset, limit int) ([]model.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func InitProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

// GetProductsByID implements ProductService.
func (s *productService) GetProductsByID(id uint) (*model.Product, error) {
	return s.productRepo.GetProductByID(id)
}

// ListProducts implements ProductService.
func (s *productService) ListProducts(offset int, limit int) ([]model.Product, error) {
	return s.productRepo.ListProducts(offset, limit)
}
