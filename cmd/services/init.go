package services

import (
	db "go-store/cmd/database"
)

func CreateOne(data interface{}) error {
	database := db.Connection()
	err := database.Create(data).Error
	return err
}

func SaveOne(data interface{}) error {
	database := db.Connection()
	err := database.Save(data).Error
	return err
}
