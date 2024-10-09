package database

import (
	"backend/model"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBInfor struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	cfg := GetDBInfor()
	dsn := cfg.DBUsername + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("初始化数据库连接失败", err)
	}

	DB = db
	if err := DB.AutoMigrate(model.User{}, model.Book{}, model.Borrow{}); err != nil {
		fmt.Println("创建表结构失败")

	}

	// 向用户表中添加管理员账号信息
	adminUser := model.User{
		UserName:    "admin",
		Name:        "管理员",
		PhoneNumber: "13733737247",
		IsAdmin:     1,
	}

	// 对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("密码加密错误", err)
		return nil, err
	}
	adminUser.Password = string(hashedPassword)

	// 查询数据库中是否存在相同的用户名
	var existingAdmin model.User
	err = DB.Where("user_name = ?", adminUser.UserName).First(&existingAdmin).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 查询过程中发生错误，返回错误信息
		fmt.Println("查询管理员用户错误", err)
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 用户名不存在，创建管理员用户
		if err := DB.Create(&adminUser).Error; err != nil {
			fmt.Println("创建管理员用户错误", err)
			return nil, err
		}
	}

	// 添加读者信息
	readerUser := model.User{
		UserName:    "reader",
		Name:        "李华",
		PhoneNumber: "13733677777",
		IsAdmin:     0,
	}
	readerUser.Password = string(hashedPassword)

	// 查询数据库中是否存在相同的用户名
	var existingReader model.User
	err = DB.Where("user_name = ?", readerUser.UserName).First(&existingReader).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 查询过程中发生错误，返回错误信息
		fmt.Println("查询读者用户错误", err)
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 用户名不存在，创建读者用户
		if err := DB.Create(&readerUser).Error; err != nil {
			fmt.Println("创建读者用户错误", err)
			return nil, err
		}
	}

	return db, nil

}

func GetDBInfor() *DBInfor {
	return &DBInfor{
		DBUsername: "root",
		DBPassword: "root",
		DBHost:     "localhost",
		DBPort:     "3306",
		DBName:     "njshw",
	}
}
