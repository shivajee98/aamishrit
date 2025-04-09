package services

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/repository"
)

type ReviewService interface {
	AddReview(review *model.Review) error
	GetReviews(productID uint) ([]model.Review, error)
	UpdateReview(reviewID uint, updated *model.Review) error
	DeleteReview(reviewID uint) error
}

type reviewService struct {
	repo repository.ReviewRepository
}

func InitReviewService(repo repository.ReviewRepository) ReviewService {
	return &reviewService{repo: repo}
}

func (s *reviewService) AddReview(review *model.Review) error {
	return s.repo.CreateReview(review)
}

func (s *reviewService) GetReviews(productID uint) ([]model.Review, error) {
	return s.repo.GetReviewsByProductID(productID)
}

func (s *reviewService) UpdateReview(reviewID uint, updated *model.Review) error {
	return s.repo.UpdateReview(reviewID, updated)
}

func (s *reviewService) DeleteReview(reviewID uint) error {
	return s.repo.DeleteReview(reviewID)
}
