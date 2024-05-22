package routes

import (
	"github.com/Pure227/Grittaya_backend/controllers"
	"github.com/gin-gonic/gin"
)

type ProductRouteController struct {
	productController controllers.ProductController
}

func NewProductRouteController(productController controllers.ProductController) ProductRouteController {
	return ProductRouteController{productController}
}

func (rc *ProductRouteController) ProductRoute(rg *gin.RouterGroup) {
	router := rg.Group("/products")

	router.GET("/", rc.productController.GetProducts)
	router.POST("/", rc.productController.CreateProduct)
	router.GET("/:id", rc.productController.GetProductByID)
	router.PUT("/:id", rc.productController.UpdateProduct)
	router.DELETE("/:id", rc.productController.DeleteProduct)
}
