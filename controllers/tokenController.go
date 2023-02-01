package controllers

import (
	"auth-golang/auth"
	"auth-golang/database"
	"auth-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Where("email= ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credetialError := user.CheckPassword(request.Password)
	if credetialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.Header("Authorization", "bearer "+tokenString)
	context.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "statusCode": 200, "datas": user})
}
