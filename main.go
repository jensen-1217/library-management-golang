package main

import (
	"backend/dao/database"
	"backend/router"
)

func main() {

	//连接数据库
	database.InitDB()

	//实例化路由
	router.InitRouter()

}
