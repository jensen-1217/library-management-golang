package controller

import (
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogoutController 处理用户退出登录请求
func Logout(c *gin.Context) {
	// 从请求中获取令牌（Token）
	token := c.Query("token")

	// 检查令牌是否存在并且有效
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is missing"})
		return
	}

	// 检查令牌是否在黑名单中
	if utils.IsTokenBlacklisted(token) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been blacklisted"})
		return
	}

	// 将令牌添加到黑名单中
	utils.AddTokenToBlacklist(token)

	// 返回响应，确认退出登录成功
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
