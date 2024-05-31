package routes

import (
	"github.com/Pure227/Grittaya_backend/controllers"
	"github.com/gin-gonic/gin"
)

type OrderRouteController struct {
	orderController *controllers.OrderController
}

func NewOrderRouteController(orderController *controllers.OrderController) *OrderRouteController {
	return &OrderRouteController{orderController}
}

func (r *OrderRouteController) OrderRoute(rg *gin.RouterGroup) {
	router := rg.Group("/orders")
	router.POST("/CreateOrder", r.orderController.CreateOrder)
	router.GET("/GetOrder", r.orderController.GetOrder)
	router.POST("/UpdateOrder/:id", r.orderController.UpdateOrder)
	router.DELETE("/DeleteOrder", r.orderController.DeleteOrder)
}
