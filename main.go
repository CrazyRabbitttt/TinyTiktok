package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default() //创建gin路由

	initRouter(r)

	r.Run() //Listen and serve on 0.0.0.0:8008

}
