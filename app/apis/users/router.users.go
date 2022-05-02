package users

import (
	"goserver/app/apis/test"

	"github.com/gofiber/fiber/v2"
)

//Account Management API
func RegisterUserRoutes(api fiber.Router) {
	account := api.Group("/account").Name("User Account Management")

	account.Get("/", test.PlaceHolder).Name("Get Account Info")
	account.Post("/signup", test.PlaceHolder).Name("Create account")
	account.Post("/signin", test.PlaceHolder).Name("Account SignIn")
	account.Post("/signout", test.PlaceHolder).Name("Account SignOut")
}
