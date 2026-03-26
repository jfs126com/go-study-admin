package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewGormDB() {
	dsn := "root:!qaz@wsx@tcp(1.95.218.59:3306)/go-study-admin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	//自动建表
	err = db.AutoMigrate(&SysUser{})
	DB = db
}
