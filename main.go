package main

import (
	"Web-Go/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default() //创建gin路由

	controller.InitDb() //初始化数据库连接

	initRouter(r)

	r.Run() //Listen and serve on 0.0.0.0:8008

}
