package controller

import (
	"backend/dao/database"
	"backend/model"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 更新用户信息
func UpdateUser(c *gin.Context) {
	// 创建结构体用于接收参数
	var userData struct {
		ID          uint   `json:"id"`
		UserName    string `json:"username"`
		Password    string `json:"password"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phonenumber"`
		IsAdmin     int    `json:"isadmin"`
	}

	// 使用 BindJSON 方法解析请求体中的 JSON 数据
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数解析失败",
		})
		return
	}

	// 查询用户
	var user model.User
	if err := database.DB.Where("id = ?", userData.ID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "用户不存在",
		})
		return
	}

	// 更新用户信息
	if userData.UserName != "" {
		user.UserName = userData.UserName
	}
	if userData.Name != "" {
		user.Name = userData.Name
	}
	if userData.PhoneNumber != "" {
		user.PhoneNumber = userData.PhoneNumber
	}
	user.IsAdmin = userData.IsAdmin

	// 判断密码是否为空，不为空则加密新密码
	if userData.Password != "" {
		hashedPassword, err := utils.HashPassword(userData.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "加密密码失败",
			})
			return
		}
		user.Password = hashedPassword
	}

	// 保存更新后的用户信息到数据库
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "更新用户信息失败",
		})
		return
	}

	// 返回更新后的用户信息
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "更新用户信息成功",
		"data":    user,
	})
}
