package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/services"
	"github.com/shivajee98/aamishrit/pkg/utils"
)

type UserHandler struct {
	userService services.UserService
}

type UserPhoneNumber struct {
	PhoneNumber string `json:"phoneNumber"`
}

func InitUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// POST /api/user/register (only used on first login if user doesn't exist)
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	// Accessing Clerk User Id From c.Locals()
	userID := c.Locals("user_id")

	if userID == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User ID not found")
	}

	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body")
	}
	user.UserID = userID.(string)

	// Fetch Phone Number
	userData, err := utils.FetchClerkUser(user.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch user from Clerk")
	}

	// Extract the phone number
	if phone, ok := userData["phone"]; ok {
		user.Phone = phone.(string)
	}
	err = h.userService.RegisterUser(&user)
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
	user, err := h.userService.GetUserByPhone(phone)
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
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var request UserPhoneNumber
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Resquest Body"})
	}

	user, err := h.userService.GetUserByPhone(request.PhoneNumber)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User Not Found"})
	}

	return c.JSON(user)
}
