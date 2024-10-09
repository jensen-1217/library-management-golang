package router

import (
	controller "backend/controller/register"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	//注册
	r.POST("/register", controller.Register)
}
