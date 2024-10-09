package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteBorrow 删除借阅信息的控制器
func DeleteBorrow(c *gin.Context) {
	var borrow model.Borrow

	// 绑定 JSON 请求体到 Borrow 结构体
	if err := c.BindJSON(&borrow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体"})
		return
	}

	// 根据 ID 查询借阅信息
	if err := database.DB.Where("id = ?", borrow.ID).First(&borrow).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到借阅信息"})
		return
	}

	// 如果还书时间为空，则不允许删除
	if borrow.Returntimestr == "" {
		c.JSON(201, gin.H{
			"code":  201,
			"error": "借阅信息尚未归还，无法删除"})
		return
	}

	// 执行删除操作
	if err := database.DB.Delete(&borrow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法删除借阅信息", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "借阅信息删除成功"})
}
