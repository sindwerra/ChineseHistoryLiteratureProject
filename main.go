package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()
	router.Use(Counter())
	index := router.Group("/")
	user := router.Group("/user")
	{
		user.GET("", func(context *gin.Context) {
			//context.String(http.StatusOK, "user page")
			context.JSON(200, gin.H{
				"key": "sindwerra",
				"value": "admin",
			})
		})
		user.GET("/register", func(context *gin.Context) {
			context.String(http.StatusOK, "registration page")
		})
		login := user.Group("/login")
		login.GET("/", func(context *gin.Context) {
			context.JSON(200, gin.H{"message": "login-page"})
		})
		login.GET("/v1", func(context *gin.Context) {
			context.JSON(200, gin.H{"message": "v1-login-page"})
		})
	}
	{
		index.GET("", func(context *gin.Context) {
			context.String(http.StatusOK, "hello gin")
		})
	}
	router.Run()
}