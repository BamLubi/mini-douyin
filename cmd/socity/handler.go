package main

import (
	"context"
	"log"
	"math/rand"
	socitydouyin "mini-douyin/cmd/socity/kitex_gen/socitydouyin"
	"mini-douyin/cmd/socity/kitex_gen/userdouyin"
	"mini-douyin/cmd/socity/rpc"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/entity"
	"mini-douyin/pkg/utils"
	"strconv"
	"time"
)

// SocityServiceImpl implements the last service interface defined in the IDL.
type SocityServiceImpl struct{}

// FavoriteAction implements the SocityServiceImpl interface.
func (s *SocityServiceImpl) FavoriteAction(ctx context.Context, req *socitydouyin.FavoriteActionRequest) (resp *socitydouyin.FavoriteActionResponse, err error) {
	// 无论是点赞还是取消点赞都修改数据库
	// TODO：使用redis优化时在考虑
	// 1. 将点赞数据先存储在redis中，用hashtable存储
	// 2. 监听redis过期，写入数据库
	randSecond := 1*60*60 + rand.Intn(30*60)
	randSecond = 10 // 测试用
	config.RD.Send("SELECT", 10)
	config.RD.Send("SET", "user_favorite_ex_"+strconv.Itoa(int(req.UserId)), "expire", "EX", randSecond)
	config.Logger.Info("user_favorite_ex_" + strconv.Itoa(int(req.UserId)))
	config.RD.Send("HSET", "user_favorite_"+strconv.Itoa(int(req.UserId)), strconv.Itoa(int(req.VideoId)), strconv.Itoa(int(req.ActionType)))
	if req.ActionType == 1 {
		config.RD.Send("HINCRBY", "video_"+strconv.Itoa(int(req.VideoId)), "favorite_count", 1)
	} else {
		config.RD.Send("HINCRBY", "video_"+strconv.Itoa(int(req.VideoId)), "favorite_count", -1)
	}
	config.RD.Send("SET", "video_ex_"+strconv.Itoa(int(req.VideoId)), "expire", "EX", 30)
	config.RD.Flush()

	defer func() {
		if err := recover(); err != nil {
			config.RD.Do("DEL", "user_favorite_ex_"+strconv.Itoa(int(req.UserId)))
			config.RD.Do("DEL", "user_favorite_"+strconv.Itoa(int(req.UserId)))

			config.Logger.Error(err.(string))
			err_msg := "FavoriteAction Redis error"
			resp = &socitydouyin.FavoriteActionResponse{StatusCode: 1, StatusMsg: &err_msg}
		}
	}()

	if _, err := config.RD.Receive(); err != nil {
		panic("Redis error 1 " + err.Error())
	}
	if _, err := config.RD.Receive(); err != nil {
		log.Println(err)
		panic("Redis error 2 " + err.Error())
	}
	if _, err := config.RD.Receive(); err != nil {
		log.Println(err)
		panic("Redis error 3 " + err.Error())
	}
	if _, err := config.RD.Receive(); err != nil {
		log.Println(err)
		panic("Redis error 4 " + err.Error())
	}
	resp = &socitydouyin.FavoriteActionResponse{StatusCode: 0}
	return
}

// CommentAction implements the SocityServiceImpl interface.
func (s *SocityServiceImpl) CommentAction(ctx context.Context, req *socitydouyin.CommentActionRequest) (resp *socitydouyin.CommentActionResponse, err error) {
	if req.ActionType == 1 {
		nowDate := time.Now().Format("01-02")
		comment := &entity.Comment{Id: utils.IdGen(), UserId: req.UserId, VideoId: req.VideoId, Content: *req.CommentText, CreateDate: nowDate, Active: 1}
		// err = config.DB.Create(comment).Error
		err = config.DB.Save(comment).Error
		if err != nil {
			config.Logger.Error(err.Error())
			err_msg := "FavoriteAction DB error"
			resp = &socitydouyin.CommentActionResponse{StatusCode: 1, StatusMsg: &err_msg}
			return
		}
		// 调用rpc获取当前用户信息
		userInfoReq := &userdouyin.UserInfoRequest{UserId: req.UserId}
		var userInfoResq *userdouyin.UserInfoResponse
		userInfoResq, err = rpc.UserClient.UserInfo(ctx, userInfoReq)
		if err != nil {
			err_msg := "FavoriteAction DB error"
			resp = &socitydouyin.CommentActionResponse{StatusCode: 1, StatusMsg: &err_msg}
			return
		} // TODO: 出错需要回滚
		user := userInfoResq.User
		commentIDL := EntityComment2IDLComment(comment, user)
		resp = &socitydouyin.CommentActionResponse{StatusCode: 0, Comment: commentIDL}
		return
	} else {
		err = config.DB.Model(&entity.Comment{}).Where("id = ?", req.CommentId).Update("active", 0).Error
		if err != nil {
			config.Logger.Error(err.Error())
			err_msg := "FavoriteAction DB error"
			resp = &socitydouyin.CommentActionResponse{StatusCode: 1, StatusMsg: &err_msg}
			return
		}
		resp = &socitydouyin.CommentActionResponse{StatusCode: 0}
		return
	}
}

// CommentList implements the SocityServiceImpl interface.
func (s *SocityServiceImpl) CommentList(ctx context.Context, req *socitydouyin.CommentListRequest) (resp *socitydouyin.CommentListResponse, err error) {
	var commentList []*entity.Comment
	err = config.DB.Table("comment").Preload("User").Where("comment.video_id = ?", req.VideoId).Order("comment.create_date DESC").Find(&commentList).Error
	if err != nil {
		config.Logger.Error(err.Error())
		err_msg := "comment DB error"
		resp = &socitydouyin.CommentListResponse{StatusCode: 1, StatusMsg: &err_msg}
		return
	}
	commentListIDL := EntityCommentList2IDLCommentList(commentList)
	resp = &socitydouyin.CommentListResponse{StatusCode: 0, CommentList: commentListIDL}

	return
}
