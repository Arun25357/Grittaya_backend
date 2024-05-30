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
	if err := ctrl.DB.First(&customer, "name = ?", payload.CustomerName).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	//Search Productname
	var setproduct models.SetProduct
	if err := ctrl.DB.First(&setproduct, "name = ?", payload.SetProductName).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	newOrder := models.Order{
		OrderDate:        payload.OrderDate,
		Status:           payload.Status,
		CustomerName:     payload.CustomerName,
		Platform:         payload.Platform,
		DeliveryType:     payload.DeliveryType,
		TotalPrice:       payload.TotalPrice,
		Discount:         payload.Discount,
		CustomerID:       customer.ID,
		SetProductID:     setproduct.ID,
		UserID:           payload.UserID,
		Postcode:         payload.Postcode,
		SetProductName:   setproduct.Name,
		Amount:           payload.Amount,
		Type:             setproduct.Type,
		Price:            setproduct.Price,
		PaymentType:      payload.PaymentType,
		LastPricePayment: payload.LastPricePayment,
	}

	if err := ctrl.DB.Create(&newOrder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newOrder)
}

func (pc *OrderController) UpdateOrder(ctx *gin.Context) {
	// Bind URI payload to get the order ID or other URI parameters
	var order models.Order
	if err := ctx.BindUri(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	// Bind JSON payload to get the update data
	var payload models.UpdateOrder
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	// Search for the product by name
	var setproduct models.SetProduct
	if err := pc.DB.First(&setproduct, "name = ?", payload.SetProductName).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Create the update order struct
	updateorder := models.UpdateOrder{
		ID:               payload.ID,
		OrderDate:        payload.OrderDate,
		Status:           payload.Status,
		CustomerName:     payload.CustomerName,
		Platform:         payload.Platform,
		DeliveryType:     payload.DeliveryType,
		TotalPrice:       payload.TotalPrice,
		Discount:         payload.Discount,
		SetProductID:     setproduct.ID,
		UserID:           payload.UserID,
		Postcode:         payload.Postcode,
		SetProductName:   setproduct.Name,
		Amount:           payload.Amount,
		Type:             setproduct.Type,
		Price:            setproduct.Price,
		PaymentType:      payload.PaymentType,
		LastPricePayment: payload.LastPricePayment,
		// AttachFile:  payload.AttachFile,
	}

	if err := pc.DB.First(&order, "ID = ?", updateorder.ID).Error; err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "400", "message": "Order not found"})
		return
	}

	if err := pc.DB.Model(&order).Where("ID = ?", updateorder.ID).Updates(updateorder).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to update Order"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Order updated successfully"})
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
