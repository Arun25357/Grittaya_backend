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
	router.POST("/CreateSetProduct", rc.setProductController.CreateSetProduct)
	router.POST("/UpdateSetProduct", rc.setProductController.UpdateSetProduct)
	router.DELETE("/DeleteSetProduct", rc.setProductController.DeleteSetProduct)
	router.GET("/GetSetProduct", rc.setProductController.GetSetProduct)
	router.GET("/GetAllSetProducts", rc.setProductController.GetAllSetProducts)
}
