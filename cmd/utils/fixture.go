package utils

import (
	"go-store/cmd/models"
	"log"

	"go-store/cmd/database"

	"github.com/jinzhu/gorm"
)

func CreateProducts(db *gorm.DB) {
	products := []models.Product{
		{Name: "Apples", Description: "Apples", Price: 100},
		{Name: "Bananas", Description: "Bananas", Price: 100},
		{Name: "Pears", Description: "Pears", Price: 100},
		{Name: "Oranges", Description: "Oranges", Price: 100},
	}
	conn := database.Connection()
	for index, product := range products {
		conn.Create(&product)
		log.Println(index)
	}

}
