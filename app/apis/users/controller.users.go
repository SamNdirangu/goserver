package users

import "github.com/gofiber/fiber/v2"

type UserControllers struct{}
type IUserControllers interface {
	GetAllUsers(fc *fiber.Ctx) error
}

func (user *UserControllers) GetAllUsers(fc *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
