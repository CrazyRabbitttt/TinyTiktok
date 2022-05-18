package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
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

var userIdsequence = int64(1) //进行id自增

type UserResponse struct {
	Response
	User User `json:"user"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	fmt.Println(token)

	//传入token查看是否是存在的
	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			//如果存在了就返回说已经存在了
			Response: Response{Statuscode: 1, StatusMsg: "User already exist"},
		})
	} else {
		//新创建一个user
		atomic.AddInt64(&userIdsequence, 1)
		newUser := User{
			Id:   userIdsequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{Statuscode: 0, StatusMsg: "Successfully made a new User"},
			UserId:   userIdsequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{Statuscode: 0, StatusMsg: "Weclome " + username},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{Statuscode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{Statuscode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{Statuscode: 1, StatusMsg: "User don't exist"},
		})
	}
}