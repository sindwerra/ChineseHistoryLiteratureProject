package main

import (
	"ginProject/service"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	// 默认添加了Logger和Recovery中间件
	app := gin.Default()
	app.Use(Counter())

	index := app.Group("/")
	{
		index.GET("", service.IndexHandler)
	}

	user := app.Group("/user")
	login := user.Group("/login")
	{
		user.GET("", service.UserIndexHandler)
		user.GET("/register", service.UserRegisterHandler)
		login.GET("/", service.UserLoginHandler)
		login.GET("/v1", service.UserLoginV1Handler)
	}

	app.Run()
}