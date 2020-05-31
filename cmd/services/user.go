package services

import (
	db "go-store/cmd/database"
	m "go-store/cmd/models"
)

func FindOneUser(condition interface{}) (model m.User, err error) {
	database := db.Connection()
	err = database.Where(condition).First(&model).Error
	return
}

func UpdateUser(user m.User, data interface{}) (err error) {
	database := db.Connection()
	err = database.Model(user).Update(data).Error
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
