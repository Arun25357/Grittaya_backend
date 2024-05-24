package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pure227/Grittaya_backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(db *gorm.DB) ProductController {
	return ProductController{DB: db}
}

func (pc *ProductController) GetAllProduct(ctx *gin.Context) {
	// Parse query parameters
	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid page number"})
		return
	}
	perPageStr := ctx.DefaultQuery("perPage", "10")
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid items per page number"})
		return
	}

	// Get total number of tickets
	var totalCount int64
	if err := pc.DB.Model(&models.Product{}).Count(&totalCount).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to get total number of products"})
		return
	}

	// Calculate offset and limit for pagination
	offset := (page - 1) * perPage
	limit := perPage

	// Retrieve tickets based on pagination
	var product []models.Product
	if err := pc.DB.Offset(offset).Limit(limit).Find(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve product"})
		return
	}

	// Convert retrieved tickets to response format
	var getProducts []*models.GetProduct
	for _, payload := range product {
		getProduct := &models.GetProduct{
			ID:          payload.ID,
			Name:        payload.Name,
			Amount:      payload.Amount,
			UnitPrice:   payload.UnitPrice,
			Type:        payload.Type,
			Category:    payload.Category,
			Description: payload.Description,
			AttachFile:  payload.AttachFile,
		}
		getProducts = append(getProducts, getProduct)
	}

	// Return paginated tickets
	ctx.JSON(http.StatusOK, gin.H{"status": "200", "data": gin.H{
		"projects":    getProducts,
		"totalCount":  totalCount,
		"currentPage": page,
		"perPage":     perPage,
		"nextPage":    fmt.Sprintf("?page=%d&perPage=%d", page+1, perPage),
		"prevPage":    fmt.Sprintf("?page=%d&perPage=%d", page-1, perPage),
	}})
}

func (pc *ProductController) GetProducts(ctx *gin.Context) {
	// Parse query parameter
	productIDStr := ctx.Params.ByName("productID")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productID < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid product ID"})
		return
	}

	// Retrieve the ticket
	var product models.Product
	if err := pc.DB.First(&product, productID).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve the product"})
		return
	}

	// Convert the ticket to response format
	getProduct := &models.GetProduct{
		ID:          product.ID,
		Name:        product.Name,
		Amount:      product.Amount,
		UnitPrice:   product.UnitPrice,
		Type:        product.Type,
		Category:    product.Category,
		Description: product.Description,
		AttachFile:  product.AttachFile,
	}

	// Return the ticket
	ctx.JSON(http.StatusOK, gin.H{"status": "200", "data": getProduct})
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var payload models.CreateProduct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name:       payload.Name,
		Amount:     payload.Amount,
		UnitPrice:  payload.UnitPrice,
		Type:       payload.Type,
		Category:   payload.Category,
		AttachFile: payload.AttachFile,
	}

	if err := pc.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	// Bind JSON payload เพื่อให้รับข้อมูลจาก request body
	var product models.Product
	if err := ctx.BindUri(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	// ทำการค้นหาสินค้าที่ต้องการอัปเดตในฐานข้อมูล
	var payload models.UpdateProduct
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	// อัปเดตข้อมูลของสินค้า
	updateproduct := models.UpdateProduct{
		ID:          payload.ID,
		Name:        payload.Name,
		Amount:      payload.Amount,
		UnitPrice:   payload.UnitPrice,
		Type:        payload.Type,
		Category:    payload.Category,
		Description: payload.Description,
		AttachFile:  payload.AttachFile,
	}

	if err := pc.DB.First(&product, "ID = ?", updateproduct.ID).Error; err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "400", "message": "Product not found"})
		return
	}

	if err := pc.DB.Model(&product).Where("ID = ?", updateproduct.ID).Updates(updateproduct).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to update product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Product updated successfully"})
}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	var product models.Product
	var payload *models.DeleteProduct
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	deleteProduct := models.Product{
		ID: payload.ID,
	}

	if err := pc.DB.First(&product, "ID = ?", deleteProduct.ID).Delete(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Product deleted successfully"})
}
