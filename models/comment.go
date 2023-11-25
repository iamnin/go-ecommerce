package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	Content   string  `gorm:"size:2028"`
	Rating    int     `gorm:"default:null"`
	Product   Product `gorm:"foreignKey:ProductId"`
	ProductId uint
	User      User `gorm:"foreignKey:UserId"`
	UserId    uint `gorm:"not null"`
}
