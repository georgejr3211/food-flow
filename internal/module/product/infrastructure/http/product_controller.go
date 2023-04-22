package http

import (
	"net/http"

	addproduct "github.com/georgejr3211/food-flow/internal/module/product/application/add_product"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	addProductUseCase *addproduct.AddProductUseCase
}

func NewProductController(addProductUseCase *addproduct.AddProductUseCase) *ProductController {
	return &ProductController{
		addProductUseCase: addProductUseCase,
	}
}

// @Summary Add a new product
// @Description Add a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body addproduct.AddProductInput true "Product"
// @Success 201 {object} addproduct.AddProductOutput
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
func (ctrl *ProductController) HandleAddProduct(c *gin.Context) {

	var input addproduct.AddProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, appErr := ctrl.addProductUseCase.Execute(&input)

	if appErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": appErr.Message})
		return
	}

	c.JSON(http.StatusCreated, output)
}
