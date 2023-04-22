package http

import (
	"github.com/georgejr3211/food-flow/internal/module/product"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(router *gin.Engine) *Router {
	return &Router{
		router: router,
	}
}

func (r *Router) Init(db *gorm.DB) {
	// Modules
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.router.Group("/api/v1")

	// Products
	productRouter := v1.Group("/products")
	productModule := product.InitProductModule(db)
	productModule.RegisterRoutes(productRouter)
}
