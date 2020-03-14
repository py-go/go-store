package services

import (
	db "go-store/cmd/database"
	"go-store/cmd/models"
	m "go-store/cmd/models"
)

//RecordNotFound
func FechProductByID(productID uint) (product m.Product, err error) {
	database := db.Connection()
	err = database.First(&product, productID).Error
	return product, err
}

func FetchProducts() (products []models.Product, err error) {
	database := db.Connection()
	err = database.Find(&products).Error
	return products, err

}

func FechProductByIDs(productIds []uint, fields []string) (products []models.Product, err error) {
	database := db.Connection()
	err = database.Select(fields).Find(&products, productIds).Error
	return products, err
}
