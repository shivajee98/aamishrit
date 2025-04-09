package handlers

import (
	"errors"
	"html"
	"log"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gofiber/fiber/v2"
	mw "github.com/shivajee98/aamishrit/internal/middleware"
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/services"
	"github.com/shivajee98/aamishrit/pkg/utils"
	"gorm.io/gorm"
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

// POST /api/user/register
// RegisterUser handles POST /api/user/register
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	// Extract Clerk ID from context
	clerkIDValue := c.Locals(mw.UserIDKey)
	clerkID, ok := clerkIDValue.(string)
	if !ok || clerkID == "" {
		log.Println("RegisterUser: missing or invalid Clerk ID")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	// Check if user already exists
	existingUser, err := h.userService.GetUserByClerkID(clerkID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("RegisterUser: error checking user existence: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	}
	if existingUser != nil {
		return fiber.NewError(fiber.StatusConflict, "User already exists")
	}

	// Parse request body
	var userModel model.User
	if err := c.BodyParser(&userModel); err != nil {
		log.Printf("RegisterUser: body parse error: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Input validation
	userModel.Name = sanitizeString(userModel.Name)
	if userModel.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Name is required")
	}

	// Fetch user details from Clerk
	ctx := c.Context()
	userDetails, err := user.Get(ctx, clerkID)
	if err != nil {
		log.Printf("RegisterUser: failed to fetch Clerk user details for %s: %v", clerkID, err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch user details")
	}

	// Extract phone number
	var phoneNumber string
	if len(userDetails.PhoneNumbers) > 0 {
		phoneNumber = userDetails.PhoneNumbers[0].PhoneNumber
	} else {
		log.Printf("RegisterUser: no phone number found for Clerk user: %s", clerkID)
		return fiber.NewError(fiber.StatusBadRequest, "Phone number is missing from Clerk profile")
	}

	userModel.ClerkID = clerkID
	userModel.Phone = phoneNumber

	// Save user
	if err := h.userService.RegisterUser(&userModel); err != nil {
		log.Printf("RegisterUser: DB error while registering user: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create user")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user": fiber.Map{
			"id":    userModel.ID,
			"name":  userModel.Name,
			"phone": userModel.Phone,
		},
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
	ClerkID := c.Locals("user_id")

	if ClerkID == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User ID not found")
	}

	var user model.User
	userData, err := utils.FetchClerkUser(ClerkID.(string))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch user from Clerk")
	}

	userExists, err := h.userService.GetUserByPhone(userData.PhoneNumber)

	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body")
	}

	if userExists == nil {
		c.Locals("user_id", nil)

		return c.Redirect("/register?error=user_not_found")
	}

	err = h.userService.UpdateUser(&user)
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

func (h *UserHandler) GetClerkUser(c *fiber.Ctx) error {
	ClerkID := c.Locals("clerk_id")
	clerkID, ok := ClerkID.(string)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid Clerk ID")
	}

	userData, err := utils.FetchClerkUser(clerkID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch user from Clerk")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": userData})

}

func sanitizeString(input string) string {
	return strings.TrimSpace(html.EscapeString(input))
}
