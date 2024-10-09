package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteBooks 批量删除图书信息
func DeleteBooks(c *gin.Context) {
	// 定义图书 ID 列表，用于接收前端传递的要删除的图书 ID
	var books []model.Book

	// 绑定 JSON 数据
	if err := c.BindJSON(&books); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid book data"})
		return
	}

	// 遍历图书 ID 列表，逐个执行删除图书的逻辑
	for _, book := range books {
		if err := database.DB.Delete(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "Failed to delete book"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Books deleted successfully"})
}
