package http

import (
	"net/http"

	"github.com/georgejr3211/food-flow/internal/module/product/infrastructure/http/middleware/validator"
	"github.com/gorilla/mux"
)

type ProductRouter struct {
	controller *ProductController
}

func NewProductRouter(controller *ProductController) *ProductRouter {
	return &ProductRouter{controller}
}

func (pr *ProductRouter) RegisterRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/products").Subrouter()
	subrouter.HandleFunc("", validator.AddProductValidatorMiddleware(pr.controller.HandleAddProduct)).Methods(http.MethodPost)
}
