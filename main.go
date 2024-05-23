package main

import (
	"log"
	"net/http"

	"github.com/Pure227/Grittaya_backend/controllers"
	"github.com/Pure227/Grittaya_backend/initializers"
	"github.com/Pure227/Grittaya_backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server                    *gin.Engine
	AuthController            controllers.AuthController
	AuthRouteController       routes.AuthRouteController
	UserController            controllers.UserController
	UserRouteController       routes.UserRouteController
	ProductController         controllers.ProductController
	ProductRouteController    routes.ProductRouteController
	SetProductController      controllers.SetProductController
	SetProductRouteController routes.SetProductRouteController
	DiscountController        controllers.DiscountController
	DiscountRouteController   routes.DiscountRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.UserRouteController(UserRouteController)

	ProductController = controllers.NewProductController(initializers.DB)
	ProductRouteController = routes.NewProductRouteController(ProductController)

	SetProductController = *controllers.NewSetProductController(initializers.DB)
	SetProductRouteController = routes.NewSetProductRouteController(SetProductController)

	DiscountController = controllers.NewDiscountController(initializers.DB)
	DiscountRouteController = routes.NewDiscountRouteController(DiscountController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "http://localhost", config.ClientOrigin}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Access-Control-Allow-Origin", "*"}

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	ProductRouteController.ProductRoute(router)
	SetProductRouteController.SetProductRoutes(router)
	DiscountRouteController.DiscountRoutes(router)

	log.Fatal(server.Run(":" + config.BackendPort))
}
