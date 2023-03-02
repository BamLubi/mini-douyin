package entity

type VideoInfo struct {
	Id            int64  `json:"id"`             // id
	UserId        string `json:"user_id"`        // 作者id
	PlayUrl       string `json:"play_url"`       // 播放地址
	CoverUrl      string `json:"cover_url"`      // 封面地址
	FavoriteCount int64  `json:"favorite_count"` // 点赞总数
	CommentCount  int64  `json:"comment_count"`  // 评论总数
	IsFavorite    string `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string `json:"title"`          // 标题
}

func (vi VideoInfo) TableName() string {
	return "videoinfo"
}
