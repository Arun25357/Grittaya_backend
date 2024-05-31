package controllers

import (
	"net/http"

	"github.com/Pure227/Grittaya_backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerController struct {
	DB *gorm.DB
}

func NewCustomerController(DB *gorm.DB) CustomerController {
	return CustomerController{DB}
}
func (cc *CustomerController) CreateCustomer(ctx *gin.Context) {
	var payload *models.CustomerDetails
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// var existingCustomer models.Customer
	// if err := cc.DB.Where("phone = ?", payload.Phone).First(&existingCustomer).Error; err == nil {
	// 	// Phone number found in the database
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"name":     existingCustomer.Name,
	// 		"location": existingCustomer.Location,
	// 		"postcode": existingCustomer.Postcode,
	// 	})
	// 	return
	// } else if err != gorm.ErrRecordNotFound {
	// 	// An error occurred during the query
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
	// 	return
	// }

	// Phone number not found, create a new customer
	newCustomer := models.Customer{
		Name:     payload.Name,
		Phone:    payload.Phone,
		Location: payload.Location,
		Postcode: payload.Postcode,
		Platform: payload.Platform,
	}

	if err := cc.DB.Create(&newCustomer).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create customer"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"name":     newCustomer.Name,
		"phone":    newCustomer.Phone,
		"location": newCustomer.Location,
		"postcode": newCustomer.Postcode,
		"platform": newCustomer.Platform,
	})
}
