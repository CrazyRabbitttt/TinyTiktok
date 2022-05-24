package controller

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 99,
		CommentCount:  0,
		IsFavorite:    true,
	},
	{
		Id:            4,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.43.104:8080/static/24_trailer.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 893,
		CommentCount:  0,
		IsFavorite:    true,
	},
	{
		Id:            5,
		Author:        DemoUser,
		PlayUrl:       "http://vjs.zencdn.net/v/oceans.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 299,
		CommentCount:  0,
		IsFavorite:    true,
	},
	{
		Id:            5,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.43.104:8080/static/24_vedio1.mov",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 23,
		CommentCount:  0,
		IsFavorite:    true,
	},
}

var DemoVideos1 = []Video{
	{
		Id:            2,
		Author:        DemoUser,
		PlayUrl:       "https://media.w3.org/2010/05/sintel/trailer.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 23,
		CommentCount:  0,
		IsFavorite:    true,
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
