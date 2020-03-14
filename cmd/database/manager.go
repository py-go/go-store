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

//db := database.GetDB().Create(&user)
func Init() *gorm.DB {

	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// host := os.Getenv("DB_HOST")
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
	db.AutoMigrate(&m.User{})
	db.AutoMigrate(&m.Partner{})
	db.AutoMigrate(&m.Product{})
	db.AutoMigrate(&m.Order{})
	db.AutoMigrate(&m.OrderItem{})

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
