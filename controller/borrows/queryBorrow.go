package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// QueryBorrow 根据查询参数查询借阅信息的控制器
func QueryBorrow(c *gin.Context) {
	// 获取查询参数
	barcode := c.Query("barcode")
	title := c.Query("title")
	userName := c.Query("username")
	pageStr := c.DefaultQuery("page", "1")    // 默认为第一页
	limitStr := c.DefaultQuery("limit", "10") // 默认每页显示10条记录

	// 将字符串参数转换为整数
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的页码"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的每页记录数"})
		return
	}

	// 构建查询条件
	db := database.DB.Model(&model.Borrow{})
	if barcode != "" {
		db = db.Where("barcode = ?", barcode)
	}
	if title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if userName != "" {
		db = db.Where("user_name LIKE ?", "%"+userName+"%")
	}

	// 查询总记录数
	var count int64
	if err := db.Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法查询总记录数", "details": err.Error()})
		return
	}

	// 分页查询
	var borrows []model.Borrow
	offset := (page - 1) * limit
	if err := db.Offset(offset).Limit(limit).Find(&borrows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法查询借阅信息", "details": err.Error()})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"data":  borrows,
		"count": count,
	})
}
