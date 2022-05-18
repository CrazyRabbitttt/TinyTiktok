package controller

//基本的Response类型

type Response struct {
	Statuscode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"` //可省略？
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`      //封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty"` //点赞总数
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`   //关注数目
	FollowerCount int64  `json:"follower_count,omitempty"` //粉丝数目
	IsFollow      bool   `json:"is_follow,omitempty"`      //是否关注了
}
