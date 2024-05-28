package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

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
	var payload *models.UserSignUpInput

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

	newUser := models.User{
		Username: payload.Username,
		Password: hashedPassword,
		Position: int(constants.Admin),
		Nickname: payload.Nickname,
		Phone:    payload.Phone,
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

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	// var payload *models.UserSignInInput

	// if err := ctx.ShouldBindJSON(&payload); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	// 	return
	// }
	// var tokenData models.Token
	// var adminData models.User
	// var result *gorm.DB
	// var user models.User
	// result = ac.DB.First(&user, "username = ?", strings.ToLower(payload.Username))
	// if result.Error != nil {
	// 	// Handle error
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Error retrieving user data"})
	// 	return
	// }

	// if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid Username or Password"})
	// 	return
	// }

	// config, _ := initializers.LoadConfig(".")

	// // Generate token that expire in 24 hours
	// token, err := utils.GenerateToken(config.TokenExpiresIn, adminData.ID, config.TokenSecret)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	// 	return
	// }

	// adminData = user
	// tokenData = models.Token{
	// 	User_ID:   adminData.ID.String(),
	// 	Token:     token,
	// 	CreatedAt: time.Now().Unix(),
	// }

	// if ac.DB.Where("user_id = ?", adminData.ID) != nil {
	// 	ac.DB.Model(&tokenData).Where("user_id = ?", adminData.ID).Delete(&tokenData)
	// }

	// ac.DB.Save(&tokenData)

	// ctx.JSON(http.StatusOK, gin.H{"status": "success", "token": token, "position": adminData.Position, "UserID": adminData.ID})
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

func (ac *AuthController) GetUserDataByToken(ctx *gin.Context) (models.User, error) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		return models.User{}, errors.New("unauthorized: missing Authorization header")
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) != 2 || fields[0] != "Bearer" {
		return models.User{}, errors.New("unauthorized: invalid Authorization header format")
	}

	config, err := initializers.LoadConfig(".")
	if err != nil {
		return models.User{}, errors.New("internal server error: failed to load config")
	}

	sub, err := utils.ValidateToken(fields[1], config.TokenSecret)
	if err != nil {
		return models.User{}, errors.New("unauthorized: invalid token")
	}

	userID := fmt.Sprint(sub)
	var user models.User
	if err := ac.DB.First(&user, "id = ?", userID).Error; err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (pc *AuthController) DeleteUser(ctx *gin.Context) {
	var user models.User
	var payload *models.UserDelete
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": err.Error()})
		return
	}

	deleteuser := models.User{
		ID: payload.ID,
	}

	if err := pc.DB.First(&user, "ID = ?", deleteuser.ID).Delete(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "Product deleted successfully"})
}
