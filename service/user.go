package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// User Index Page godoc
// @Summary 用户页面引导页
// @Description 用户页面引导页
// @ID 1
// @Tags User
// @Accept  json
// @Produce  json
// @Router /user [get]
// @Success 200 object model.Result
func UserIndexHandler (context *gin.Context) {
	context.JSON(200, gin.H{
		"key": "sindwerra",
		"value": "admin",
	})
}

// User Index Page godoc
// @Summary 用户注册页面
// @Description 用户注册页面
// @ID 2
// @Tags User
// @Accept  json
// @Produce  json
// @Router /user/register [get]
func UserRegisterHandler (context *gin.Context) {
	context.String(http.StatusOK, "registration page")
}

// User Index Page godoc
// @Summary 用户登录页面
// @Description 用户登录页面
// @ID 3
// @Tags User
// @Accept  json
// @Produce  json
// @Router /user/login [get]
func UserLoginHandler (context *gin.Context) {
	context.JSON(200, gin.H{"message": "login-page"})
}

// User Index Page godoc
// @Summary 用户登录v1页面
// @Description 用户登录v1页面
// @ID 4
// @Tags User
// @Accept  json
// @Produce  json
// @Router /user/login/v1 [get]
func UserLoginV1Handler (context *gin.Context) {
	context.JSON(200, gin.H{"message": "v1-login-page"})
}