package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	addproduct "github.com/georgejr3211/food-flow/internal/module/product/application/add_product"
)

type ProductController struct {
	addProductUseCase *addproduct.AddProductUseCase
}

func NewProductController(addProductUseCase *addproduct.AddProductUseCase) *ProductController {
	return &ProductController{
		addProductUseCase: addProductUseCase,
	}
}

func (c *ProductController) HandleAddProduct(w http.ResponseWriter, r *http.Request) {
	var input addproduct.AddProductInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
		return
	}

	output, appErr := c.addProductUseCase.Execute(&input)

	if appErr != nil {
		http.Error(w, fmt.Sprintf("error: %v", appErr), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
