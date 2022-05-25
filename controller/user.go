package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//首先使用内存进行管理，内存中的map
var usersLoginInfo = map[string]User{
	"shaoguixinwoshimima": {
		Id:            1,
		Name:          "shaoguixin",
		FollowCount:   1002,
		FollowerCount: 778,
		IsFollow:      true,
	},
}

func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{Statuscode: 1, StatusMsg: "用户名或者密码不能为空！"},
		})
	}
	token := username + password + "lala"

	var dbUser User
	dbUser.Id = 0
	dbUser.Name = username
	dbUser.Password = password
	db.Table("tik_user").Where("Name = ?", username).Find(&dbUser)
	fmt.Println("传入的token:", token)
	if dbUser.Id > 0 {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{Statuscode: 1, StatusMsg: "抱歉，该用户已经存在！"},
		})
	}
	db.Table("tik_user").Create(&dbUser)

	//传入的是&，目前已经是
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{Statuscode: 0, StatusMsg: "用户" + username + "注册成功！"},
		UserId:   dbUser.Id,
		Token:    token,
	})

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	fmt.Println(username, password, "efhduiwefhui")
	token := username + password + "lala"

	var dbUser User

	db.Table("tik_user").Where("Name = ?", username).Find(&dbUser)

	fmt.Println("Login user", dbUser.Name)

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{Statuscode: 0, StatusMsg: "用户" + username + "登陆成功！"},
		UserId:   dbUser.Id,
		Token:    token,
	})
	//var userInfo UserLoginInfo
	////根据token进行查找
	//db.Where("Token = ?", token).Find(&userInfo) //查询到的结果拿到user中
	//
	//fmt.Println("Login, user: ", userInfo.Name)
	//
	//c.JSON(http.StatusOK, UserLoginResponse{
	//	Response: Response{Statuscode: 0, StatusMsg: "Weclome " + userInfo.Name},
	//	UserId:   userInfo.UserId,
	//	Token:    token,
	//})

}

func UserInfo(c *gin.Context) {
	userId := c.Query("user_id")
	var dbUser User

	db.Table("tik_user").Where("user_id = ?", userId).Find(&dbUser)

	c.JSON(http.StatusOK, UserResponse{
		Response: Response{Statuscode: 0, StatusMsg: "查询用户信息成功"},
		User:     dbUser,
	})
}
