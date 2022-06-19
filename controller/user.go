package controller

import (
	"Web-Go/Common"
	"Web-Go/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserIdTokenResponse struct {
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

type UserRegisterResponse struct {
	Common.Response
	UserIdTokenResponse
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
	db.Table("tik_user").Where("name = ?", username).First(&dbUser)
	if dbUser.Id > 0 {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{Statuscode: 1, StatusMsg: "抱歉，该用户已经存在！"},
		})
		return
	}
	db.Table("tik_user").Create(&dbUser)

	//传入的是&，目前已经是
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{Statuscode: 0, StatusMsg: "用户" + username + "注册成功！"},
		UserId:   dbUser.Id,
		Token:    token,
	})

}

//用户登陆的处理函数，例如处理一下是否是存在等
func UserregisterService(userName string, passWord string) (UserIdTokenResponse, error) {

	var userResponse = UserIdTokenResponse{}

	//1.Legal check
	err := service.IsUserLegal(userName, passWord)
	if err != nil {
		return userResponse, err
	}
	//2.Create New User
	newUser, err := service.CreateNewUser(userName, passWord)
	if err != nil {
		return userResponse, err
	}
	print(newUser)
	//3.颁发token
	return userResponse, nil
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	fmt.Println(username, password, "efhduiwefhui")
	token := username + password + "lala"

	var dbUser User

	db.Table("tik_user").Where("Name = ?", username).Find(&dbUser)

	fmt.Println("Login user:", dbUser)

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{Statuscode: 0, StatusMsg: "用户" + username + "登陆成功！"},
		UserId:   dbUser.Id,
		Token:    token,
	})
}

func UserInfo(c *gin.Context) {
	userId := c.Query("user_id")

	var dbUser User

	fmt.Println("传入的user_id", userId)
	db.Table("tik_user").Where("id = ?", userId).Find(&dbUser)

	fmt.Println("查到的用户信息：", dbUser)

	if dbUser.Id != 0 {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{Statuscode: 0, StatusMsg: "查询用户信息成功"},
			User: User{
				Id:            dbUser.Id,
				Name:          dbUser.Name,
				FollowCount:   188,
				FollowerCount: 199,
				IsFollow:      true,
			},
		})
	} else {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{Statuscode: 1, StatusMsg: "查询失败"},
			UserId:   dbUser.Id,
		})
	}

}
