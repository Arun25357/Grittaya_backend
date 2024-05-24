package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
			Price:       payload.Price,
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
		Price:       product.Price,
		Type:        product.Type,
		Category:    product.Category,
		Description: product.Description,
		AttachFile:  product.AttachFile,
	}

	// Return the ticket
	ctx.JSON(http.StatusOK, gin.H{"status": "200", "data": getProduct})
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var payload models.CreateProduct
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	file, err := ctx.FormFile("attach_file")

	// Check if file was uploaded
	var AttachFile string
	if err != nil {
		if err != http.ErrMissingFile {
			log.Println("File upload error:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to upload file"})
			return
		}
		// No file uploaded
		AttachFile = ""
	} else {
		// Process file data
		ext := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		path := "public/product"

		// Check if the file extension is valid
		validExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
		validExtension := false
		for _, validExt := range validExtensions {
			if strings.EqualFold(ext, validExt) {
				validExtension = true
				break
			}
		}

		if !validExtension {
			log.Println("Invalid file extension:", ext)
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid file extension"})
			return
		}

		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				log.Println("Failed to create directory:", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to create directory"})
				return
			}
		}

		pathWithTime := filepath.Join(path, strconv.FormatInt(time.Now().Unix(), 10)+"-"+originalFileName+ext)
		AttachFile = pathWithTime

		// Save the uploaded file
		if err := ctx.SaveUploadedFile(file, pathWithTime); err != nil {
			log.Println("Failed to save file:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to save file"})
			return
		}
	}

	product := models.Product{
		Name:       payload.Name,
		Amount:     payload.Amount,
		Price:      payload.Price,
		Type:       payload.Type,
		Category:   payload.Category,
		AttachFile: AttachFile,
	}

	if err :=pc.DB.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "400", "message": "Can't create new ticket"})
		return
	}

	if err := pc.DB.Save(product).Error; err != nil {
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
