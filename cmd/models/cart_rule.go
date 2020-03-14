package models

import (
	"github.com/jinzhu/gorm"
)

type Discount struct {
	gorm.Model
	Type       int               `gorm:"not null"`
	Percerange float64           `gorm:"default:0"`
	Rules      []ProductCartRule `gorm:"foreignKey:DiscountId"`
}

type ProductCartRule struct {
	gorm.Model
	Discount   Discount
	DiscountId uint `gorm:"not null"`
	// Partner   Partner `gorm:"association_foreignkey:PartnerId:"`
	Product   Product `gorm:"foreignkey:ProductId"`
	ProductId uint

	Quantity   int     `gorm:"not null"`
	Percerange float64 `gorm:"default:0"`
	// TrackingNumber string

}

// 	Partner   Partner `gorm:"association_foreignkey:PartnerId:"`
// 	PartnerId uint

// 	User   User `gorm:"foreignKey:UserId:"`
// 	d_type uint `gorm:"default:null"`

// 	AmountTotal float64
// 	Discount    float64
// 	Amount      float64
// // DiscountTable(D1):

// // d_type: direct
// // percerange: 30

// // ProductCartRule(R1):

// // Product: oranges
// // qty: null
// discount: D1
