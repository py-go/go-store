package services

import (
	db "go-store/cmd/database"
	m "go-store/cmd/models"
)

// TrackingNumber
func FetchOrdersByUser(user m.User) (orders []m.Order, err error) {
	database := db.Connection()
	err = database.Model(m.Order{User: user}).Preload("OrderItems").Find(&orders).Error
	// var partner m.Partner
	// database.Model(&orders).Related(&partner)
	// order.Partner = partner
	return
}

func FetchOrderDetails(orderID uint) (order m.Order, err error) {
	database := db.Connection()
	err = database.Model(m.Order{}).Preload("OrderItems").First(&order, orderID).Error
	var partner m.Partner
	database.Model(&order).Related(&partner)
	order.Partner = partner
	return order, err
}
