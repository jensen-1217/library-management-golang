package router

import (
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 初始化路由
	r := gin.Default()

	// 配置CORS中间件
	utils.UseCorsMiddleware(r)
	// 配置静态文件服务
	r.Static("/public", "./public")
	//登录
	Login(r)
	//注册
	Register(r)
	//退出
	Logout(r)
	//用户
	User(r)
	//借阅
	Borrow(r)
	//图书信息
	Bookinfo(r)
	//修改密码
	Password(r)
	//在30002端口启动服务
	r.Run(":30002")
	return r
}
