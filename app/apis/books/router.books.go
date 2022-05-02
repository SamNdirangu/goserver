package books

import "github.com/gofiber/fiber/v2"

//Books API
func RegisterBookRoutes(api fiber.Router) {
	books := api.Group("/books")

	books.Post("/", CreateBook).Name("Create Book")
	books.Get("/", GetAllBooks).Name("GetAllBooks")
	books.Get("/:id", GetBookById).Name("GetBookByID")
	books.Patch("/:id", UpdateBookById).Name("Update Book")
	books.Delete("/:id", DeleteBookById).Name("Delete Book")
}
