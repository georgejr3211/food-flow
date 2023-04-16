package repository

import (
	"github.com/georgejr3211/food-flow/internal/module/product/domain/entity"
	"github.com/georgejr3211/food-flow/internal/module/product/infrastructure/database/gorm/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Save(product *entity.Product) error {
	productModel := model.NewProductModel()
	productModel.FromEntity(product)
	return r.db.Create(productModel).Error
}
