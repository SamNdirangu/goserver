package products

import (
	"goserver/app/database"
	"goserver/app/functions"

	"github.com/gofiber/fiber/v2"
)

type ProductControllers struct{}
type IProductControllers interface {
	CreateProduct(ctx *fiber.Ctx) error
	GetAllProducts(ctx *fiber.Ctx) error
	GetProductById(ctx *fiber.Ctx) error
	UpdateProductById(fc *fiber.Ctx) error
	DeleteProductById(fc *fiber.Ctx) error
}

//
//***********************************************************
type queryModel struct {
	Name        string
	Description string
	Price       string
	Featured    string
	Rating      string
	Company     string
	Sorting     string //Direction of sorting
	Page        int    // Page to display
	Limit       int    // Number of items to display per page
}

//***********************************************************
func CreateProduct(ctx *fiber.Ctx) error {
	db := database.GetDB()
	var product Product

	ctx.BodyParser(&product)
	db.Create(&product)
	return ctx.Status(fiber.StatusCreated).JSON(product)
}

func GetAllProducts(ctx *fiber.Ctx) error {
	//Get all the passed queries
	var query queryModel
	ctx.QueryParser(&query)

	db := database.GetDB()
	//Filters passed via query========================================
	//Schema based filters -------------------------------------------
	if query.Name != "" { //Name Filter
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Featured != "" {
		db = db.Where("featured=?", query.Featured)
	}
	if query.Company != "" {
		db = db.Where("company=?", query.Company)
	}
	//Numerical based filters -----------------------------------------
	if query.Price != "" {
		db, _ = functions.NumericalQueryBuilder(db, query.Price, "price")
	}
	if query.Rating != "" {
		db, _ = functions.NumericalQueryBuilder(db, query.Rating, "rating")
	}
	//Paging and Sorting=================================================
	//Sorting -----------------------------------------------------------
	if query.Sorting != "" {
		sortingQuery := functions.SortQueryBuilder(query.Sorting)
		db = db.Order(sortingQuery)
	}
	// Paging -----------------------------------------------------------
	limit, itemOffset := functions.PagingQueryBuilder(query.Page, query.Limit)
	db = db.Limit(limit).Offset(itemOffset)

	//Now get them products ================================================
	var products []Product
	err := db.Find(&products).Error
	if err != nil {
		//TODO Implement better errorhandler
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	//TODO implement better json output
	return ctx.Status(fiber.StatusOK).JSON(&products)
}
