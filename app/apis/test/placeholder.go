package test

import "github.com/gofiber/fiber/v2"

func PlaceHolder(fc *fiber.Ctx) error {
	return fc.Status(fiber.StatusNotImplemented).SendString("API is not yet implemented this is placeholder response")
}
