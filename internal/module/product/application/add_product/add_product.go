package addproduct

import (
	"github.com/georgejr3211/food-flow/internal/module/product/domain/entity"
	"github.com/georgejr3211/food-flow/internal/module/product/domain/repository"
	"github.com/georgejr3211/food-flow/internal/utils/logger"
)

type AddProductUseCase struct {
	productRepository repository.ProductRepository
}

func NewAddProductUseCase(productRepository repository.ProductRepository) *AddProductUseCase {
	return &AddProductUseCase{
		productRepository: productRepository,
	}
}

func (uc *AddProductUseCase) Execute(input *AddProductInput) (*AddProductOutput, *AddProductError) {
	product, err := entity.NewProduct(input.Name, input.Description, input.Price, input.Quantity, input.Status, "")
	if err != nil {
		return nil, &AddProductError{Message: err.Error()}
	}

	err = uc.productRepository.Save(product)
	if err != nil {
		return nil, &AddProductError{Message: err.Error()}
	}

	go func() {
		logger.LogToFile("Product added: " + product.Name)
	}()

	return &AddProductOutput{
		Product: ProductOutput{
			ID:          product.ID.Value(),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			Status:      product.Status,
		},
	}, nil
}
