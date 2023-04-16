package product

import (
	addproduct "github.com/georgejr3211/food-flow/internal/module/product/application/add_product"
	"github.com/georgejr3211/food-flow/internal/module/product/infrastructure/database/gorm/repository"
	"github.com/georgejr3211/food-flow/internal/module/product/infrastructure/http"
	"gorm.io/gorm"
)

func InitProductModule(db *gorm.DB) *http.ProductRouter {
	repo := repository.NewProductRepository(db)
	uc := addproduct.NewAddProductUseCase(repo)
	ctrl := http.NewProductController(uc)
	router := http.NewProductRouter(ctrl)

	return router
}
