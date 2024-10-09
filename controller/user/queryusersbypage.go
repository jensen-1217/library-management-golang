package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 分页查询用户信息
func QueryUsersByPage(c *gin.Context) {
	// 解析请求参数
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	username := c.Query("username")

	// 将字符串参数转换为整数
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid limit"})
		return
	}

	// 构建查询
	db := database.DB.Model(&model.User{})
	if username != "" {
		db = db.Where("user_name = ?", username)
	}

	// 查询总数
	var count int64
	if err := db.Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to count users"})
		return
	}

	// 查询用户列表
	var users []model.User
	offset := (page - 1) * limit
	if err := db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to query users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"data":  users,
		"count": count})
}
