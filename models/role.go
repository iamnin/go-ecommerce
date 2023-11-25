package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model

	Name        string
	Description string
	Users       []User     `gorm:"many2many:users_roles"`
	UserRoles   []UserRole `gorm:"foreignKey:RoleId"`
}

type UserRole struct {
	User   User `gorm:"association_foreignkey:UserId"`
	UserId uint
	Role   User `gorm:"association_foreignkey:RoleId"`
	RoleId uint
}
