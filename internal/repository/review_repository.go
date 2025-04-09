package repository

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	CreateReview(review *model.Review) error
	GetReviewsByProductID(productID uint) ([]model.Review, error)
	UpdateReview(reviewID uint, updated *model.Review) error
	DeleteReview(reviewID uint) error
}

type reviewRepository struct {
	db *gorm.DB
}

func InitReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) CreateReview(review *model.Review) error {
	return r.db.Create(review).Error
}

func (r *reviewRepository) GetReviewsByProductID(productID uint) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Where("product_id = ?", productID).Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) UpdateReview(reviewID uint, updated *model.Review) error {
	return r.db.Model(&model.Review{}).Where("id = ?", reviewID).Updates(updated).Error
}

func (r *reviewRepository) DeleteReview(reviewID uint) error {
	return r.db.Delete(&model.Review{}, reviewID).Error
}
