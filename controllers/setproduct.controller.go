package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pure227/Grittaya_backend/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SetProductController struct {
	DB *gorm.DB
}

func NewSetProductController(db *gorm.DB) *SetProductController {
	return &SetProductController{DB: db}
}

func (pc *ProductController) GetAllSetProduct(ctx *gin.Context) {
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
	var setproduct []models.SetProduct
	if err := pc.DB.Offset(offset).Limit(limit).Find(&setproduct).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve Setproduct"})
		return
	}

	// Convert retrieved tickets to response format
	var getsetproducts []*models.GetSetProduct
	for _, payload := range getsetproducts {
		getsetproduct := &models.GetSetProduct{
			ID:        payload.ID,
			Name:      payload.Name,
			Amount:    payload.Amount,
			Price:     payload.Price,
			Type:      payload.Type,
			ProductID: payload.ProductID,
			Status:    payload.Status,
			// AttachFile:  payload.AttachFile,
		}
		getsetproducts = append(getsetproducts, getsetproduct)
	}

	// Return paginated tickets
	ctx.JSON(http.StatusOK, gin.H{"status": "200", "data": gin.H{
		"products":    getsetproducts,
		"totalCount":  totalCount,
		"currentPage": page,
		"perPage":     perPage,
		"nextPage":    fmt.Sprintf("?page=%d&perPage=%d", page+1, perPage),
		"prevPage":    fmt.Sprintf("?page=%d&perPage=%d", page-1, perPage),
	}})
}
type GetSetProductByID struct {
	ID string `uri:"id"`
}

func (ctrl *SetProductController) CreateSetProduct(c *gin.Context) {
	var input models.CreateSetProduct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	setProduct := models.SetProduct{
		ID:     uuid.NewV4(),
		Name:   input.Name,
		Amount: input.Amount,
		Price:  input.Price,
		Status: input.Status,
	}

	if err := ctrl.DB.Create(&setProduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, setProduct)
}

func (ctrl *SetProductController) GetSetProduct(c *gin.Context) {
	id := c.Param("id")
	var setProduct models.SetProduct

	if err := ctrl.DB.Where("id = ?", id).First(&setProduct).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "SetProduct not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, setProduct)
}

func (ctrl *SetProductController) UpdateSetProduct(c *gin.Context) {
	id := c.Param("id")
	var setProduct models.SetProduct
	if err := ctrl.DB.Where("id = ?", id).First(&setProduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SetProduct not found"})
		return
	}

	var input models.UpdateSetProduct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctrl.DB.Model(&setProduct).Updates(input)

	c.JSON(http.StatusOK, setProduct)
}

func (ctrl *SetProductController) DeleteSetProduct(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.DB.Where("id = ?", id).Delete(&models.SetProduct{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "SetProduct not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
