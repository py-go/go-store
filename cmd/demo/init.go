package demo

import (
	"go-store/cmd/models"
	"log"

	"go-store/cmd/database"

	"github.com/jinzhu/gorm"
)

func CreateDemoData(db *gorm.DB) {
	CreateProducts(db)
	CreateDiscounts(db)

}
func CreateProducts(db *gorm.DB) {

	products := []models.Product{
		{Name: "Apples1", Description: "Apples1", Price: 100.7},
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

func CreateDiscounts(db *gorm.DB) {
	discounts := []models.Discount{

		{
			Type:       0,
			Percerange: 30,
		},
		{
			Type:       1,
			Percerange: 10,
			Rules: []models.ProductCartRule{
				{Quantity: 7, ProductID: 5},
			},
		},
		{
			Type:       2,
			Percerange: 10,
			Rules: []models.ProductCartRule{
				{Quantity: 4, ProductID: 3},
				{Quantity: 2, ProductID: 4},
			},
		},
	}
	conn := database.Connection()
	for index, discount := range discounts {
		conn.Create(&discount)
		log.Println(index)
	}

}