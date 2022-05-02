package test

import "github.com/gofiber/fiber/v2"

func HeartBeat(fc *fiber.Ctx) error {
	return fc.Status(fiber.StatusOK).SendString("App is alive")
}
