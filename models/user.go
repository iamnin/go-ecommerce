package models

import "gorm.io/gorm"

type User struct {
	//Generate timestamp and Id
	gorm.Model
	FirstName string `gorm:"varchar(255);not null"`
	LastName  string `gorm:"varchar(255);not null"`
	UserName  string `gorm:"column:username"`
	Email     string `gorm:"column:email;uniqueIndex"`
	Password  string `gorm:"password;not null"`

	Comments []Comment `gorm:"foreignKey:UserId"`

	Roles     []Role     `gorm:"many2many:users_roles;"`
	UserRoles []UserRole `gorm:"foreignKey:UserId"`
}
