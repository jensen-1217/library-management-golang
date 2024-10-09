package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Barcode            string `json:"barcode" gorm:"not null;unique" binding:"required"`
	Title              string `json:"title" gorm:"not null" binding:"required"`
	Classification     string `json:"classification"`
	ClassificationName string `json:"classification_name" `
	Author             string `json:"author" gorm:"not null" binding:"required"`
	Publisher          string `json:"publisher"`
	Price              string `json:"price" gorm:"not null" binding:"required"`
	ISBN               string `json:"ISBN"`
	Info               string `json:"info"`
	Status             int    `json:"status"`
	Bookimg            string `json:"bookimg"`
}
