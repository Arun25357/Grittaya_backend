package controllers

import (
	"net/http"

	"github.com/Pure227/Grittaya_backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{DB: db}
}

func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var payload models.CreateOrder
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the customer exists
	var customer models.Customer
	if err := ctrl.DB.First(&customer, "id = ?", payload.CustomerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	//Search Productname
	var product models.Product
	if err := ctrl.DB.First(&product, "name = ?", payload.SetProductName).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	newOrder := models.CreateOrder{
		OrderDate:        payload.OrderDate,
		Status:           payload.Status,
		CustomerUsername: payload.CustomerUsername,
		Platform:         payload.Platform,
		DeliveryType:     payload.DeliveryType,
		TotalPrice:       payload.TotalPrice,
		Discount:         payload.Discount,
		SetProductID:     product.ID,
		CustomerID:       payload.CustomerID,
		UserID:           payload.UserID,
		Postcode:         payload.Postcode,
		SetProductName:   product.Name,
		Amount:           payload.Amount,
		Type:             product.Type,
		Price:            product.Price,
		PaymentType:      payload.PaymentType,
		LastPricePayment: payload.LastPricePayment,
	}

	if err := ctrl.DB.Create(&newOrder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newOrder)
}

func (ctrl *OrderController) GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := ctrl.DB.Where("id = ?", id).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (ctrl *OrderController) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.DB.Where("id = ?", id).Delete(&models.Order{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
