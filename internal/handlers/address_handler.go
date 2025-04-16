package handlers

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/services"
	"gorm.io/gorm"
)

type AddressHandler struct {
	service     services.AddressService
	userService services.UserService
}

func InitAddressHandler(s services.AddressService, u services.UserService) *AddressHandler {
	return &AddressHandler{service: s, userService: u}
}

// POST /api/address
func (h *AddressHandler) CreateAddress(c *fiber.Ctx) error {
	var address model.Address
	if err := c.BodyParser(&address); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid address format")
	}

	// Extract Clerk ID from context
	clerkIDValue := c.Locals("clerk_id")
	clerkID, ok := clerkIDValue.(string)
	if !ok || clerkID == "" {
		log.Println("RegisterUser: missing or invalid Clerk ID")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	// Get the userId by clerk Id
	user, err := h.userService.GetUserByClerkID(clerkID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("RegisterUser: error checking user existence: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	}

	userID := user.ID
	address.UserID = userID

	if err := h.service.CreateAddress(&address); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"message": "Address created"})
}

// GET /api/address
func (h *AddressHandler) GetAllAddresses(c *fiber.Ctx) error {
	// Extract Clerk ID from context
	clerkIDValue := c.Locals("clerk_id")
	clerkID, ok := clerkIDValue.(string)
	if !ok || clerkID == "" {
		log.Println("RegisterUser: missing or invalid Clerk ID")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	addresses, err := h.service.GetAddressesByUserID(clerkID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(addresses)
}

// GET /api/address/:id
func (h *AddressHandler) GetAddressByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid address ID")
	}

	address, err := h.service.GetAddressByID(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Address not found")
	}

	return c.JSON(address)
}

// PUT /api/address/:id
func (h *AddressHandler) UpdateAddress(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid address ID")
	}

	var address model.Address
	if err := c.BodyParser(&address); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid address format")
	}

	address.ID = uint(id)
	if err := h.service.UpdateAddress(&address); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"message": "Address updated"})
}

// DELETE /api/address/:id
func (h *AddressHandler) DeleteAddress(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid address ID")
	}

	if err := h.service.DeleteAddress(uint(id)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"message": "Address deleted"})
}

// PUT /api/address/:id/default
func (h *AddressHandler) SetDefaultAddress(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid address ID")
	}

	userID := c.Locals("user_id").(uint)
	if err := h.service.SetDefaultAddress(userID, uint(id)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"message": "Default address updated"})
}

// GET /api/address/default
func (h *AddressHandler) GetDefaultAddress(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	address, err := h.service.GetDefaultAddress(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Default address not set")
	}

	return c.JSON(address)
}
