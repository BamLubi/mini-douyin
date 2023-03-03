package entity

type UserInfo struct {
	Id              int64  `json:"id" gorm:"primary_key"`            // id
	Name            string `json:"name"`                             // 用户名称
	FollowCount     int64  `gorm:"default:0" json:"follow_count"`    // 关注总数
	FollowerCount   int64  `gorm:"default:0" json:"follower_count"`  // 粉丝总数
	IsFollow        bool   `gorm:"default:0" json:"is_follow"`       // true-已关注，false-未关注
	Avatar          string `json:"avatar"`                           // 头像
	BackgroundImage string `json:"background_image"`                 // 个人页顶部大图
	Signature       string `json:"signature"`                        // 个人简介
	TotalFavorited  int64  `gorm:"default:0" json:"total_favorited"` // 获赞数量
	WorkCount       int64  `gorm:"default:0" json:"work_count"`      // 作品数量
	FavoriteCount   int64  `gorm:"default:0" json:"favorite_count"`  // 点赞数量
}

func (ui UserInfo) TableName() string {
	return "userinfo"
}
