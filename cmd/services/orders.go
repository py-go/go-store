package services

import (
	db "go-store/cmd/database"
	m "go-store/cmd/models"
)

// TrackingNumber
func FetchOrdersByUser(user m.User) (orders []m.Order, err error) {
	database := db.Connection()
	err = database.Preload("OrderItems").Find(&orders, m.Order{UserID: user.ID}).Error
	return
}

func FetchOrderDetails(orderID uint) (order m.Order, err error) {
	database := db.Connection()
	err = database.Preload("OrderItems").First(&order, orderID).Error
	return order, err
}
