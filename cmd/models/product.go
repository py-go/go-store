package models

import (
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"size:280;not null"`
	Description string `gorm:"not null"`
	Slug        string `gorm:"unique_index;not null"`
	Price       int    `gorm:"not null"`
}

func (product *Product) BeforeSave() (err error) {
	product.Slug = slug.Make(product.Name)
	return
}
