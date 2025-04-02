package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/services"
)

type CartHandler struct {
	cartService services.CartService
}

func InitCartHandler(cartService services.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

func (h *CartHandler) AddToCart(c *fiber.Ctx) error {
	var cart model.Cart
	if err := c.BodyParser(&cart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid error"})
	}

	err := h.cartService.AddToCart(&cart)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "added to cart"})
}

func (h *CartHandler) GetCart(c *fiber.Ctx) error {
	userID, err := strconv.Atoi("user_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	cart, err := h.cartService.GetCartByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(cart)
}

func (h * CartHandler) UpdateCardItem(c *fiber.Ctx) error {
	
}