package controller

import (
	"backend/dao/database"
	"backend/model"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 添加图书信息的控制器
// AddBookInfo 添加图书信息
func AddBookInfo(c *gin.Context) {
	var book model.Book

	// 绑定 JSON 数据
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查图书的 Barcode 是否已存在

	if err := database.DB.Where("barcode = ?", book.Barcode).First(&book).Error; err == nil {
		c.JSON(202, gin.H{
			"code":  202,
			"error": "图书条码号重复"})
		return
	}

	// 处理分类名称
	book.ClassificationName = utils.GetClassificationName(book.Classification)

	// 将图书信息保存到数据库
	if err := database.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "图书信息添加成功", "data": book})
}
