package controller

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var DB sql.DB

func InitDb() {
	//数据库的连接
	//go原生db
	DB, _ := sql.Open("mysql", "root:shaoguixin1+@(127.0.0.1:3306)/tiktok") // 设置连接数据库的参数
	defer DB.Close()                                                        //关闭数据库
	err := DB.Ping()                                                        //连接数据库
	if err != nil {
		fmt.Println("原生数据库连接失败")
		return
	} else {
		fmt.Println("原生连接数据库成功！")
	}

	//gorm:

	dsn := "root:shaoguixin1+@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"

	db1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db1 //将数据传送到全局变量中

	if err != nil {
		fmt.Println("数据库连接失败！")
		return
	} else {
		fmt.Println("数据库连接成功！")
	}

}
