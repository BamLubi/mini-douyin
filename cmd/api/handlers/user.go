package handlers

import (
	"context"
	"mini-douyin/cmd/api/kitex_gen/userdouyin"
	"mini-douyin/cmd/api/middleware"
	"mini-douyin/cmd/api/rpc"
	"strconv"

	config "mini-douyin/pkg/configs"

	"github.com/gin-gonic/gin"
)

// 用户注册
func UserRegister(c *gin.Context) {
	// 获取参数，参数为 Param 格式
	req := &userdouyin.UserRegisterRequest{Username: c.Query("username"), Password: c.Query("password")}
	// 调用RPC执行用户注册，返回用户的id即可
	resp, err := rpc.UserClient.UserRegister(context.Background(), req)
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}
	// 根据用户的id生成相应的token
	if resp.StatusCode == 0 {
		token, err := middleware.GenTokenStringFromUserId(resp.UserId)
		if err != nil {
			c.JSON(400, resp)
			return
		}
		resp.Token = token
	}
	c.JSON(200, resp)
}

// 用户个人信息
func UserInfo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	// 校验id与token解析id是否一致
	if err := middleware.CheckUserId2TokenString(id, c.Query("token")); err != nil {
		config.Logger.Error(err.Error())
		c.JSON(400, &userdouyin.UserInfoResponse{StatusCode: 1})
		return
	}
	// 发送请求
	req := &userdouyin.UserInfoRequest{UserId: id}
	resp, err := rpc.UserClient.UserInfo(context.Background(), req)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	c.JSON(200, resp)
}

func PublishList(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	// 校验id与token解析id是否一致
	if err := middleware.CheckUserId2TokenString(id, c.Query("token")); err != nil {
		config.Logger.Error(err.Error())
		c.JSON(400, &userdouyin.UserInfoResponse{StatusCode: 1})
		return
	}
	req := &userdouyin.PublishListRequest{UserId: id}
	resp, err := rpc.UserClient.PublishList(context.Background(), req)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	c.JSON(200, resp)
}

func FavoriteList(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	// 校验id与token解析id是否一致
	if err := middleware.CheckUserId2TokenString(id, c.Query("token")); err != nil {
		config.Logger.Error(err.Error())
		c.JSON(400, &userdouyin.FavoriteListResponse{StatusCode: 1})
		return
	}
	req := &userdouyin.FavoriteListRequest{UserId: id}
	resp, err := rpc.UserClient.FavoriteList(context.Background(), req)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	c.JSON(200, resp)
}
