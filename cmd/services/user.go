package services

import (
	db "go-store/cmd/database"
	m "go-store/cmd/models"
)

func FindOneUser(condition interface{}) (m.User, error) {
	database := db.Connection()
	var model m.User

	err := database.Where(condition).First(&model).Error
	return model, err
}

func UpdateUser(user m.User, data interface{}) error {
	database := db.Connection()
	err := database.Model(user).Update(data).Error
	return err
}

func GetLastOrder(user m.User) error {
	// database := db.Connection()
	// database.Model(user)
	// err := database.Model(user).Update(data).Error
	// return err

	// var model m.User
	// err := database.Last(&model).Error
	// err := database.Where(condition).Preload("Roles").First(&user).Error
	// return user, err
	// return model, err
	return nil
}
