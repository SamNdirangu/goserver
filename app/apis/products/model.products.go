package products

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Featured    bool    `json:"featured"`
	Rating      int     `json:"rating"`
	Company     string  `json:"company" sql:"type:ENUM('Toyota', 'Hamfrey','lenovo')"` // MySQL
}
