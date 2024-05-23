package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"github.com/Pure227/Grittaya_backend/models"
)

// Controller for SetProduct
type SetProductController struct {
	DB *gorm.DB
}

func NewSetProductController(db *gorm.DB) *SetProductController {
	return &SetProductController{DB: db}
}

func (spc *SetProductController) CreateSetProduct(ctx *gin.Context) {
	var payload struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
		Items       []struct {
			ProductID uuid.UUID `json:"product_id" binding:"required"`
			Quantity  int       `json:"quantity" binding:"required"`
		} `json:"items" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	setProduct := models.SetProduct{
		ID:          uuid.NewV4(),
		Name:        payload.Name,
		Description: payload.Description,
	}

	for _, item := range payload.Items {
		setProductItem := models.SetProductItem{
			ID:           uuid.NewV4(),
			SetProductID: setProduct.ID,
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
		}
		setProduct.Items = append(setProduct.Items, setProductItem)
	}

	if err := spc.DB.Create(&setProduct).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to create set product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Set product created successfully"})
}

func (spc *SetProductController) UpdateSetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	setProductID, err := uuid.FromString(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid set product ID"})
		return
	}

	var payload struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Items       []struct {
			ProductID uuid.UUID `json:"product_id"`
			Quantity  int       `json:"quantity"`
		} `json:"items"`
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	var setProduct models.SetProduct
	if err := spc.DB.First(&setProduct, "id = ?", setProductID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Set product not found"})
		return
	}

	if payload.Name != "" {
		setProduct.Name = payload.Name
	}
	if payload.Description != "" {
		setProduct.Description = payload.Description
	}

	if len(payload.Items) > 0 {
		if err := spc.DB.Where("set_product_id = ?", setProductID).Delete(&models.SetProductItem{}).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to update set product items"})
			return
		}
		for _, item := range payload.Items {
			setProductItem := models.SetProductItem{
				ID:           uuid.NewV4(),
				SetProductID: setProductID,
				ProductID:    item.ProductID,
				Quantity:     item.Quantity,
			}
			setProduct.Items = append(setProduct.Items, setProductItem)
		}
	}

	if err := spc.DB.Save(&setProduct).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to update set product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Set product updated successfully"})
}

func (spc *SetProductController) DeleteSetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	setProductID, err := uuid.FromString(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid set product ID"})
		return
	}

	if err := spc.DB.Where("id = ?", setProductID).Delete(&models.SetProduct{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to delete set product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Set product deleted successfully"})
}

func (spc *SetProductController) GetSetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	setProductID, err := uuid.FromString(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid set product ID"})
		return
	}

	var setProduct models.SetProduct
	if err := spc.DB.Preload("Items").First(&setProduct, "id = ?", setProductID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Set product not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "data": setProduct})
}

func (spc *SetProductController) GetAllSetProducts(ctx *gin.Context) {
	var setProducts []models.SetProduct
	if err := spc.DB.Preload("Items").Find(&setProducts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "Failed to retrieve set products"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "data": setProducts})
}
