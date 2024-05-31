package main

import (
	"fmt"
	"log"

	"github.com/Pure227/Grittaya_backend/initializers"
	"github.com/Pure227/Grittaya_backend/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.Customer{})
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Order{})
	initializers.DB.AutoMigrate(&models.Product{})
	initializers.DB.AutoMigrate(&models.SetProduct{})
	initializers.DB.AutoMigrate(&models.Discount{})
	initializers.DB.AutoMigrate(&models.Token{})
	initializers.DB.AutoMigrate(&models.OrderDetail{})
	fmt.Println("? Migration complete")
}
