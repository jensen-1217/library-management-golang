package utils

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UseCorsMiddleware(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:30003"},                 // 允许的源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},           // 允许的HTTP方法
		AllowHeaders:     []string{"Content-Type", "Authorization", "token"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                         // 暴露的响应头
		AllowCredentials: true,                                               // 允许携带认证信息（如Cookie）
		MaxAge:           12 * time.Hour,                                     // 预检请求（OPTIONS）的有效期
	}))
}
