package router

import (
	controller "backend/controller/logout"

	"github.com/gin-gonic/gin"
)

func Logout(r *gin.Engine) {
	//退出
	r.POST("/user/logout", controller.Logout)
}
