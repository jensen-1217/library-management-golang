package router

import (
	controller "backend/controller/password"

	"github.com/gin-gonic/gin"
)

func Password(r *gin.Engine) {
	//修改密码
	r.POST("/user/alterPassword", controller.ChangePassword)
}
