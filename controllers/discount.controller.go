package controllers

import (
	"net/http"
	"time"

	"github.com/Pure227/Grittaya_backend/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type DiscountController struct {
	DB *gorm.DB
}

func NewDiscountController(DB *gorm.DB) DiscountController {
	return DiscountController{DB}
}

func (dc *DiscountController) CreateDiscount(ctx *gin.Context) {
	var payload struct {
		Percent int       `json:"percent" binding:"required"`
		Baht    int       `json:"bagt" binding:"required"`
		Expir   time.Time `json:"expir" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	discount := models.Discount{
		ID:      uuid.NewV4(),
		Baht:    payload.Baht,
		Percent: payload.Percent,
	}

	if err := dc.DB.Create(&discount).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to create discount"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Discount created successfully", "data": discount})
}

func (dc *DiscountController) GetDiscounts(ctx *gin.Context) {
	var discounts []models.Discount
	if err := dc.DB.Find(&discounts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve discounts"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "200", "data": discounts})
}

func (dc *DiscountController) UpdateDiscount(ctx *gin.Context) {
	var payload struct {
		ID      uuid.UUID `json:"id" binding:"required"`
		Baht    int       `json:"baht"`
		Percent int       `json:"percent"`
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	var discount models.Discount
	if err := dc.DB.First(&discount, "id = ?", payload.ID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Discount not found"})
		return
	}

	discount.Percent = payload.Percent
	discount.Baht = payload.Baht

	if err := dc.DB.Save(&discount).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to update discount"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Discount updated successfully", "data": discount})
}

func (dc *DiscountController) DeleteDiscount(ctx *gin.Context) {
	var payload struct {
		ID uuid.UUID `json:"id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	if err := dc.DB.Delete(&models.Discount{}, "id = ?", payload.ID).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to delete discount"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Discount deleted successfully"})
}
