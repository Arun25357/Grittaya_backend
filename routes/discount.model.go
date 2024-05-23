package routes

import (
	"github.com/Pure227/Grittaya_backend/controllers"
	"github.com/gin-gonic/gin"
)

type DiscountRouteController struct {
	discountController controllers.DiscountController
}

func NewDiscountRouteController(discountController controllers.DiscountController) DiscountRouteController {
	return DiscountRouteController{discountController}
}

func (rc *DiscountRouteController) DiscountRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/discounts")

	router.POST("/create", rc.discountController.CreateDiscount)
	router.GET("/get", rc.discountController.GetDiscounts)
	router.POST("/update", rc.discountController.UpdateDiscount)
	router.DELETE("/delete", rc.discountController.DeleteDiscount)
}