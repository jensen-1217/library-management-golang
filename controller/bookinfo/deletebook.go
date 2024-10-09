package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteBookInfo 删除图书信息
func DeleteBookInfo(c *gin.Context) {
	var book model.Book

	// 绑定 JSON 数据
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	// 在数据库中查找要删除的图书信息
	if err := database.DB.Where("id = ?", book.ID).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到要删除的图书信息"})
		return
	}

	// 判断图书状态是否为0（已借出）
	if book.Status == 0 {
		c.JSON(202, gin.H{
			"code":  202,
			"error": "图书已借出，无法删除"})
		return
	}

	// 删除图书信息
	if err := database.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除图书信息失败", "details": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "图书信息删除成功"})
}
