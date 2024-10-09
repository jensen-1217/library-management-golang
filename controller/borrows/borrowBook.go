package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// BorrowBook 处理借书请求的控制器
func BorrowBook(c *gin.Context) {
	// 从请求中获取图书条形码
	barcode := c.Query("barcode")
	username := c.Query("username")
	if barcode == "" && username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少用户名或图书条形码"})
		return
	}

	// 检查图书是否可借
	book := model.Book{}
	if err := database.DB.Where("barcode = ?", barcode).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到对应的图书"})
		return
	}

	// 检查图书是否已被借阅
	if book.Status == 0 {
		c.JSON(201, gin.H{
			"code":  201,
			"error": "图书已被借出"})
		return
	}

	// 更新对应图书的状态为已归还
	if err := database.DB.Model(&book).Where("barcode = ?", book.Barcode).Update("status", 0).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法更新图书状态", "details": err.Error()})
		return
	}

	// 创建借阅记录
	borrow := model.Borrow{
		Barcode:       barcode,
		Title:         book.Title,
		UserName:      username, // TODO: 替换为实际用户
		Borrowtimestr: time.Now().Format("2006-01-02"),
	}

	// 将借阅记录保存到数据库
	if err := database.DB.Create(&borrow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建借阅记录"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "借书成功"})
}
