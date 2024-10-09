package controller

import (
	"backend/dao/database"
	"backend/model"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserInfoHandler 用于获取用户信息的处理函数
func UserInfo(c *gin.Context) {
	// 从请求头中获取 Token
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "未提供身份验证令牌"})
		return
	}

	// 解析 Token
	claims, err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "无效的身份验证令牌"})
		return
	}

	// 查询用户
	var user model.User
	if err := database.DB.Where("user_name = ?", claims.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  404,
			"error": "用户不存在"})
		return
	}

	// 构建用户信息响应
	userInfo := map[string]interface{}{
		"username":    user.UserName,
		"name":        user.Name,
		"phonenumber": user.PhoneNumber,
		"isadmin":     user.IsAdmin,
		// 其他用户信息
	}

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": userInfo})
}
