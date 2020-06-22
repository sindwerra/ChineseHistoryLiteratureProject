package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var totalCount int = 0
var sucCount int = 0

func Counter() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 请求前部分
		totalCount++
		fmt.Printf("%d times request\n", totalCount)
		context.Next()
		// 响应后部分
		if statusCode := context.Writer.Status(); statusCode == 200 {
			sucCount++
		}
		fmt.Printf("%d times succeed\n", sucCount)
	}
}
