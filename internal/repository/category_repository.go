package repository

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category *model.Category) error
	GetAllCategories() ([]model.Category, error)
	GetCategoryByID(id uint) (*model.Category, error)
	UpdateCategory(category *model.Category) error
	DeleteCategory(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func InitCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetCategoryByID(id uint) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *categoryRepository) UpdateCategory(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(id uint) error {
	return r.db.Delete(&model.Category{}, id).Error
}
