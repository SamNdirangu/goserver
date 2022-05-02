package books

import (
	"goserver/app/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BookControllers struct{}
type IBookControllers interface {
	CreateBook(ctx *fiber.Ctx) error
	GetAllBooks(ctx *fiber.Ctx) error
	GetBookById(ctx *fiber.Ctx) error
	UpdateBookById(fc *fiber.Ctx) error
	DeleteBookById(fc *fiber.Ctx) error
}

//***********************************************************
func CreateBook(ctx *fiber.Ctx) error {
	db := database.GetDB()
	var book Book

	ctx.BodyParser(&book)
	db.Create(&book)
	return ctx.Status(fiber.StatusCreated).JSON(book)
}

//***********************************************************
func GetAllBooks(fc *fiber.Ctx) error {
	db := database.GetDB()
	var books Book
	//Todo implement filtering
	db.Find(&books, Book{Rating: 4})

	return fc.Status(fiber.StatusOK).JSON(&books)
}

//***********************************************************
func GetBookById(ctx *fiber.Ctx) error {
	db := database.GetDB()
	var book Book
	//Get the id from the url params
	id, err := strconv.ParseInt(ctx.Params("id"), 0, 0)
	if err != nil {
		//Supplied ID params is not an int
		return ctx.Status(fiber.StatusNotFound).
			SendString("No such book found")
	}
	db.Find(&book, id) //Get the book with the given ID
	if book.ID == 0 {  //If ID is zero no book was found with such ID
		return ctx.Status(fiber.StatusNotFound).
			SendString("No such book found")
	}
	return ctx.Status(fiber.StatusOK).JSON(&book)
}

//***********************************************************
func DeleteBookById(ctx *fiber.Ctx) error {
	db := database.GetDB()
	var book Book
	//Get the id from the url params
	id, err := strconv.ParseInt(ctx.Params("id"), 0, 0)
	if err != nil {
		//Supplied ID params is not an int
		return ctx.Status(fiber.StatusNotFound).
			SendString("No such book found")
	}
	db.Find(&book, id) //Get the book with the given ID
	if book.ID == 0 {  //If ID is zero no book was found with such ID
		return ctx.Status(fiber.StatusNotFound).
			SendString("No such book found")
	}
	db.Delete(&book, id)
	return ctx.Status(fiber.StatusOK).JSON(&book)
}

//***********************************************************
func UpdateBookById(ctx *fiber.Ctx) error {
	db := database.GetDB()
	var book Book
	var bookUpdate Book
	//Get the id from the url params
	id, err := strconv.ParseInt(ctx.Params("id"), 0, 0)
	if err != nil {
		//Supplied ID params is not an int
		return ctx.Status(fiber.StatusNotFound).
			SendString("No such book found")
	}
	db.Find(&book, id) //Get the book with the given ID
	if book.ID == 0 {  //If ID is zero no book was found with such ID
		return ctx.Status(fiber.StatusNotFound).
			SendString("No such book found")
	}
	ctx.BodyParser(&bookUpdate)
	if bookUpdate.Name != "" {
		book.Name = bookUpdate.Name
	}
	if bookUpdate.Author != "" {
		book.Author = bookUpdate.Author
	}
	if bookUpdate.Publication != "" {
		book.Publication = bookUpdate.Publication
	}
	if bookUpdate.Description != "" {
		book.Description = bookUpdate.Description
	}
	if bookUpdate.Rating > 1 || bookUpdate.Rating < 5 {
		book.Rating = bookUpdate.Rating
	}
	if bookUpdate.Year > 1700 {
		book.Year = bookUpdate.Year
	}

	db.Save(&book)
	return ctx.Status(fiber.StatusCreated).JSON(&book)
}
