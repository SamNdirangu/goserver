package products

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(api fiber.Router) {
	products := api.Group("/products")

	products.Get("/", GetAllProducts).Name("GetAllProducts")
	//products.Get("/:id", bookController.GetBookById).Name("GetBookByID")
	products.Post("/", CreateProduct).Name("Create Product")
	//products.Patch("/:id", bookController.UpdateBookById).Name("Update Book")
	//products.Delete("/:id", bookController.DeleteBookById).Name("Delete Book")
}
