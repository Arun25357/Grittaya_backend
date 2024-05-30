package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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

	//  c.JSON(http.StatusOK, payload.ListProducts)

	// fmt.Println(payload.ListProducts);
	for i := 0; i < len(payload.ListProducts); i++ {
		var e = payload.ListProducts[i]
		var setproduct models.SetProduct
		if err := ctrl.DB.First(&setproduct, "name = ?", e.SetProductName).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
	}

	newOrder := models.Order{
		OrderDate:    payload.OrderDate,
		Status:       payload.Status,
		CustomerName: payload.CustomerName,
		Location:     payload.Location,
		Platform:     payload.Platform,
		DeliveryType: payload.DeliveryType,
		TotalPrice:   payload.TotalPrice,
		Discount:     payload.Discount,
		CustomerID:   customer.ID,
		Phone:        payload.Phone,
		UserID:       payload.UserID,
		Postcode:     payload.Postcode,
		Amount:       payload.Amount,
		// SetProductID:     setproduct.ID,
		// SetProductName:   setproduct.Name,
		// Type:             setproduct.Type,
		// Price:            setproduct.Price,
		PaymentType:      payload.PaymentType,
		LastPricePayment: payload.LastPricePayment,
	}

	// for i := 0; i < len(payload.ListProducts); i++ {
	// 	var e = payload.ListProducts[i]
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

	// Retrieve customer details
	var customer models.Customer
	if err := pc.DB.First(&customer, payload.CustomerID).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve customer"})
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
		Phone:            payload.Phone,
		Location:         payload.Location,
		Platform:         payload.Platform,
		DeliveryType:     payload.DeliveryType,
		TotalPrice:       payload.TotalPrice,
		CustomerID:       customer.ID,
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

func (pc *OrderController) GetOrder(ctx *gin.Context) {
	// Parse query parameters
	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid page number"})
		return
	}
	perPageStr := ctx.DefaultQuery("perPage", "500")
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid items per page number"})
		return
	}

	// Get total number of tickets
	var totalCount int64
	if err := pc.DB.Model(&models.Order{}).Count(&totalCount).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to get total number of products"})
		return
	}

	// Calculate offset and limit for pagination
	offset := (page - 1) * perPage
	limit := perPage

	// Retrieve orders based on pagination
	var orders []models.Order
	if err := pc.DB.Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve product"})
		return
	}

	// Convert retrieved orders to response format
	var getOrders []*models.GetOrder
	for _, payload := range orders {
		// Retrieve set product details
		var setproduct models.SetProduct
		if err := pc.DB.First(&setproduct, payload.SetProductID).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve set product"})
			return
		}

		// Retrieve customer details
		var customer models.Customer
		if err := pc.DB.First(&customer, payload.CustomerID).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve customer"})
			return
		}

		getOrder := &models.GetOrder{
			ID:               payload.ID,
			OrderDate:        payload.OrderDate,
			Status:           payload.Status,
			CustomerName:     payload.CustomerName,
			Location:         payload.Location,
			CustomerID:       customer.ID,
			Phone:            payload.Phone,
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
		getOrders = append(getOrders, getOrder)
	}

	// Return paginated orders
	ctx.JSON(http.StatusOK, gin.H{"status": "200", "data": gin.H{
		"orders":      getOrders,
		"totalCount":  totalCount,
		"currentPage": page,
		"perPage":     perPage,
		"nextPage":    fmt.Sprintf("?page=%d&perPage=%d", page+1, perPage),
		"prevPage":    fmt.Sprintf("?page=%d&perPage=%d", page-1, perPage),
	}})
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
