package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name         string       `gorm:"not null"`
	Description  string       `gorm:"default:null"`
	Slug         string       `gorm:"uniqueIndex"`
	Products     []Product    `gorm:"many2many:products_tags;"`
	Images       []FileUpload `gorm:"foreignKey:TagId"`
	IsNewsRecord bool         `gorm:"-;default:false"` // Virtual Field, so it is not persisted in the Db. This is used in FirstOrCreate()
}

//func (a *Tag) BeforeSave() (err error) {
//	a.Slug = slug.Make(a.Name)
//	return
//}
