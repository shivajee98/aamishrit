package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/services"
)

type ReviewHandler struct {
	service services.ReviewService
}

func InitReviewHandler(service services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

func (h *ReviewHandler) AddReview(c *fiber.Ctx) error {
	var review model.Review
	if err := c.BodyParser(&review); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if review.Rating < 1 || review.Rating > 5 {
		return fiber.NewError(fiber.StatusBadRequest, "Rating must be between 1 and 5")
	}

	if err := h.service.AddReview(&review); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Review added successfully"})
}

func (h *ReviewHandler) GetReviews(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid product ID")
	}

	reviews, err := h.service.GetReviews(uint(productID))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(reviews)
}

func (h *ReviewHandler) UpdateReview(c *fiber.Ctx) error {
	reviewID, err := strconv.Atoi(c.Params("review_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid review ID")
	}

	var updated model.Review
	if err := c.BodyParser(&updated); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if updated.Rating < 1 || updated.Rating > 5 {
		return fiber.NewError(fiber.StatusBadRequest, "Rating must be between 1 and 5")
	}

	err = h.service.UpdateReview(uint(reviewID), &updated)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"message": "Review updated successfully"})
}

func (h *ReviewHandler) DeleteReview(c *fiber.Ctx) error {
	reviewID, err := strconv.Atoi(c.Params("review_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid review ID")
	}

	err = h.service.DeleteReview(uint(reviewID))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"message": "Review deleted successfully"})
}
