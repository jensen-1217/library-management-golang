package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ChangePassword 修改密码的处理函数
func ChangePassword(c *gin.Context) {
	// 解析请求参数
	username := c.Query("username")
	oldPassword := c.Query("oldPassword")
	newPassword := c.Query("newPassword")

	// 查询用户信息
	var user model.User
	if err := database.DB.Where("user_name = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// 验证旧密码是否匹配
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		c.JSON(202, gin.H{
			"code":    202,
			"message": "Incorrect old password"})
		return
	}

	// 哈希新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash new password"})
		return
	}

	// 更新用户密码为新密码
	database.DB.Model(&user).Update("password", string(hashedPassword))

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
