package repository

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	AddReview(review *model.Review) error
	GetReviewsByProduct(productID uint) ([]model.Review, error)
	UpdateReview(review *model.Review) error
	DeleteReview(reviewID uint) error
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) AddReview(review *model.Review) error {
	return r.db.Create(review).Error
}

func (r *reviewRepository) GetReviewsByProduct(productID uint) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Where("product_id = ?", productID).Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) UpdateReview(review *model.Review) error {
	return r.db.Save(review).Error
}

func (r *reviewRepository) DeleteReview(reviewID uint) error {
	return r.db.Delete(&model.Review{}, reviewID).Error
}
