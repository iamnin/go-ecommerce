package models

import "gorm.io/gorm"

type FileUpload struct {
	gorm.Model

	Filename     string
	FilePath     string
	OriginalName string
	FileSize     string

	Tag   Tag  `gorm:"association_foreignkey:TagId"`
	TagId uint `gorm:"default:null"`

	Category   Category `gorm:"association_foreignkey:CategoryId"`
	CategoryId uint

	Product   Category `gorm:"association_foreignkey:ProductId"`
	ProductId uint     `gorm:"default:null"`
}
