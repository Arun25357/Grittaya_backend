package routes

import (
	"github.com/Pure227/Grittaya_backend/controllers"
	"github.com/Pure227/Grittaya_backend/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.GET("/login", rc.authController.SignInUser)
	router.POST("/logout", middleware.MiddlewareUser(), rc.authController.LogoutUser)
	router.DELETE("/detele",rc.authController.DeleteUser)
}
