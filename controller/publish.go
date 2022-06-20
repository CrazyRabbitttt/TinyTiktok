package controller

import (
	"Web-Go/Common"
	"Web-Go/ConnSql"
	"Web-Go/Model"
	"Web-Go/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

//传入的参数：data, token, title
func Publish(c *gin.Context) {
	db := ConnSql.ThemodelOfSql()
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

	filename := filepath.Base(data.Filename)             //去除掉路径的文件名称
	finalName := fmt.Sprintf("%d_%s", user.Id, filename) //最终存储的文件名称
	saveFile := filepath.Join("../public/", finalName)   //拼接成为绝对路径

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

//将视频上传到本地的服务器上面
func Publish1(c *gin.Context) {
	//1. 通过token 获得用户名
	token := c.PostForm("token")
	userId := service.GetUserByToken(token)
	println("通过token获取到的用户id：", userId)
	title := c.PostForm("title")
	println("获取到的title:", title)
	//2.进行传入视频文件位置的设置
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)            //去除掉路径的文件名称
	finalName := fmt.Sprintf("%d_%s", userId, filename) //最终存储的文件名称
	saveFile := filepath.Join("./public/", finalName)   //拼接成为绝对路径
	println("The save path is :", saveFile)

	if err := c.SaveUploadedFile(data, saveFile); err != nil { //进行文件的存储
		c.JSON(http.StatusOK, Common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error() + "文件存储失败",
		})
		return
	}
	c.JSON(http.StatusOK, Common.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " upload successfully",
	})

	//进行video的信息的配置
	playUrl := "http://192.168.43.104:8080/" + "public/" + finalName
	coverUrl := "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"

	video := Model.Video{
		Model:         gorm.Model{},
		AuthorId:      uint(userId),
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
	}
	service.CreateVideo(&video)
}

//
//func PublishList(c *gin.Context) {
//	c.JSON(http.StatusOK, VideoListResponse{
//		Response: Response{
//			Statuscode: 0,
//		},
//		VideoList: DemoVideos,
//	})
//}
