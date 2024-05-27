package routes

import (
	"github.com/Pure227/Grittaya_backend/controllers"
	"github.com/Pure227/Grittaya_backend/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("/users")
	router.GET("/profile", middleware.MiddlewareUser(), uc.userController.GetUser)

}
