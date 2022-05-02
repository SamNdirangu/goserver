package books

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Rating      int    `json:"rating"`
	ISBN        int    `json:"isbn"`
}
