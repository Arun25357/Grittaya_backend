package routes

import (
	"github.com/Pure227/Grittaya_backend/controllers"
	"github.com/gin-gonic/gin"
)

type CustomerRouteController struct {
	customerController controllers.CustomerController
}

func NewCustomerRouteController(customerController controllers.CustomerController) CustomerRouteController {
	return CustomerRouteController{customerController}
}
func (rc *CustomerRouteController) CustomerRoute(rg *gin.RouterGroup) {
	router := rg.Group("/customer")

	router.POST("/create", rc.customerController.CreateCustomer)
}
