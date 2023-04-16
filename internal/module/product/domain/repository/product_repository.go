package repository

import (
	"github.com/georgejr3211/food-flow/internal/module/product/domain/entity"
)

type ProductRepository interface {
	// FindAll() ([]*entity.Product, error)
	// FindByID(id valueobject.ProductID) (*entity.Product, error)
	// Delete(id valueobject.ProductID)
	Save(product *entity.Product) error
	// Update(product *entity.Product) (*entity.Product, error)
}
