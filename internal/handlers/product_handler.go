package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/services"
	"github.com/shivajee98/aamishrit/internal/uploader"
	"github.com/shivajee98/aamishrit/pkg/utils"
)

type ProductHandler struct {
	productService     services.ProductService
	cloudinaryUploader *uploader.CloudinaryUploader
}

func InitProductHandler(productService services.ProductService, uploader *uploader.CloudinaryUploader) *ProductHandler {
	return &ProductHandler{
		productService:     productService,
		cloudinaryUploader: uploader,
	}
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

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Form Data"})
	}

	product := &model.Product{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		Price:       utils.ParseFloat(c.FormValue("price")),
		Categories:  c.FormValue("category"),
		Images:      []string{},
	}

	// Upload Images
	files := form.File["productImages"]
	if files == nil || len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No product images provided"})
	}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open uploaded file"})
		}
		defer src.Close()

		imageURL, err := h.cloudinaryUploader.Upload(src)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Image Upload Failed"})
		}

		product.Images = append(product.Images, imageURL)
	}

	// Save the product
	if err := h.productService.CreateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Product ID"})
	}

	// Fetch existing product
	existingProduct, err := h.productService.GetProductsByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product Not Found"})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Form Data"})
	}

	// Update only if field is provided
	name := c.FormValue("name")
	if name != "" {
		existingProduct.Name = name
	}
	description := c.FormValue("description")
	if description != "" {
		existingProduct.Description = description
	}
	priceStr := c.FormValue("price")
	if priceStr != "" {
		existingProduct.Price = utils.ParseFloat(priceStr)
	}
	category := c.FormValue("category")
	if category != "" {
		existingProduct.Categories = category
	}

	// Handle optional image update
	files := form.File["productImages"]
	if files != nil && len(files) > 0 {
		// clear old images if required, or append (your business logic)
		existingProduct.Images = []string{} // or append to existing

		for _, file := range files {
			src, err := file.Open()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open uploaded file"})
			}
			defer src.Close()

			imageURL, err := h.cloudinaryUploader.Upload(src)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Image Upload Failed"})
			}

			existingProduct.Images = append(existingProduct.Images, imageURL)
		}
	}

	if err := h.productService.UpdateProduct(existingProduct); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product"})
	}

	return c.JSON(existingProduct)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	err = h.productService.DeleteProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
