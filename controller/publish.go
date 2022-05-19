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
	token := c.PostForm("token")

	//判断用户是否是存在的
	//------------ string -> User
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{Statuscode: 1, StatusMsg: "User don't exits"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Statuscode: 1,
			StatusMsg:  err.Error(),
		})
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token] //获得user
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)

	saveFile := filepath.Join("./public/", finalName) //加在./public/目录下

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			Statuscode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Statuscode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}
