package router

import (
	controller "backend/controller/bookinfo"

	"github.com/gin-gonic/gin"
)

func Bookinfo(r *gin.Engine) {
	//增加图书
	r.POST("/bookInfo/addBookInfo", controller.AddBookInfo)
	//分页查询图书
	r.GET("/bookInfo/queryBookInfosByPage", controller.QueryBookInfo)
	//更新图书信息
	r.PUT("/bookInfo/updateBookInfo", controller.UpdateBookInfo)
	//获取图书数量
	r.GET("/bookInfo/getCount", controller.GetBookCount)
	//删除图书
	r.DELETE("/bookInfo/deleteBookInfo", controller.DeleteBookInfo)
	//批量删除图书
	r.DELETE("/bookInfo/deleteBookInfos", controller.DeleteBooks)
	// 上传图书信息表格
	r.POST("/bookInfo/batchAddBookInfos", controller.UploadBookInfo)
}
