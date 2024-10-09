package router

import (
	controller "backend/controller/borrows"

	"github.com/gin-gonic/gin"
)

func Borrow(r *gin.Engine) {
	// 查询图书信息路由
	r.GET("/borrow/queryBorrowsByPage", controller.QueryBorrow)
	// 归还图书路由
	r.POST("/borrow/returnBook", controller.ReturnBook)
	//借书
	r.POST("/borrow/borrowBook", controller.BorrowBook)
	//删除借阅信息
	r.DELETE("/borrow/deleteBorrow", controller.DeleteBorrow)
	//批量删除借阅信息
	r.DELETE("/borrow/deleteBorrows", controller.DeleteBorrows)
}
