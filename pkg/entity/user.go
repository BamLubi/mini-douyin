package entity

import "strconv"

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Salt     string `json:"salt"`
	Hash     string `json:"hash"`
}

func (u User) TableName() string {
	return "user"
}

func (u User) String() string {
	return "Username:" + u.Username + ",Id:" + strconv.Itoa(int(u.Id)) + ",Salt:" + u.Salt + ",Hash:" + u.Hash
}
