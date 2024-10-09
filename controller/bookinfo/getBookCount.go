package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBookCount 获取图书总数
func GetBookCount(c *gin.Context) {
	var count int64

	// 查询数据库获取图书总数
	if err := database.DB.Model(&model.Book{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取图书总数"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "图书总数获取成功",
		"data":    count,
	})
}
