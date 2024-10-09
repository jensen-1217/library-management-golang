package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ReturnBook 处理图书归还请求的控制器
func ReturnBook(c *gin.Context) {
	// 获取请求参数
	idStr := c.Query("ID")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的借阅记录 ID"})
		return
	}

	// 从数据库中查找对应的借阅记录
	var borrow model.Borrow
	if err := database.DB.First(&borrow, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到对应的借阅记录", "details": err.Error()})
		return
	}

	// 获取当前时间并格式化为字符串
	currentTime := time.Now().Format("2006-01-02")

	// 更新借阅记录中的归还时间字段
	borrow.Returntimestr = currentTime

	// 保存更新后的借阅记录到数据库中
	if err := database.DB.Save(&borrow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法更新借阅记录", "details": err.Error()})
		return
	}

	// 更新对应图书的状态为已归还
	var book model.Book
	if err := database.DB.Model(&book).Where("barcode = ?", borrow.Barcode).Update("status", 1).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法更新图书状态", "details": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "图书归还成功"})
}
