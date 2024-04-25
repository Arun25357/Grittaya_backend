package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Pure227/Grittaya_backend/constants"
	"github.com/Pure227/Grittaya_backend/initializers"
	"github.com/Pure227/Grittaya_backend/models"
	"github.com/Pure227/Grittaya_backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

// [...] SignUp User
func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var payload *models.AdminSignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if payload.Password != payload.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	newUser := models.Admin{
		Username: payload.Username,
		Password: hashedPassword,
		Position: int(constants.Admin),
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	ac.DB.Save(newUser)
	ctx.JSON(http.StatusCreated, gin.H{"status": "OK", "message": "success"})
}

// [...] SignIn User
func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var payload *models.AdminSignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var tokenData models.Token
	var adminData models.Admin
	var result *gorm.DB
	var user models.Admin
	result = ac.DB.Where("username = ?", payload.Username).First(&adminData)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid Username or Password"})
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid Username or Password"})
		return
	}

	config, _ := initializers.LoadConfig(".")

	// Generate token that expire in 24 hours
	token, err := utils.GenerateToken(config.TokenExpiresIn, adminData.ID, config.TokenSecret)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	tokenData = models.Token{
		User_ID:   adminData.ID.String(),
		Token:     token,
		CreatedAt: time.Now().Unix(),
	}

	if ac.DB.Where("user_id = ?", adminData.ID) != nil {
		ac.DB.Model(&tokenData).Where("user_id = ?", adminData.ID).Delete(&tokenData)
	}

	ac.DB.Save(&tokenData)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "token": token, "position": adminData.Position})
}

// [...] SignOut User
func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	userID := GetUserIDByToken(ctx)
	var tokenData *models.Token
	if ac.DB.First(&tokenData, "user_id", userID).Delete(&tokenData).Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "You're already logged out"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

func GetUserIDByToken(ctx *gin.Context) (response string) {
	var token string
	authorizationHeader := ctx.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if authorizationHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if len(fields) != 2 || fields[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	token = fields[1]

	config, _ := initializers.LoadConfig(".")
	sub, err := utils.ValidateToken(token, config.TokenSecret)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	useridstring := fmt.Sprint(sub)
	// fmt.Println(useridstring)
	return useridstring
}

func (ac *AuthController) GetUserDataByToken(ctx *gin.Context) (res models.Admin) {
    authorizationHeader := ctx.GetHeader("Authorization")
    if authorizationHeader == "" {
        ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    fields := strings.Fields(authorizationHeader)
    if len(fields) != 2 || fields[0] != "Bearer" {
        ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    config, err := initializers.LoadConfig(".")
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    sub, err := utils.ValidateToken(fields[1], config.TokenSecret)
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    userID := fmt.Sprint(sub)
    var user models.Admin
    if err := ac.DB.First(&user, "id = ?", userID).Error; err != nil {
        ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    return user
}