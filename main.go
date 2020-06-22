package main

import (
	_ "ginProject/docs"
	"ginProject/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin swagger
// @version 1.0
// @description swagger示例

// @contact.name sindwerra
// @contact.email sindwerra@hotmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

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

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run()
}