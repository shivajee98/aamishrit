package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/services"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) PlaceOrder(c *fiber.Ctx) error {
	var order model.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := h.orderService.PlaceOrder(&order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Order placed successfully"})
}

func (h *OrderHandler) GetOrder(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("order_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
	}

	order, err := h.orderService.GetOrder(uint(orderID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(order)
}

func (h *OrderHandler) GetUserOrders(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	orders, err := h.orderService.GetUserOrders(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(orders)
}

func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("order_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
	}

	var data struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err = h.orderService.UpdateOrderStatus(uint(orderID), data.Status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Order status updated successfully"})
}

func (h *OrderHandler) CancelOrder(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("order_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
	}

	err = h.orderService.CancelOrder(uint(orderID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Order cancelled successfully"})
}
