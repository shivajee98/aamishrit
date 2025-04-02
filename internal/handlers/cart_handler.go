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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	err := h.cartService.AddToCart(&cart)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "added to cart"})
}

func (h *CartHandler) GetCart(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Query(("user_id")))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	cart, err := h.cartService.GetCartByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(cart)
}

func (h *CartHandler) UpdateCartItem(c *fiber.Ctx) error {
	cartID, err := strconv.Atoi(c.Params("cart_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid cart ID"})
	}

	var cart model.Cart
	if err := c.BodyParser(&cart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	// Pass cartID from URL param, not from request body
	err = h.cartService.UpdateCartItem(uint(cartID), &cart)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Cart updated successfully"})
}

func (h *CartHandler) RemoveFromCart(c *fiber.Ctx) error {
	cartID, err := strconv.Atoi(c.Query("cart_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid cart id"})
	}

	err = h.cartService.RemoveFromCart(uint(cartID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Item removed from cart"})
}

func (h *CartHandler) ClearCart(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user Id"})
	}
	err = h.cartService.ClearCart(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Cart Cleared"})
}
