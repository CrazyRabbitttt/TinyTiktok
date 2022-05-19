package main

import (
	"Web-Go/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func lalala() {
	fmt.Println("efswhiuwefuiwefguiweghui")
}

func initRouter(r *gin.Engine) {
	//	公共目录用来服务静态的资源

	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	//basic apis

	apiRouter.GET("/feed/", controller.Feed) //feed流的handler

	apiRouter.POST("/user/register/", controller.Register)

	apiRouter.POST("/user/login/", controller.Login)

	apiRouter.GET("/user/", controller.UserInfo)

	apiRouter.POST("/publish/action/", controller.Publish)

}
