package entity

type Comment struct {
	Id         int64  `gorm:"primary_key" json:"id"`
	UserId     int64  `json:"user_id"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}
