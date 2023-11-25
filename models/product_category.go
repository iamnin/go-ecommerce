package models

type ProductCategory struct {
	Category   Category `gorm:"association_foreignkey:CategoryId"`
	CategoryId uint
	Product    Product `gorm:"association_foreignkey:ProductId"`
	ProductId  uint
}
