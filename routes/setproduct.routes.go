package routes

import (
	"github.com/Pure227/Grittaya_backend/controllers"
	"github.com/gin-gonic/gin"
)

type SetProductRouteController struct {
	setProductController controllers.SetProductController
}

func NewSetProductRouteController(setProductController controllers.SetProductController) SetProductRouteController {
	return SetProductRouteController{setProductController}
}

func (rc *SetProductRouteController) SetProductRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/setproducts")
	router.POST("/", rc.setProductController.CreateSetProduct)
	router.PUT("/:id", rc.setProductController.UpdateSetProduct)
	router.DELETE("/:id", rc.setProductController.DeleteSetProduct)
	router.GET("/:id", rc.setProductController.GetSetProduct)
	router.GET("/", rc.setProductController.GetAllSetProducts)
}
