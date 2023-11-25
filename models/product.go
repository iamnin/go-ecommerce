package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name        string `gorm:"size:280;not null"`
	Description string `gorm:"not null"`
	Slug        string `gorm:"uniqueIndex;not null"`
	Price       int    `gorm:"not null"`
	Stock       int    `gorm:"not null"`

	Tags        []Tag        `gorm:"many2many:products_tags;"`
	ProductTags []ProductTag `gorm:"foreign:ProductId"`

	Categories        []Category        `gorm:"many2many:products_categories;"`
	ProductCategories []ProductCategory `gorm:"foreignKey:ProductId"`

	Comments      []Comment    `gorm:"foreignKey:ProductId"`
	Images        []FileUpload `gorm:"foreignKey:ProductId"`
	CommentsCount int          `gorm:"-"`
}
