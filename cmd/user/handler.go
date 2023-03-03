package main

import (
	"context"
	"errors"
	"mini-douyin/cmd/user/kitex_gen/base"
	userdouyin "mini-douyin/cmd/user/kitex_gen/userdouyin"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/entity"
	"mini-douyin/pkg/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *userdouyin.UserRegisterRequest) (resp *userdouyin.UserRegisterResponse, err error) {
	// 密码加盐值后hash
	salt := utils.SaltGen()
	password := req.Password + salt
	hash := utils.Encrypt(password)

	user := &entity.User{Id: utils.IdGen(), Username: req.Username, Salt: salt, Hash: hash}
	userinfo := &entity.UserInfo{Id: user.Id, Name: user.Username, IsFollow: true}

	// 启用事务创建user数据和userinfo数据
	err = RegisterTransaction(user, userinfo)
	if err != nil {
		config.Logger.Error("UserRegister: " + err.Error())

		err_msg := err.Error()
		resp = &userdouyin.UserRegisterResponse{StatusCode: 1, StatusMsg: &err_msg, UserId: user.Id, Token: ""}
		return
	}

	// 返回结果
	resp = &userdouyin.UserRegisterResponse{StatusCode: 0, UserId: user.Id, Token: ""}
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *userdouyin.UserLoginRequest) (resp *userdouyin.UserLoginResponse, err error) {
	// 从数据库中获取用户的盐值和hash，匹配是否一致，
	var user entity.User
	config.DB.Model(&entity.User{}).Where("username = ?", req.Username).First(&user)
	if utils.Encrypt(req.Password+user.Salt) != user.Hash {
		err_msg := "password wrong"
		resp = &userdouyin.UserLoginResponse{StatusCode: 1, StatusMsg: &err_msg, UserId: -1, Token: ""}
		err = errors.New("password wrong")
		return
	}
	resp = &userdouyin.UserLoginResponse{StatusCode: 0, UserId: user.Id, Token: ""}
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *userdouyin.UserInfoRequest) (resp *userdouyin.UserInfoResponse, err error) {
	var userinfo base.User
	err = config.DB.Table("userinfo").Where("id = ?", req.UserId).First(&userinfo).Error
	if err != nil {
		err_msg := "UserInfo DB error"
		resp = &userdouyin.UserInfoResponse{StatusCode: 1, StatusMsg: &err_msg}
		return
	}
	resp = &userdouyin.UserInfoResponse{StatusCode: 0, User: &userinfo}
	return
}

// PublishList implements the UserServiceImpl interface.
func (s *UserServiceImpl) PublishList(ctx context.Context, req *userdouyin.PublishListRequest) (resp *userdouyin.PublishListResponse, err error) {
	var videoList []*entity.Video
	err = config.DB.Table("videoinfo").Preload("User").Where("videoinfo.user_id = ?", req.UserId).Find(&videoList).Error
	if err != nil {
		err_msg := "PublishList DB error"
		resp = &userdouyin.PublishListResponse{StatusCode: 1, StatusMsg: &err_msg}
		return
	}
	videoListIDL := EntityVideList2IDLVideoList(videoList)
	resp = &userdouyin.PublishListResponse{StatusCode: 0, VideoList: videoListIDL}
	return
}
