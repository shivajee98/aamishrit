package services

import (
	"errors"

	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/repository"
)

type ReviewService interface {
	AddReview(review *model.Review) error
	GetReviews(productID uint) ([]model.Review, error)
	UpdateReview(reviewID uint, updatedReview *model.Review) error
	DeleteReview(reviewID uint) error
}

type reviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) ReviewService {
	return &reviewService{repo: repo}
}

func (s *reviewService) AddReview(review *model.Review) error {
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	return s.repo.AddReview(review)
}

func (s *reviewService) GetReviews(productID uint) ([]model.Review, error) {
	return s.repo.GetReviewsByProduct(productID)
}

func (s *reviewService) UpdateReview(reviewID uint, updatedReview *model.Review) error {
	existingReviews, err := s.repo.GetReviewsByProduct(updatedReview.ProductID)
	if err != nil {
		return err
	}

	found := false
	for _, r := range existingReviews {
		if r.ID == reviewID {
			found = true
			break
		}
	}

	if !found {
		return errors.New("review not found")
	}

	return s.repo.UpdateReview(updatedReview)
}

func (s *reviewService) DeleteReview(reviewID uint) error {
	return s.repo.DeleteReview(reviewID)
}
