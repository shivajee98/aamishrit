package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/services"
)

type ProductHandler struct {
	productService services.ProductService
}

func InitProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "Invalid Product ID"})
	}

	product, err := h.productService.GetProductsByID(uint(id))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product Not Found"})
	}

	return c.JSON(product)
}

func (h *ProductHandler) ListProducts(c *fiber.Ctx) error {
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	product, err := h.productService.ListProducts(offset, limit)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(product)

}
