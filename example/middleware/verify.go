package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// LoginVerify 中间件: 验证登录状态
func LoginVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("------verify")
		c.Next()
	}
}

// LoginNotVerify 中间件: 不验证登录状态
func LoginNotVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("------not verify")
		c.Next()
	}
}
