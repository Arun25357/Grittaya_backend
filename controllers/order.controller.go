package controllers

import (
	"errors"
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

	amount := 0
	for _, product := range payload.ListProducts {
		amount += product.Amount
	}

	newOrder := models.Order{
		OrderDate:        payload.OrderDate,
		Status:           0,
		CustomerName:     payload.CustomerName,
		Location:         payload.Location,
		Platform:         payload.Platform,
		DeliveryType:     payload.DeliveryType,
		TotalPrice:       payload.TotalPrice,
		Discount:         payload.Discount,
		CustomerID:       customer.ID,
		Phone:            payload.Phone,
		UserID:           payload.UserID,
		Postcode:         payload.Postcode,
		Amount:           amount,
		PaymentType:      payload.PaymentType,
		LastPricePayment: payload.LastPricePayment,
	}

	// for i := 0; i < len(payload.ListProducts); i++ {
	// 	var e = payload.ListProducts[i]
	if err := ctrl.DB.Create(&newOrder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderDetail := []models.OrderDetail{}
	for _, product := range payload.ListProducts {

		temp := models.OrderDetail{
			OrderID:      int(newOrder.ID),
			SetProductID: product.SetProductID,
			Amount:       product.Amount,
		}

		orderDetail = append(orderDetail, temp)
	}

	if err := ctrl.DB.Create(&orderDetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newOrder)
}

type GetOrder struct {
	ID string `uri:"id"`
}

func (pc *OrderController) UpdateOrder(ctx *gin.Context) {
	// Bind URI payload to get the order ID or other URI parameters
	var id GetOrder
	if err := ctx.BindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	// Bind JSON payload to get the update data
	var payload models.UpdateOrder
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	var orderDetail []models.OrderDetail

	if err := pc.DB.Where("order_id = ?", id.ID).Find(&orderDetail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Product not found"})
			fmt.Println(err)
			return
		}

		fmt.Println(1)
		fmt.Println(err)
		return
	}

	if len(orderDetail) > 0 {
		if err := pc.DB.Delete(&orderDetail).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Product not found"})
				fmt.Println(err)
				return
			}

			fmt.Println(1)
			fmt.Println(err)
			return
		}
	}

	amount := 0
	for _, product := range payload.ListProducts {
		amount += product.Amount
	}

	u64, err := strconv.ParseUint(id.ID, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	wd := uint(u64)

	newOrder := models.Order{
		ID:               uint(wd),
		OrderDate:        payload.OrderDate,
		Status:           0,
		CustomerName:     payload.CustomerName,
		Location:         payload.Location,
		Platform:         payload.Platform,
		DeliveryType:     payload.DeliveryType,
		TotalPrice:       payload.TotalPrice,
		Discount:         payload.Discount,
		Phone:            payload.Phone,
		UserID:           payload.UserID,
		Postcode:         payload.Postcode,
		Amount:           amount,
		PaymentType:      payload.PaymentType,
		LastPricePayment: payload.LastPricePayment,
	}

	if err := pc.DB.Save(&newOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Product not found"})
			fmt.Println(err)
			return
		}

		fmt.Println(1)
		fmt.Println(err)
		return
	}

	orderDetail = []models.OrderDetail{}
	for _, product := range payload.ListProducts {

		temp := models.OrderDetail{
			OrderID:      int(newOrder.ID),
			SetProductID: product.SetProductID,
			Amount:       product.Amount,
		}

		orderDetail = append(orderDetail, temp)
	}

	if err := pc.DB.Create(&orderDetail).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Order updated successfully"})
}

// // Retrieve customer details
// var customer models.Customer
// if err := pc.DB.First(&customer, payload.CustomerID).Error; err != nil {
// 	ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve customer"})
// 	return
// }

// Search for the product by name
// var setproduct models.SetProduct
// if err := pc.DB.First(&setproduct, "name = ?", payload.SetProductName).Error; err != nil {
// 	ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
// 	return
// }

// Create the update order struct
// for i := 0; i < len(payload.ListProducts); i++ {
// 	var e = payload.ListProducts[i]
// 	var setproduct models.SetProduct
// 	if err := pc.DB.First(&setproduct, "name = ?", e.SetProductName).Error; err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
// 		return
// 	}
// }

// amount := 0
// for _, product := range payload.ListProducts {
// 	amount += product.Amount
// }

// newOrder := models.Order{
// 	ID:           payload.ID,
// 	OrderDate:    payload.OrderDate,
// 	Status:       0,
// 	CustomerName: payload.CustomerName,
// 	Location:     payload.Location,
// 	Platform:     payload.Platform,
// 	DeliveryType: payload.DeliveryType,
// 	TotalPrice:   payload.TotalPrice,
// 	Discount:     payload.Discount,
// 	// CustomerID:       customer.ID,
// 	Phone:            payload.Phone,
// 	UserID:           payload.UserID,
// 	Postcode:         payload.Postcode,
// 	Amount:           amount,
// 	PaymentType:      payload.PaymentType,
// 	LastPricePayment: payload.LastPricePayment,
// }

// if err := pc.DB.First(&order, "ID = ?", newOrder).Error; err != nil {
// 	ctx.JSON(http.StatusBadGateway, gin.H{"status": "400", "message": "Order not found"})
// 	return
// }

// if err := pc.DB.Model(&order).Where("ID = ?", newOrder.ID).Updates(newOrder).Error; err != nil {
// 	ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to update Order"})
// 	return
// }

// if err := pc.DB.Delete(&newOrder).Error; err != nil {
// 	ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "An error occurred while logging out"})
// 	return
// }

// ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Order updated successfully"})

func (oc *OrderController) GetOrder(ctx *gin.Context) {
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

	orders2 := []models.Order{}
	if err := oc.DB.Find(&orders2).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to get total number of orders"})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"status": "200", "data": orders2})
	return

	// Get total number of orders
	// var totalCount int64
	// if err := oc.DB.Model(&models.Order{}).Count(&totalCount).Error; err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to get total number of orders"})
	// 	return
	// }

	// ctx.JSON(http.StatusInternalServerError, gin.H{"status": "200", "data": totalCount})
	// return

	// Calculate offset and limit for pagination
	offset := (page - 1) * perPage
	limit := perPage

	// Retrieve orders based on pagination
	var orders []models.Order
	if err := oc.DB.Offset(offset).Limit(limit).Preload("ListProducts").Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve orders"})
		return
	}
	// Convert retrieved orders to response format
	// var getOrders []*models.GetOrder
	// for _, order := range orders {
	// 	// Retrieve set product details
	// 	var setproduct models.SetProduct
	// 	if err := oc.DB.First(&setproduct, order.SetProductID).Error; err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve set product"})
	// 		return
	// 	}

	// 	// Check if the customer exists
	// 	var customer models.Customer
	// 	if err := oc.DB.First(&customer, "name = ?", order.CustomerName).Error; err != nil {
	// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
	// 		return
	// 	}

	// 	// Create GetOrder instance
	// 	getOrder := &models.GetOrder{
	// 		OrderDate:        order.OrderDate,
	// 		Status:           order.Status,
	// 		CustomerName:     order.CustomerName,
	// 		Location:         order.Location,
	// 		Platform:         order.Platform,
	// 		DeliveryType:     order.DeliveryType,
	// 		TotalPrice:       order.TotalPrice,
	// 		Discount:         order.Discount,
	// 		Phone:            order.Phone,
	// 		UserID:           order.UserID,
	// 		Postcode:         order.Postcode,
	// 		Amount:           order.Amount,
	// 		PaymentType:      order.PaymentType,
	// 		LastPricePayment: order.LastPricePayment,
	// 		ListProducts:     order.ListProducts, // Include ListProducts
	// 	}
	// 	getOrders = append(getOrders, getOrder)
	// }

	// // Return paginated orders
	// ctx.JSON(http.StatusOK, gin.H{"status": "200", "data": gin.H{
	// 	"orders":      getOrders,
	// 	"totalCount":  totalCount,
	// 	"currentPage": page,
	// 	"perPage":     perPage,
	// 	"nextPage":    fmt.Sprintf("?page=%d&perPage=%d", page+1, perPage),
	// 	"prevPage":    fmt.Sprintf("?page=%d&perPage=%d", page-1, perPage),
	// }})
}

// func (ctrl *OrderController) DeleteOrder(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := ctrl.DB.Where("id = ?", id).Delete(&models.Order{}).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
// 			return
// 		}
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

//		c.Status(http.StatusNoContent)
//	}func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
func (oc *OrderController) DeleteOrder(ctx *gin.Context) {
	var id GetOrder
	
	if err := ctx.BindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}
	
	var order []models.OrderDetail
	var order2 []models.Order
	if err := oc.DB.Where("order_id = ?", id.ID).Find(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Product not found"})
			fmt.Println(err)
			return
		}

		fmt.Println(order)
		fmt.Println(err)
		return
	}

	if len(order) > 0 {
		if err := oc.DB.Delete(&order).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Product not found"})
				fmt.Println(err)
				return
			}

			fmt.Println(2)
			fmt.Println(err)
			return
		}
	}
	if err := oc.DB.Where("id = ?", id.ID).Find(&order2).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Product not found"})
			fmt.Println(err)
			return
		}

		fmt.Println(order)
		fmt.Println(err)
		return
	}
	
	if len(order2) > 0 {
		if err := oc.DB.Delete(&order2).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Product not found"})
				fmt.Println(err)
				return
			}

			fmt.Println(2)
			fmt.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Product deleted successfully"})
}
