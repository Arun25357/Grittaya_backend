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

	router.GET("/GetProducts", rc.productController.GetProducts)
	router.POST("/CreateProduct", rc.productController.CreateProduct)
	router.GET("/GetAllProduct", rc.productController.GetAllProduct)
	router.POST("/UpdateProduct", rc.productController.UpdateProduct)
	router.DELETE("/DeleteProduct", rc.productController.DeleteProduct)
}
