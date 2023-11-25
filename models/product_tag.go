package models

type ProductTag struct {
	Tag       Tag `gorm:"association_foreignkey:TagId"`
	TagId     uint
	Product   Product `gorm:"association_foreignkey:ProductId"`
	ProductId uint
}

func (*ProductTag) TableName() string {
	return "products_tags"
}
