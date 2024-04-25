package controllers

import (
	"net/http"

	"github.com/Pure227/Grittaya_backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	ac := NewAuthController(uc.DB)
	getUserDataByToken := ac.GetUserDataByToken(ctx)

	userData := models.Admin{
		ID:       getUserDataByToken.ID,
		Username: getUserDataByToken.Username,
		Nickname: getUserDataByToken.Nickname,
		Position: getUserDataByToken.Position,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userData}})
}
