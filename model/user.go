package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string `json:"username" gorm:"not null;unique" binding:"required"`
	Name        string `json:"name" gorm:"not null" binding:"required"`
	PhoneNumber string `json:"phonenumber" gorm:"not null;unique" binding:"required"`
	Password    string `json:"password" gorm:"not null" binding:"required"`
	IsAdmin     int    `json:"isadmin" gorm:"not null"`
}
