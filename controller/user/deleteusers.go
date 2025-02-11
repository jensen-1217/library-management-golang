package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 删除多个用户信息
func DeleteUsers(c *gin.Context) {
	// 定义结构体数组，用于接收前端传递的用户信息列表
	var users []model.User
	if err := c.BindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid user data"})
		return
	}

	// 遍历用户列表，逐个执行删除用户的逻辑
	for _, user := range users {
		// 查询未归还的借阅记录
		var borrowCount int64
		if err := database.DB.Model(&model.Borrow{}).Where("user_name = ? AND returntimestr=?", user.UserName, "").Count(&borrowCount).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "Failed to query borrow records"})
			return
		}

		// 如果存在未归还的借阅记录，则拒绝删除用户
		if borrowCount > 0 {
			c.JSON(201, gin.H{
				"code":    201,
				"message": "There are still borrow records not returned for user: " + user.UserName})
			return
		}

		// 执行删除用户的逻辑
		if err := database.DB.Delete(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "Failed to delete user: " + user.UserName})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Users deleted successfully"})
}
