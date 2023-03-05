package entity

type Comment struct {
	Id         int64  `gorm:"primary_key" json:"id"`
	UserId     int64  `json:"user_id"`
	VideoId    int64  `json:"video_id"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
	Active     int    `json:"active"`

	User UserInfo `json:"user" gorm:"foreignKey:Id;references:UserId;"`
}

func (c Comment) TableName() string {
	return "comment"
}
