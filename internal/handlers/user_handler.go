package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/services"
)

type UserHandler struct {
	userService services.UserService
}

func InitUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

// POST /api/user/register (only used on first login if user doesn't exist)
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body")
	}
	err := h.userService.RegisterUser(&user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
	})
}

// GET /api/user/:phone
func (h *UserHandler) GetUserByPhone(c *fiber.Ctx) error {
	phone := c.Params("phone")
	user, err := h.userService.GetUser(phone)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}
	return c.JSON(user)
}

// PUT /api/user/update
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body")
	}
	err := h.userService.UpdateUser(&user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"message": "user updated successfully",
	})
}

// GET /api/user/me
func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	phone := c.Locals("userPhone").(string)
	user, err := h.userService.GetUser(phone)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}
	return c.JSON(user)
}
