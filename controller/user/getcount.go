package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCount(c *gin.Context) {
	// 查询用户总数
	var count int64
	if err := database.DB.Model(&model.User{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "无法查询用户数量",
		})
		return
	}

	// 返回用户数量
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "成功获取用户数量",
		"count":   count,
	})
}
