package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/services"
)

type PaymentHandler struct {
	paymentService *services.PaymentService
}

func NewPaymentHandler(paymentService *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) CreateOrder(c *fiber.Ctx) error {
	var request struct {
		Amount  float64 `json:"amount"`
		UserID  uint    `json:"user_id"`
		OrderID uint    `json:"order_id"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	payment, err := h.paymentService.CreateOrder(request.Amount, request.UserID, request.OrderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(payment)
}

func (h *PaymentHandler) VerifyPayment(c *fiber.Ctx) error {
	transactionID := c.Params("transaction_id")
	razorpayOrderID := c.Params("order_id")

	if transactionID == "" || razorpayOrderID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing parameters"})
	}

	err := h.paymentService.VerifyPayment(transactionID, razorpayOrderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Payment verified successfully"})
}
