package repository

import (
	"github.com/shivajee98/aamishrit/internal/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(id uint) (*model.Product, error)
	ListProducts(offset, limit int) ([]model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func InitProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error

	if err != nil {
		return nil, err
	}
	return &product, nil

}

func (r *productRepository) ListProducts(offset, limit int) ([]model.Product, error) {
	var products []model.Product
	err := r.db.Offset(offset).Limit(limit).Find(&products).Error
	return products, err
}

func (r *productRepository) CreateProduct(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) UpdateProduct(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}
