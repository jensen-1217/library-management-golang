package router

import (
	controller "backend/controller/user"

	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	//获取用户数据
	r.GET("/user/info", controller.UserInfo)
	//获取用户信息
	r.GET("/user/queryUsersByPage", controller.QueryUsersByPage)
	//添加用户
	r.POST("/user/addUser", controller.AddUser)
	//获取用户数量
	r.GET("/user/getCount", controller.GetCount)
	//更新用户数据
	r.PUT("/user/updateUser", controller.UpdateUser)
	//删除用户
	r.DELETE("/user/deleteUser", controller.DeleteUser)
	//批量删除用户
	r.DELETE("/user/deleteUsers", controller.DeleteUsers)
}
