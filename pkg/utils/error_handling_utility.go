package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
)

func CheckError(messege string, err error) {
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(os.Stderr, "Error %s: %v\n", messege, err)
		os.Exit(1)
	}
}

func FiberError(c *fiber.Ctx, msg string, err error) error {
	log.Println(msg+":", err)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": msg})
}
