package controllers

import "github.com/gofiber/fiber/v2"

func GetAllJobs(fc *fiber.Ctx) error {
	return fc.Status(fiber.StatusOK).SendString("Many Jobs")
}

func GetJob(fc *fiber.Ctx) error {
	return fc.Status(fiber.StatusOK).SendString("Job 1")
}
func CreateJob(fc *fiber.Ctx) error {
	return fc.Status(fiber.StatusOK).SendString("Create a Job")
}

func UpdateJob(fc *fiber.Ctx) error {
	return fc.Status(fiber.StatusOK).SendString("Update a Job")
}

func DeleteJob(fc *fiber.Ctx) error {
	return fc.Status(fiber.StatusOK).SendString("Delete a Job")
}
