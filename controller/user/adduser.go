package controller

import (
	"backend/dao/database"
	"backend/model"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddUser 用户添加处理函数
func AddUser(c *gin.Context) {
	// 创建结构体用于接收参数
	var user model.User

	// 使用 BindJSON 方法解析请求体中的 JSON 数据
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数解析失败",
		})
		return
	}

	// 将密码进行哈希加密
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "密码加密失败",
		})
		return
	}

	// 更新用户密码为哈希密码
	user.Password = hashedPassword

	// 创建用户记录
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(202, gin.H{
			"code":    202,
			"message": "用户添加失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "用户添加成功",
	})
}
