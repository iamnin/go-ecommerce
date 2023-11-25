package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	Name        string       `gorm:"not null"`
	Description string       `gorm:"default:null"`
	Slug        string       `gorm:"uniqueIndex"`
	Products    []Product    `gorm:"many2many:products_categories;"`
	Images      []FileUpload `gorm:"foreignKey:CategoryId"`
	IsNewRecord bool         `gorm:"-;default:false"`
}
