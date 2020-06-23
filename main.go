package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "historyProject/docs"
	"historyProject/middleware"
	"historyProject/service"
	"historyProject/service/search"
	"log"
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
	search.InitElasticSearch()

	// 默认添加了Logger和Recovery中间件
	app := gin.Default()
	app.Use(middleware.Counter())

	index := app.Group("/")
	user := app.Group("/user")
	login := user.Group("/login")
	elasticService := app.Group("/elastic")
	swaggerService := app.Group("/swagger")
	{
		index.GET("", service.IndexHandler)
		user.GET("", service.UserIndexHandler)
		user.GET("/register", service.UserRegisterHandler)
		login.GET("/", service.UserLoginHandler)
		login.GET("/v1", service.UserLoginV1Handler)
		elasticService.GET("/search", search.Endpoint)
		elasticService.POST("/documents", search.PostDocument)
		swaggerService.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}