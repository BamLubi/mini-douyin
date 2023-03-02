package main

import (
	"context"
	"errors"
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

	// 写入数据库
	err = config.DB.Create(&user).Error
	if err != nil {
		config.Logger.Error("UserRegister" + err.Error())

		err_msg := err.Error()
		resp = &userdouyin.UserRegisterResponse{StatusCode: 1, StatusMsg: &err_msg, UserId: user.Id, Token: ""}
		return resp, err
	}

	// 返回结果
	resp = &userdouyin.UserRegisterResponse{StatusCode: 0, UserId: user.Id, Token: ""}
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *userdouyin.UserLoginRequest) (resp *userdouyin.UserLoginResponse, err error) {
	// TODO: Your code here...
	// 从数据库中获取用户的盐值和hash，匹配是否一致，
	var user entity.User
	config.DB.Model(&entity.User{}).Where("username = ?", req.Username).First(&user)
	if utils.Encrypt(req.Password+user.Salt) != user.Hash {
		err_msg := "password wrong"
		resp = &userdouyin.UserLoginResponse{StatusCode: 1, StatusMsg: &err_msg, UserId: -1, Token: ""}
		err = errors.New("password wrong")
	}
	resp = &userdouyin.UserLoginResponse{StatusCode: 0, UserId: user.Id, Token: ""}
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *userdouyin.UserInfoRequest) (resp *userdouyin.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the UserServiceImpl interface.
func (s *UserServiceImpl) PublishList(ctx context.Context, req *userdouyin.PublishListRequest) (resp *userdouyin.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}
