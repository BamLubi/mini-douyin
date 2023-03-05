package entity

type Favorite struct {
	Id         int64 `gorm:"primary_key" json:"id"`
	UserId     int64 `json:"user_id"`
	VideoId    int64 `json:"video_id"`
	ActionType int32 `json:"action_type"`

	Video VideoInfo `json:"video" gorm:"foreignKey:Id;references:VideoId;"`
}
