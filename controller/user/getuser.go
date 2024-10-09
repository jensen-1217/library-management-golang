package controller

import (
	"backend/dao/database"
	"backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	var count int64

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// 获取当前页数据
	database.DB.Limit(limit).Offset(offset).Find(&users)

	// 获取总记录数
	database.DB.Model(&model.User{}).Count(&count)

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"count": count,
		"data":  users,
	})
}
