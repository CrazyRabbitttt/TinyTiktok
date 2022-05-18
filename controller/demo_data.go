package controller

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 99,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "测试评论呐，啦啦啦",
		CreateDate: "12-30",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "Shaoguixin",
	FollowCount:   10000,
	FollowerCount: 1000000,
	IsFollow:      false,
}
