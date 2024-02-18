package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Category string `json:"category" gorm:"type:varchar(255);unique;not null"`
}

type Product struct {
	gorm.Model
	Name        string   `json:"name" gorm:"type:varchar(255);unique;not null"`
	Description string   `json:"description" gorm:"type:varchar(255)"`
	Price       float64  `json:"price" gorm:"type:decimal(22,2)"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID"`
}

func (Product) TableName() string {
	return "product"
}
func (Category) TableName() string {
	return "category"
}
