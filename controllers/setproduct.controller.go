package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/satori/go.uuid"
	"github.com/Pure227/Grittaya_backend/models"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	setProduct := models.SetProduct{
		ID:        uuid.NewV4(),
		Name:      input.Name,
		Amount:    input.Amount,
		Price:     input.Price,
		Status:    input.Status,
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
