package rest

import (
	"github.com/gin-gonic/gin"
	swaggerfile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "example.com/m/internal/api-gateway/docs"
)

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.Default()
	productGroup := router.Group("/products")
	{
		productGroup.POST("", h.CreateProduct)
		// productGroup.GET("/:id", h.GetProductByID)
		// productGroup.GET("/page", h.GetPagesProduct)
		// productGroup.DELETE("/:id", h.DeleteProductByID)
		// productGroup.PUT("/:id", h.UpdateProductByID)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfile.Handler))
	return router
}
