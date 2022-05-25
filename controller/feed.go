package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//feed流

type FeedResponse struct {
	Response          //基本回复
	VideoList []Video `json:"video_list,omitempty"` //Video 列表
	NextTime  int64   `json:"next_time,omitempty"`
}

//Feed same demo video list for every request

func Feed(c *gin.Context) { //现在返回的是StatusCode :0 & Unixcurtime :23563263

	//进行VedioList的查找
	c.JSON(http.StatusOK, FeedResponse{ //将给的结构体序列化成为Json， 传到response Body中
		Response:  Response{Statuscode: 0}, //成功就是返回0
		VideoList: DemoVideos,              //Video List
		NextTime:  time.Now().Unix(),
	})
}
