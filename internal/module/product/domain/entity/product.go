package entity

import (
	"errors"

	"github.com/georgejr3211/food-flow/internal/module/product/domain/entity/valueobject"
)

type Product struct {
	ID          valueobject.ProductID
	Name        string
	Description string
	Price       float64
	Quantity    int
	Status      bool
}

func NewProduct(name, description string, price float64, quantity int, status bool, id string) (*Product, error) {
	productID := valueobject.NewProductId(id)

	product := &Product{
		ID:          *productID,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		Status:      status,
	}

	if err := product.validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}

	if p.Description == "" {
		return errors.New("description is required")
	}

	if p.Price <= 0 {
		return errors.New("price should be greater than zero")
	}

	if p.Quantity <= 0 {
		return errors.New("quantity should be greater than zero")
	}

	return nil
}
