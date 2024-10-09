package controller

import (
	"backend/dao/database"
	"backend/model"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateBookInfo 更新图书信息
func UpdateBookInfo(c *gin.Context) {
	var newBook model.Book

	// 绑定 JSON 数据
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 根据图书ID查找旧记录
	var oldBook model.Book
	if err := database.DB.First(&oldBook, newBook.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find old book record"})
		return
	}

	// 更新字段
	oldBook.Barcode = newBook.Barcode
	oldBook.Title = newBook.Title
	oldBook.Classification = newBook.Classification
	oldBook.ClassificationName = utils.GetClassificationName(newBook.Classification) // 获取分类名称
	oldBook.Author = newBook.Author
	oldBook.Publisher = newBook.Publisher
	oldBook.Price = newBook.Price
	oldBook.ISBN = newBook.ISBN
	oldBook.Info = newBook.Info
	oldBook.Status = newBook.Status
	oldBook.Bookimg = newBook.Bookimg

	// 执行更新操作
	if err := database.DB.Save(&oldBook).Error; err != nil {
		c.JSON(202, gin.H{
			"code":  202,
			"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "图书信息更新成功",
		"data":    oldBook,
	})
}
