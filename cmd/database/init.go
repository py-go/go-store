package database

import (
	m "go-store/cmd/models"
	"log"
	"os"
	"path"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func Init() *gorm.DB {

	dbInstance, err := gorm.Open("sqlite3", path.Join(".", "app.db"))
	if err != nil {
		log.Fatalf("failed to open DB connection: %v", err)
	}

	dbInstance.DB().SetMaxIdleConns(10)
	dbInstance.LogMode(true)
	SetConnection(dbInstance)
	available := IsConnectionAvailable(5)
	if !available {
		log.Fatalf("Failed to Ping DB connection: %v", err)
	}

	log.Printf("Done.")
	AutoMigrate()
	return dbInstance
}

func IsConnectionAvailable(retries int) bool {
	err := Connection().DB().Ping()
	if err != nil {
		if retries == 0 {
			return false
		} else {
			time.Sleep(2 * time.Second)
			return IsConnectionAvailable(retries - 1)
		}
	}

	return true
}

func Connection() *gorm.DB {
	return db
}

func SetConnection(dbInstance *gorm.DB) {
	db = dbInstance
}

func AutoMigrate() {
	// db.DropTableIfExists(
	// 	&m.Order{},
	// 	&m.Partner{},
	// 	&m.OrderItem{})
	db.AutoMigrate(&m.User{})
	db.AutoMigrate(&m.Partner{})
	db.AutoMigrate(&m.Product{})
	db.AutoMigrate(&m.Order{}, &m.OrderItem{})
	db.AutoMigrate(&m.Discount{}, &m.ProductCartRule{})

}

func DropTableIfExists() {
	db.DropTableIfExists(
		&m.User{},
		&m.Partner{},
		&m.Product{},
		&m.Order{},
		&m.OrderItem{})
}
func RemoveDb(db *gorm.DB) error {
	db.Close()
	err := os.Remove(path.Join(".", "app.db"))
	return err
}
