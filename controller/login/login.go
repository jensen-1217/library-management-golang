package controller

import (
	"backend/dao/database"
	"backend/model"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 登录
func Login(ac *gin.Context) {
	// 创建结构体用于接收参数
	var loginInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 使用 BindJSON 方法解析请求体中的 JSON 数据
	if err := ac.BindJSON(&loginInfo); err != nil {
		ac.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数解析失败",
		})
		return
	}

	// 从结构体中获取用户名和密码
	username := loginInfo.Username
	password := loginInfo.Password

	// 查询用户
	var user model.User
	if err := database.DB.Where("user_name = ?", username).First(&user).Error; err != nil {
		ac.JSON(202, gin.H{
			"code":    202,
			"message": "用户名或密码错误",
		})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ac.JSON(202, gin.H{
			"code":    202,
			"message": "用户名或密码错误",
		})
		return
	}

	// 生成 JWT Token
	token, err := utils.GenerateToken(username)
	if err != nil {
		// 处理生成令牌失败的情况
		ac.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "无法生成 token",
		})
	}
	// 返回 Token
	ac.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "登录成功",
		"data": gin.H{
			"token": token,
		},
	})

}
