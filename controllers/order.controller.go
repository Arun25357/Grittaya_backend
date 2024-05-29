package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/satori/go.uuid"
	"github.com/Pure227/Grittaya_backend/models"
)

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{DB: db}
}


// func (ctrl *OrderController) CreateOrder(c *gin.Context) {
// 	var input models.CreateOrder
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	order := models.Order{
// 		ID:               uuid.NewV4(),
// 		OrderDate:        input.OrderDate,
// 		Status:           input.Status,
// 		CustomerUsername: input.CustomerUsername,
// 		Platform:         input.Platform,
// 		DeliveryType:     input.DeliveryType,
// 		TotalPrice:       input.TotalPrice,
// 		Discount:         input.Discount,
// 		SetproductID:     input.SetproductID,
// 		CustomerID:       input.CustomerID,
// 		UserID:           input.UserID,
// 	}

// 	if err := ctrl.DB.Create(&order).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, order)
// }

func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var input models.CreateOrder
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    // Check if the customer exists
    var customer models.Customer
    if err := ctrl.DB.First(&customer, "id = ?", input.CustomerID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
        return
    }

	order := models.Order{
		ID:               uuid.NewV4(),
		OrderDate:        input.OrderDate,
		Status:           input.Status,
		CustomerUsername: input.CustomerUsername,
		Platform:         input.Platform,
		DeliveryType:     input.DeliveryType,
		TotalPrice:       input.TotalPrice,
		Discount:         input.Discount,
		SetproductID:     input.SetproductID,
		CustomerID:       input.CustomerID,
		UserID:           input.UserID,
        Postcode:         input.Postcode,
        SetproductName:   input.Name,
        Amount:           input.Amount,
        Type:             input.Type,
        Price:            input.Price,
        PaymentType:      input.PaymentType,
        TotalpricePayment: input.TotalpricePayment,
	}

	if err := ctrl.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
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
