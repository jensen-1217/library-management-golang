package model

import (
	"gorm.io/gorm"
)

type Borrow struct {
	gorm.Model
	Barcode       string `json:"barcode"`
	Title         string `json:"title"`
	UserName      string `json:"username"`
	Borrowtimestr string `json:"borrowtimestr"`
	Returntimestr string `json:"returntimestr"`
}
