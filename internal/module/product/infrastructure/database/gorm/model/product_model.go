package model

import (
	"github.com/georgejr3211/food-flow/internal/module/product/domain/entity"
	"github.com/georgejr3211/food-flow/internal/module/product/domain/entity/valueobject"
	"github.com/google/uuid"
)

type ProductModel struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name        string
	Description string
	Price       float64
	Quantity    int
	Status      bool
}

func (ProductModel) TableName() string {
	return "produtos"
}

func NewProductModel() *ProductModel {
	return &ProductModel{}
}

func (p *ProductModel) ToEntity() *entity.Product {
	return &entity.Product{
		ID:          *valueobject.NewProductId(p.ID.String()),
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Status:      p.Status,
	}
}

func (p *ProductModel) FromEntity(product *entity.Product) {
	p.ID = uuid.MustParse(product.ID.Value())
	p.Name = product.Name
	p.Description = product.Description
	p.Price = product.Price
	p.Quantity = product.Quantity
	p.Status = product.Status
}
