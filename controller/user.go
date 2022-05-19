package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

//用户登陆的时候的结构体，需要添加上token，同时添加好Name & Id

type UserLoginInfo struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Token string
}

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
		//下面不需要设置Id了，因为我们在数据库中设置了自增Id了
		atomic.AddInt64(&userIdsequence, 1) //Id是递增的，每次都增加1，但是需要从users中
		newUser := User{
			Id:   userIdsequence,
			Name: username,
		}
		//usersLoginInfo[token] = newUser
		db.Create(newUser) //将数据添加到表中,创建数据添加到User表中，没有密码
		//数据库自动添加Id，我们还是需要查出来然后进行返回回去

		fmt.Println("The newUser's id :", newUser.Id)
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

	var user UserLoginInfo
	//根据token进行查找
	db.Where("Token = ?", "token").Find(&user) //查询到的结果拿到user中

	fmt.Println("Login, user: ", user.Name)

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{Statuscode: 0, StatusMsg: "Weclome " + username},
		UserId:   user.Id,
		Token:    token,
	})

	//下面是从内存中map进行查询，只能单次查询
	/*
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
		}*/
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
