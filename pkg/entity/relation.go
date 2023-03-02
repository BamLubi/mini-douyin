package entity

type Relation struct {
	Id       int64 `gorm:"primary_key" json:"id"`
	UserId   int64 `json:"user_id"`
	TargetId int64 `json:"target_id"`
	IsActive int   `json:"is_active"`
}