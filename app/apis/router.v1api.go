package apis

import (
	"goserver/app/apis/books"
	"goserver/app/apis/products"
	"goserver/app/apis/users"

	"github.com/gofiber/fiber/v2"
)

func RegisterV1APIs(api fiber.Router) {
	products.RegisterProductRoutes(api)
	books.RegisterBookRoutes(api)
	users.RegisterUserRoutes(api)
}
