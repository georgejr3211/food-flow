package addproduct_test

import (
	"fmt"
	"testing"

	addproduct "github.com/georgejr3211/food-flow/internal/module/product/application/add_product"
	"github.com/georgejr3211/food-flow/internal/module/product/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Save(product *entity.Product) error {
	args := m.Called(product)
	fmt.Println(product)

	return args.Error(0)
}

func TestAddProductUseCase(t *testing.T) {
	// Create a mock product repository
	productRepository := new(MockProductRepository)

	// Set up input data and return value
	productRepository.On("Save", mock.AnythingOfType("*entity.Product")).Return(nil)

	addProductUseCase := addproduct.NewAddProductUseCase(productRepository)

	testInput := &addproduct.AddProductInput{
		Name:        "Test Product",
		Description: "A test product",
		Price:       9.99,
		Quantity:    100,
		Status:      true,
	}

	output, err := addProductUseCase.Execute(testInput)

	assert.IsType(t, "", output.Product.ID)

	expectedOutput := &addproduct.AddProductOutput{
		Product: addproduct.ProductOutput{
			ID:          output.Product.ID,
			Name:        "Test Product",
			Description: "A test product",
			Price:       9.99,
			Quantity:    100,
			Status:      true,
		},
	}

	assert.Equal(t, expectedOutput, output)
	assert.Nil(t, err)

	productRepository.AssertCalled(t, "Save", mock.AnythingOfType("*entity.Product"))
}
