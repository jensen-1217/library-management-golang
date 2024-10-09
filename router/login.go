package router

import (
	controller "backend/controller/login"
	"backend/dao"

	"github.com/gin-gonic/gin"
)

func Login(r *gin.Engine) {

	r.POST("/initializeDatabase", dao.SeedData)

	//登录
	r.POST("/login", controller.Login)
}
