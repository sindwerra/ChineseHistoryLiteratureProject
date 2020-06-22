package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index Page godoc
// @Summary Index引导页
// @Description Index引导页
// @ID 1
// @Tags Index
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Router / [get]
func IndexHandler (context *gin.Context) {
	context.String(http.StatusOK, "hello gin")
}
