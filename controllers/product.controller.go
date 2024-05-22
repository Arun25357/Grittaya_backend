package controllers

import (
	"errors"
	"net/http"

	"github.com/Pure227/Grittaya_backend/models"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(db *gorm.DB) ProductController {
	return ProductController{DB: db}
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	var products []models.Product
	if err := pc.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var payload models.CreateProduct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name:        payload.Name,
		Amount:      payload.Amount,
		UnitPrice:   payload.UnitPrice,
		Type:        payload.Type,
		Category:    payload.Category,
		Description: payload.Description,
	}

	if err := pc.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.FromString(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	var product models.Product
	if err := pc.DB.First(&product, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	// Bind JSON payload เพื่อให้รับข้อมูลจาก request body
	var payload models.UpdateProduct
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ทำการค้นหาสินค้าที่ต้องการอัปเดตในฐานข้อมูล
	var product models.Product
	if err := pc.DB.First(&product, "ID = ?", payload.ID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Product not found"})
		return
	}

	// อัปเดตข้อมูลของสินค้า
	product.Name = payload.Name
	product.Amount = payload.Amount
	product.UnitPrice = payload.UnitPrice
	product.Type = payload.Type
	product.Category = payload.Category
	product.Description = payload.Description

	// บันทึกการอัปเดตลงในฐานข้อมูล
	if err := pc.DB.Save(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to update product"})
		return
	}

	// ส่งคำตอบว่าอัปเดตสินค้าสำเร็จ
	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Product updated successfully"})
}


func (pc *ProductController) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.FromString(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := pc.DB.Delete(&models.Product{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
