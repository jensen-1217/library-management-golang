package controller

import (
	"backend/dao/database"
	"backend/model"
	"backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 检查用户名是否已存在

	if err := database.DB.Where("user_name = ?", user.UserName).First(&user).Error; err == nil {
		c.JSON(202, gin.H{
			"code":  202,
			"error": "用户名或手机号已存在"})
		return
	}

	// 检查手机号是否已存在
	if err := database.DB.Where("phone_number = ?", user.PhoneNumber).First(&user).Error; err == nil {
		c.JSON(202, gin.H{
			"code":  202,
			"error": "用户名或手机号已存在"})
		return
	}
	// 验证密码长度
	if len(user.Password) < 6 {
		c.JSON(400, gin.H{"error": "密码长度至少为6个字符"})
		return
	}

	// 对密码进行哈希处理
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "密码哈希处理失败"})
		return
	}
	user.Password = hashedPassword

	// 设置默认角色
	user.IsAdmin = 0

	// 创建用户
	if err := database.DB.Create(&user).Error; err != nil {
		fmt.Println("创建用户时发生错误:", err)
		c.JSON(400, gin.H{
			"code":  400,
			"error": "创建用户失败"})
		return
	}

	c.JSON(200, user)
}
