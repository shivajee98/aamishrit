package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/services"
)

type CategoryHandler struct {
	service services.CategoryService
}

func InitCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var category model.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.service.CreateCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(category)
}

func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(categories)
}

func (h *CategoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	category, err := h.service.GetCategoryByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
	}
	return c.JSON(category)
}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var updated model.Category
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	updated.ID = uint(id)
	if err := h.service.UpdateCategory(&updated); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(updated)
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.service.DeleteCategory(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
