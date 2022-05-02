package errorHandlers

import "github.com/gofiber/fiber/v2"

func NotFoundHandler(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendString("Ooops 404 you lost")
}
