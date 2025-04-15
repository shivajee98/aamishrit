package services

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/repository"
)

type CategoryService interface {
	CreateCategory(category *model.Category) error
	GetAllCategories() ([]model.Category, error)
	GetCategoryByID(id uint) (*model.Category, error)
	UpdateCategory(category *model.Category) error
	DeleteCategory(id uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func InitCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(category *model.Category) error {
	return s.repo.CreateCategory(category)
}

func (s *categoryService) GetAllCategories() ([]model.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *categoryService) GetCategoryByID(id uint) (*model.Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *categoryService) UpdateCategory(category *model.Category) error {
	return s.repo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repo.DeleteCategory(id)
}
