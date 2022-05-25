package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

func Publish(c *gin.Context) {
	token := c.PostForm("token") //获得body传入的token

	//TODO: 判断传入的token是否有对应的用户
	data, err := c.FormFile("data") //从body中获得文件
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Statuscode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//通过token 找到对应的user，查询两次表
	var user User
	var userlogin UserLoginInfo
	db.Where("Token = ?", token).First(&userlogin)
	var queryId = userlogin.UserId
	db.Where("Id = ?", queryId).First(&user) //通过两张表进行查询

	filename := filepath.Base(data.Filename)                                   //去除掉路径的文件名称
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)                       //最终存储的文件名称
	saveFile := filepath.Join("http://192.168.43.104:8080/public/", finalName) //拼接成为绝对路径

	fmt.Println("The savepath is : ", saveFile)

	if err := c.SaveUploadedFile(data, saveFile); err != nil { //进行文件的存储
		c.JSON(http.StatusOK, Response{
			Statuscode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Statuscode: 0,
		StatusMsg:  finalName + " uploaded successfully!",
	})
}

func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			Statuscode: 0,
		},
		VideoList: DemoVideos,
	})
}
