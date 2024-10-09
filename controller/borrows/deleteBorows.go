package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteBorrows 批量删除借阅信息
func DeleteBorrows(c *gin.Context) {
	// 定义借阅记录列表，用于接收前端传递的要删除的借阅记录
	var borrows []model.Borrow

	// 绑定 JSON 数据
	if err := c.BindJSON(&borrows); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid borrow data"})
		return
	}

	// 遍历借阅记录列表，逐个执行删除借阅记录的逻辑
	for _, borrow := range borrows {
		// 如果借阅记录未归还图书，则不允许删除
		if borrow.Returntimestr == "" {
			c.JSON(201, gin.H{
				"code":    201,
				"message": "Borrow record with ID " + strconv.Itoa(int(borrow.ID)) + " has not been returned yet"})
			return
		}

		// 执行删除借阅记录的逻辑
		if err := database.DB.Delete(&borrow).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "Failed to delete borrow"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Borrows deleted successfully"})
}
