package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	OrderStatus    int `gorm:"default:0"`
	TrackingNumber string

	OrderItems []OrderItem `gorm:"foreignKey:OrderId"`

	Address   Address `gorm:"association_foreignkey:AddressId;"`
	AddressId uint

	User            User `gorm:"foreignKey:UserId"`
	UserId          uint `gorm:"default:null"`
	OrderItemsCount int  `gorm:"-"`
}
