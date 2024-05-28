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
	return UserController{DB: DB}
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	ac := NewAuthController(uc.DB)

	userData, err := ac.GetUserDataByToken(ctx)
	if err != nil {
		// Handle error appropriately (e.g., return unauthorized or bad request)
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid token"})
		return
	}

	// 2. Use a dedicated user model (improves data management and security)
	user := models.User{
		ID:       userData.ID,
		Username: userData.Username,
		Nickname: userData.Nickname,
		Position: userData.Position,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": user}})
}
func (uc *UserController) Test(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": nil})
}
