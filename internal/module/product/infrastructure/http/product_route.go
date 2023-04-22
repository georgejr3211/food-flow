package http

import (
	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
	controller *ProductController
}

func NewProductRouter(controller *ProductController) *ProductRouter {
	return &ProductRouter{controller}
}

func (pr *ProductRouter) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/", pr.controller.HandleAddProduct)
	// router.HandleFunc("", validator.AddProductValidatorMiddleware(pr.controller.HandleAddProduct)).Methods(http.MethodPost)
}
