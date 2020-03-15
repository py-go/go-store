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

func FetchProductCartRule(user m.User) (discounts []m.Discount, err error) {
	database := db.Connection()
	// err = database.Preload("Rules.Product", "id  IN (?)", []int{999}, func(db *gorm.DB) *gorm.DB {
	// 	return db.Where("id  IN (?)", []int{991119})
	// }).Find(&discounts).Error

	err = database.Preload("Rules").Preload("Rules.Product", "id  IN (?)", []int{999}).Find(&discounts).Error

	return
}

func FetchProductCartRuleV1(productIds []uint) (discounts []m.Discount, err error) {
	database := db.Connection()
	// db.Table("users").Select("users.name, emails.email"
	// ).Joins("left join emails on emails.user_id = users.id").Scan(&results)
	// db.Table("rules").Select("rules.name, products.price"
	// ).Joins("left join products on products.user_id = rules.id").Scan(&rules)
	// err = database.Select("product_cart_rules.quantity , discounts.*").Joins(
	// 	"join products on products.id = product_cart_rules.product_id AND products.id IN (?)",
	// 	productIds).Joins(
	// 	"join product_cart_rules on product_cart_rules.discount_id = discounts.id").Find(
	// 	&discounts).Error

	err = database.Joins(
		"join product_cart_rules on product_cart_rules.discount_id = discounts.id").Find(
		&discounts).Error

	// 		Discount
	// DiscountID
	// Product
	// ProductID
	// Quantity
	//.Where(
	// "product_cart_rules.product_id  IN (?)", []int{1, 5})
	// .Select("products.name, emails.email").Error
	// DB.Joins("left join emails on emails.user_id = users.id").First(&user).Error != nil

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
