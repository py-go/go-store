package services

import (
	db "go-store/cmd/database"
	m "go-store/cmd/models"
)

func FetchOrderDetails(orderID uint) (order m.Order, err error) {
	database := db.Connection()
	err = database.Preload("OrderItems").First(&order, orderID).Error
	return
}

// TrackingNumber
// func FetchOrdersByUser(user m.User) (orders []m.Order, err error) {
// 	database := db.Connection()
// 	err = database.Preload("OrderItems").Find(&orders, m.Order{UserID: user.ID}).Error
// 	return
// }

func FetchOrdersByUser(condition interface{}) (order m.Order, err error) {
	database := db.Connection()
	err = database.Where(condition).Preload("OrderItems").Find(&order).Error
	return
}
