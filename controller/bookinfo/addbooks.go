package controller

import (
	"backend/dao/database"
	"backend/model"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

// UploadBookInfo 上传图书信息表格
func UploadBookInfo(c *gin.Context) {
	// 从请求中获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 打开文件
	excelFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer excelFile.Close()

	// 解析 Excel 文件
	xlFile, err := xlsx.OpenReaderAt(excelFile, file.Size)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 遍历 Excel 文件中的每个工作表
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			// 如果是第一行（标题行），则跳过
			if i == 0 {
				continue
			}
			// 解析每一行数据
			var book model.Book
			cells := row.Cells
			if len(cells) < 9 {
				fmt.Println("行数据不完整")
				continue
			}
			book.Title = cells[0].String()
			book.Author = cells[1].String()
			book.Price = cells[2].String()
			book.Barcode = cells[3].String()
			book.ISBN = cells[4].String()
			book.Classification = cells[5].String()
			book.Publisher = cells[6].String()
			book.Bookimg = cells[7].String()
			book.Info = cells[8].String()
			book.Status = 1

			// 检查图书的 Barcode 是否已存在
			var existingBook model.Book
			if err := database.DB.Where("barcode = ?", book.Barcode).First(&existingBook).Error; err == nil {
				fmt.Printf("图书条码号重复: %s\n", book.Barcode)
				continue
			}

			// 处理分类名称
			book.ClassificationName = utils.GetClassificationName(book.Classification)

			// 将图书信息保存到数据库
			if err := database.DB.Create(&book).Error; err != nil {
				fmt.Printf("图书信息添加失败: %v\n", err)
				continue
			}

			fmt.Printf("图书信息添加成功: %s\n", book.Title)
		}
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "图书信息添加成功"})
}
