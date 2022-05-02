package tasks

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterTaskRoutes(api fiber.Router) {
	tasks := api.Group("/tasks")

	tasks.Get("/", GetAllTasks)
	tasks.Get("/:id", GetTask)
	tasks.Post("/:id", CreateTask)
	tasks.Patch("/:id", UpdateTask)
	tasks.Delete("/:id", DeleteTask)
}
