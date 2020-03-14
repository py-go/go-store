package services

import (
	db "go-store/cmd/database"
	"go-store/cmd/models"
	m "go-store/cmd/models"

	"github.com/jinzhu/gorm"
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

func FetchProductCartRule(user m.User) (discounts []m.Discount, err error) {
	database := db.Connection()
	err = database.Preload("Rules.Product", "id  IN (?)", []int{999}, func(db *gorm.DB) *gorm.DB {
		return db.Where("id  IN (?)", []int{991119})
	}).Find(&discounts).Error
	return
}

func FetchProductCartRuleV1(user m.User) (rules []m.ProductCartRule, err error) {
	// database := db.Connection()
	// db.Table("rules").Select("rules.name, products.price"
	// ).Joins("left join products on products.user_id = rules.id").Scan(&rules)

	// DB.Joins("join rules on rules.product_id = product.id"
	// ).Where("name = ?", "joins").Find(&rules)

	// db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org"
	// ).Joins("JOIN credit_cards ON credit_cards.user_id = users.id"
	// ).Where("credit_cards.number = ?", "411111111111").Find(&user)
	return
}

// var user m.User
// database.Model(&order).Related(&user)
// order.User = user
// err = database.Preload("Rules", "product_id  IN (?)", []uint{5}).Where("active = ?", true).Find(&rules).Error

func FetchProductCartRulebyProductIDs(productIds []uint) (rules []m.ProductCartRule, err error) {
	database := db.Connection()
	err = database.Preload("Discount").Where("product_id  IN (?)", productIds).Find(&rules).Error
	return
}
