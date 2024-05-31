package controllers

import (
	// "errors"
	// "os"
	// "path/filepath"
	// "strings"
	// "time"
	"errors"
	"fmt"
	"log"
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
	perPageStr := ctx.DefaultQuery("perPage", "500")
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
			Price:       payload.Price,
			Type:        payload.Type,
			Category:    payload.Category,
			Description: payload.Description,
			// AttachFile:  payload.AttachFile,
		}
		getProducts = append(getProducts, getProduct)
	}

	// Return paginated tickets
	ctx.JSON(http.StatusOK, gin.H{"status": "200", "data": gin.H{
		"products":    getProducts,
		"totalCount":  totalCount,
		"currentPage": page,
		"perPage":     perPage,
		"nextPage":    fmt.Sprintf("?page=%d&perPage=%d", page+1, perPage),
		"prevPage":    fmt.Sprintf("?page=%d&perPage=%d", page-1, perPage),
	}})
}

type GetProductByID struct {
	ID string `uri:"id"`
}

func (pc *ProductController) GetProducts(ctx *gin.Context) {
	// Parse query parameter

	req := GetProductByID{}
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	fmt.Println(req.ID)
	
	// Retrieve the product
	var product models.Product

	if err := pc.DB.Where("id = ?", req.ID).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Product not found"})
			fmt.Println(err)
			return
		}

		fmt.Println(1)
		fmt.Println(err)
		return
	}

	// Convert the product to response format
	getProduct := &models.GetProduct{
		ID:          product.ID,
		Name:        product.Name,
		Amount:      product.Amount,
		Price:       product.Price,
		Type:        product.Type,
		Category:    product.Category,
		Description: product.Description,
	}

	// Return the product
	ctx.JSON(http.StatusOK, gin.H{"status": "200", "data": getProduct})
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var payload models.CreateProduct
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(payload)
	// return
	product := models.Product{
		Name:     payload.Name,
		Amount:   payload.Amount,
		Price:    payload.Price,
		Type:     payload.Type,
		Category: payload.Category,
		// AttachFile: attachFile,
	}

	if err := pc.DB.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "400", "message": "Can't create new ticket"})
		return
	}

	if err := pc.DB.Save(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to save ticket"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "201", "data": product})
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
		Price:       payload.Price,
		Type:        payload.Type,
		Category:    payload.Category,
		Description: payload.Description,
		// AttachFile:  payload.AttachFile,
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
