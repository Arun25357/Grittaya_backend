package routes

import (
	"github.com/Pure227/Grittaya_backend/controllers"
	"github.com/gin-gonic/gin"
)

type SetProductRouteController struct {
	setProductController *controllers.SetProductController
}

func NewSetProductRouteController(setProductController *controllers.SetProductController) *SetProductRouteController {
	return &SetProductRouteController{setProductController}
}

func (r *SetProductRouteController) SetProductRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/setproducts")
	router.POST("/CreateSetProduct", r.setProductController.CreateSetProduct)
	router.GET("/GetSetProduct:id", r.setProductController.GetSetProduct)
	router.POST("/UpdateSetProduct", r.setProductController.UpdateSetProduct)
	router.DELETE("/DeleteSetProduct:id", r.setProductController.DeleteSetProduct)
}
