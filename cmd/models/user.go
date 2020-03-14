package models

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Partner struct {
	gorm.Model
	Street    string `gorm:"not null"`
	City      string `gorm:"not null"`
	Country   string `gorm:"not null"`
	ZipCode   string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`

	User   User
	UserId uint    `gorm:"default:null"` // Guest Users may place an order, so they should be able to create an Partner with nullable UserId
	Orders []Order `gorm:"foreignKey:PartnerID"`
}

type User struct {
	gorm.Model
	Username    string  `gorm:"column:Username;unique_index"`
	Password    string  `gorm:"column:password"`
	Orders      []Order `gorm:"foreignkey:UserId"`
	LastOrderId uint    `gorm:"default:null"`
}

func registerNewUser(Username, password string) (*User, error) {
	return nil, errors.New("The password can't be empty")
}

func isUserValid(Username, password string) bool {
	return false
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

func (u *User) IsValidPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (user *User) SignClaims() string {
	jwt_token := jwt.New(jwt.SigningMethodHS512)
	jwt_token.Claims = jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 90).Unix(),
	}
	token, _ := jwt_token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token
}
