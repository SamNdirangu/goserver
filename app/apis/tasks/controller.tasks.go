package tasks

import "github.com/gofiber/fiber/v2"

func GetAllTasks(fc *fiber.Ctx) error {
	return fc.Status(fiber.StatusOK).SendString("Task 1")
}

func GetTask(fc *fiber.Ctx) error {
	return fc.Status(fiber.StatusOK).SendString("Task 1")
}
func CreateTask(fc *fiber.Ctx) error {
	return fc.Status(fiber.StatusOK).SendString("Task 1")
}

func UpdateTask(fc *fiber.Ctx) error {
	return fc.Status(fiber.StatusOK).SendString("Task 1")
}

func DeleteTask(fc *fiber.Ctx) error {
	return fc.Status(fiber.StatusOK).SendString("Task 1")
}
