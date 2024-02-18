package repository

import (
	"log"

	"product-service/model"

	"gorm.io/gorm"
)

type ProductRepo interface {
	Migrate() error
}

type productRepo struct {
	DB *gorm.DB
}

func (c productRepo) Migrate() error {
	log.Print("[ProductRepository]...Migrate")
	return c.DB.AutoMigrate(&model.Product{})
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return productRepo{
		DB: db,
	}
}
