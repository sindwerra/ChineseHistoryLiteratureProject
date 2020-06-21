package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserIndexHandler (context *gin.Context) {
	context.JSON(200, gin.H{
		"key": "sindwerra",
		"value": "admin",
	})
}

func UserRegisterHandler (context *gin.Context) {
	context.String(http.StatusOK, "registration page")
}

func UserLoginHandler (context *gin.Context) {
	context.JSON(200, gin.H{"message": "login-page"})
}

func UserLoginV1Handler (context *gin.Context) {
	context.JSON(200, gin.H{"message": "v1-login-page"})
}