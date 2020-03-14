package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderStatus int         `gorm:"default:0"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderId"`

	Partner   Partner `gorm:"association_foreignkey:PartnerId:"`
	PartnerId uint

	User   User `gorm:"foreignKey:UserId"`
	UserId uint `gorm:"default:null"`

	AmountTotal float64
	Discount    float64
	Amount      float64
}

func (order *Order) Status() string {
	switch order.OrderStatus {
	case 0:
		return "draft"
	case 1:
		return "confirmed"
	default:
		return "unknown"
	}
}

func (order *Order) TrackingNumber() string {
	return fmt.Sprintf("SOOO-%v", order.ID)
}

type OrderItem struct {
	gorm.Model
	Order   Order
	OrderId uint `gorm:"not null"`

	Product   Product
	ProductId uint `gorm:"not null"`

	Slug        string `gorm:"not null"`
	ProductName string `gorm:"not null"`
	Price       int    `gorm:"not null"`
	Quantity    int    `gorm:"not null"`

	User   User `gorm:"association_foreignkey:UserId:"`
	UserId uint `gorm:"default:null"`
}
