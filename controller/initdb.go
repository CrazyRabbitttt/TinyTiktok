package controller

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {
	//进行数据库的连接

	dsn := "root:woshimima@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"

	db1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db1 //将数据传送到全局变量中

	if err != nil {
		fmt.Println("数据库连接失败！")
		return
	} else {
		fmt.Println("数据库连接成功！")
	}

}
