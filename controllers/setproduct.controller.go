package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Pure227/Grittaya_backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/satori/go.uuid" // Ensuring the same UUID library is used as in the models
)

type SetProductController struct {
	DB *gorm.DB
}

func NewSetProductController(db *gorm.DB) *SetProductController {
	return &SetProductController{DB: db}
}

func (ctrl *SetProductController) CreateSetProduct(c *gin.Context) {
	var input models.CreateSetProduct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	setProduct := models.SetProduct{
		ID:        uuid.NewV4(), // Use NewV4() for satori/go.uuid
		Name:      input.Name,
		Amount:    input.Amount,
		Price:     input.Price,
		Status:    input.Status,
		Type:      input.Type,
	}

	if err := ctrl.DB.Create(&setProduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Can't create new set product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "201", "data": setProduct})
}

func (ctrl *SetProductController) GetSetProduct(c *gin.Context) {
	id := c.Param("id")

	var setProduct models.SetProduct
	if err := ctrl.DB.Where("id = ?", id).First(&setProduct).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Set product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve set product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "200", "data": setProduct})
}

func (ctrl *SetProductController) UpdateSetProduct(c *gin.Context) {
	id := c.Param("id")

	var setProduct models.SetProduct
	if err := ctrl.DB.Where("id = ?", id).First(&setProduct).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Set product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve set product"})
		return
	}

	var input models.UpdateSetProduct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DB.Model(&setProduct).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to update set product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "200", "data": setProduct})
}

func (ctrl *SetProductController) DeleteSetProduct(c *gin.Context) {
	var input models.DeleteSetProduct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	if err := ctrl.DB.Where("id = ?", input.ID).Delete(&models.SetProduct{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Set product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to delete set product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Set product deleted successfully"})
}

func (ctrl *SetProductController) GetAllSetProducts(c *gin.Context) {
	// Parse query parameters
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid page number"})
		return
	}
	perPageStr := c.DefaultQuery("perPage", "50")
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid items per page number"})
		return
	}

	// Get total number of set products
	var totalCount int64
	if err := ctrl.DB.Model(&models.SetProduct{}).Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to get total number of set products"})
		return
	}

	// Calculate offset and limit for pagination
	offset := (page - 1) * perPage
	limit := perPage

	// Retrieve set products based on pagination
	var setProducts []models.SetProduct
	if err := ctrl.DB.Offset(offset).Limit(limit).Find(&setProducts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve set products"})
		return
	}

	// Convert retrieved set products to response format
	var getSetProducts []*models.GetSetProduct
	for _, payload := range setProducts {
		getSetProduct := &models.GetSetProduct{
			ID:          payload.ID,
			Name:        payload.Name,
			Amount:      payload.Amount,
			Price:       payload.Price,
			Status:      payload.Status,
			Type:        payload.Type,
		}
		getSetProducts = append(getSetProducts, getSetProduct)
	}

	// Return paginated set products
	c.JSON(http.StatusOK, gin.H{"status": "200", "data": gin.H{
		"setProducts": getSetProducts,
		"totalCount":  totalCount,
		"currentPage": page,
		"perPage":     perPage,
		"nextPage":    fmt.Sprintf("?page=%d&perPage=%d", page+1, perPage),
		"prevPage":    fmt.Sprintf("?page=%d&perPage=%d", page-1, perPage),
	}})
}
