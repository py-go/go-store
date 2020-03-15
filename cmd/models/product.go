package models

import (
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"size:280;not null"`
	Description string  `gorm:"not null"`
	Slug        string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
}

func (product *Product) BeforeSave() (err error) {
	product.Slug = slug.Make(product.Name)
	return
}

type Discount struct {
	gorm.Model
	Type       int               `gorm:"not null"`
	Percerange float64           `gorm:"default:0"`
	Rules      []ProductCartRule `gorm:"foreignKey:DiscountID"`
}

type ProductCartRule struct {
	gorm.Model
	Discount   Discount
	DiscountID uint `gorm:"not null"`

	Product   Product
	ProductID uint `gorm:"default:null"`
	Quantity  int  `gorm:"default:0"`
}
