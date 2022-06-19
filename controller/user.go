package controller

import (
	"Web-Go/Common"
	"Web-Go/ConnSql"
	"Web-Go/Model"
	"Web-Go/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserIdTokenResponse struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserRegisterResponse struct {
	Common.Response
	UserIdTokenResponse
}

//用户注册的主函数， 最上层的接口函数
func UserRegister(c *gin.Context) {
	//传进来的参数的获取
	username := c.Query("username")
	password := c.Query("password")

	//进行service层次的处理
	registerResponse, err := UserRegisterService(username, password)

	//将响应进行返回
	if err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserRegisterResponse{
		Response:            Common.Response{StatusCode: 0},
		UserIdTokenResponse: registerResponse,
	})
	return
}

//用户进行登陆的处理函数，鉴别是否是存在的等
func UserRegisterService(userName string, passWord string) (UserIdTokenResponse, error) {

	var userResponse = UserIdTokenResponse{}

	//1.Legal check
	err := service.IsUserLegal(userName, passWord)
	if err != nil {
		return userResponse, err
	}
	//2.Create New User, 返回的对象中只是有用户名、密码
	var newUser Model.User
	newUser, err = service.CreateNewUser(userName, passWord)
	if err != nil {
		if err == Common.ErrorUserExits {
			//print("Error : user exist....")
			return userResponse, Common.ErrorUserExits //将err继续传递
		}
	}
	//进行token的颁发

	token := newUser.Name + newUser.Password + "bing"

	userResponse = UserIdTokenResponse{
		UserId: newUser.Id,
		Token:  token,
	}
	return userResponse, nil
}

func Login(c *gin.Context) {
	db := ConnSql.ThemodelOfSql()
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
	db := ConnSql.ThemodelOfSql()
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
